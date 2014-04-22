package controllers

import (
    "encoding/json"
    "fmt"
    "github.com/uas/connect"
    "github.com/uas/models"
    "github.com/uas/utils"
    "strconv"
)

func (c *BaseController) Handler() *Handler {
    return new(Handler)
}

type Handler struct {
    Controller
}

func (this *Handler) Index() {
    var (
        request  interface{}
        response = ""
    )
    this.Response.Header().Set("Access-Control-Allow-Origin", "*")
    this.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    this.Response.Header().Set("Content-type", "application/json")

    decoder := json.NewDecoder(this.Request.Body)
    err := decoder.Decode(&request)
    utils.HandleErr("[Handler] Decode :", err)
    data := request.(map[string]interface{})

    switch data["action"] {
    case "register":
        login, password := data["login"].(string), data["password"].(string)
        response = this.HandleRegister(login, password)
        fmt.Fprintf(this.Response, "%s", response)
        break

    case "login":
        login, password := data["login"].(string), data["password"].(string)
        response = this.HandleLogin(login, password)
        fmt.Fprintf(this.Response, "%s", response)
        break

    case "logout":
        response = this.HandleLogout()
        fmt.Fprintf(this.Response, "%s", response)
        break

    case "getId":
        sess := this.Session.SessionStart(this.Response, this.Request)
        ans := map[string]interface{}{"id": sess.Get("id")}
        res, err := json.Marshal(ans)
        utils.HandleErr("[Handle.Index] json.Marshal: ", err)
        fmt.Fprintf(this.Response, "%s", string(res))
        break

    case "select":
        tableName := data["table"].(string)
        fields := data["fields"].([]interface{})
        base := new(models.ModelManager)

        var model models.Entity
        var length int

        switch tableName {
        case "Users":
            length = len(base.Users().UserColumns)
            model = base.Users().Entity
            break
        case "Contests":
            length = len(base.Contests().Columns)
            model = base.Contests().Entity
            break
        }
        p := make([]string, length)
        j := 0
        for i, v := range fields {
            if v != nil {
                p[i] = v.(string)
                j++
            }
        }
        pp := make([]string, j)
        copy(pp[:], p[:j])
        answer := model.Select(nil, pp...)
        response, err := json.Marshal(answer)
        utils.HandleErr("[HandleLogin] json.Marshal: ", err)
        fmt.Fprintf(this.Response, "%s", response)

    case "update":
        tableName := data["table"].(string)
        base := new(models.ModelManager)

        var model models.Entity
        var length int

        switch tableName {
        case "Users":
            length = len(base.Users().UserColumns)
            model = base.Users().Entity
            break
        case "Contests":
            length = len(base.Contests().Columns)
            model = base.Contests().Entity
            break
        }

        d := data["data"].(map[string](interface{}))
        fields := utils.ArrayInterfaceToString(d["fields"].([]interface{}), length-1)
        values := d["userData"].([]interface{})

        model.Update(fields, values, fmt.Sprintf("id=%s", data["id"]))

        response, err := json.Marshal(map[string]interface{}{"result": "ok"})
        utils.HandleErr("[HandleLogin] json.Marshal: ", err)
        fmt.Fprintf(this.Response, "%s", response)
        break

    case "editBlank":
        id := data["id"].(string)
        inf := data["data"].([]interface{})

        columns := connect.DBGetColumnNames("blank_" + id)
        fmt.Println(columns)

        used := make(map[string][]int, len(columns))
        for i, v := range columns {
            used[v] = make([]int, 2)
            used[v][0] = 0
            used[v][1] = i
        }
        used["id"][0] = 1

        n := len(columns)

        db := connect.DBConnect()
        defer connect.DBClose(db)

        base := new(models.ModelManager)
        blanks := base.Blanks()

        for i := 0; i < len(inf); i = i + 4 {
            colNameObj := inf[i].(map[string]interface{})
            colNameDBObj := inf[i+1].(map[string]interface{})
            colTypeObj := inf[i+2].(map[string]interface{})
            colLenObj := inf[i+3].(map[string]interface{})

            var type_ string
            switch colTypeObj["value"].(string) {
            case "input":
                type_ = "varchar(" + colLenObj["value"].(string) + ")"
                break
            case "select":
                type_ = "varchar(" + colLenObj["value"].(string) + ")[]"
                break
            case "date":
                type_ = colLenObj["value"].(string)
                break
            }

            if utils.IsExist(columns, colNameDBObj["value"].(string)) == false {
                query := "ALTER TABLE blank_" + id + " ADD COLUMN " + colNameDBObj["value"].(string) + " " + type_
                stmt, err := db.Prepare(query)
                defer connect.DBClose(stmt)
                utils.HandleErr("[Handler.Index->editBlank] Prepare: ", err)
                _, err = stmt.Exec()
                utils.HandleErr("[Handler.Index->editBlank] Exec: ", err)

                blanks.Update([]string{"columns[" + strconv.Itoa(n) + "]"}, []interface{}{colNameDBObj["value"].(string)}, fmt.Sprintf("contest_id=%s", id))
                blanks.Update([]string{"colNames[" + strconv.Itoa(n) + "]"}, []interface{}{colNameObj["value"].(string)}, fmt.Sprintf("contest_id=%s", id))
                blanks.Update([]string{"types[" + strconv.Itoa(n) + "]"}, []interface{}{type_}, fmt.Sprintf("contest_id=%s", id))
                n++
            } else {
                used[colNameDBObj["value"].(string)][0] = 1
                query := "ALTER TABLE blank_" + id + " ALTER COLUMN " + colNameDBObj["value"].(string) + " TYPE " + type_
                stmt, err := db.Prepare(query)
                defer connect.DBClose(stmt)
                utils.HandleErr("[Handler.Index->editBlank] Prepare: ", err)
                _, err = stmt.Exec()
                utils.HandleErr("[Handler.Index->editBlank] Exec: ", err)
                blanks.Update([]string{"types[" + strconv.Itoa(used[colNameDBObj["value"].(string)][1]) + "]"}, []interface{}{type_}, fmt.Sprintf("contest_id=%s", id))
            }
        }

        for i, v := range used {
            if v[0] == 0 {
                query := "UPDATE blanks SET columns = "
                query += "(SELECT del_elem_by_index ("
                query += "(SELECT columns FROM blanks WHERE contest_id =" + id + "),"
                query += strconv.Itoa(v[1]) + "))"
                query += "WHERE contest_id =" + id

                stmt, err := db.Prepare(query)
                defer connect.DBClose(stmt)
                utils.HandleErr("[Handler.Index->editBlank] Prepare: ", err)
                _, err = stmt.Exec()
                utils.HandleErr("[Handler.Index->editBlank] Exec: ", err)

                query = "UPDATE blanks SET colnames = "
                query += "(SELECT del_elem_by_index ("
                query += "(SELECT colnames FROM blanks WHERE contest_id =" + id + "),"
                query += strconv.Itoa(v[1]) + "))"
                query += "WHERE contest_id =" + id

                stmt, err = db.Prepare(query)
                defer connect.DBClose(stmt)
                utils.HandleErr("[Handler.Index->editBlank] Prepare: ", err)
                _, err = stmt.Exec()
                utils.HandleErr("[Handler.Index->editBlank] Exec: ", err)

                query = "UPDATE blanks SET types = "
                query += "(SELECT del_elem_by_index ("
                query += "(SELECT types FROM blanks WHERE contest_id =" + id + "),"
                query += strconv.Itoa(v[1]) + "))"
                query += "WHERE contest_id =" + id

                stmt, err = db.Prepare(query)
                defer connect.DBClose(stmt)
                utils.HandleErr("[Handler.Index->editBlank] Prepare: ", err)
                _, err = stmt.Exec()
                utils.HandleErr("[Handler.Index->editBlank] Exec: ", err)

                query = "ALTER TABLE blank_" + id + " DROP COLUMN IF EXISTS " + i
                stmt, err = db.Prepare(query)
                defer connect.DBClose(stmt)
                utils.HandleErr("[Handler.Index->editBlank] Prepare: ", err)
                _, err = stmt.Exec()
                utils.HandleErr("[Handler.Index->editBlank] Exec: ", err)
            }
        }

        query := "UPDATE blanks SET columns = "
        query += "(SELECT del_null_from_arr ("
        query += "(SELECT columns FROM blanks WHERE contest_id =" + id + ")))"
        query += "WHERE contest_id =" + id

        stmt, err := db.Prepare(query)
        defer connect.DBClose(stmt)
        utils.HandleErr("[Handler.Index->editBlank] Prepare: ", err)
        _, err = stmt.Exec()
        utils.HandleErr("[Handler.Index->editBlank] Exec: ", err)

        query = "UPDATE blanks SET colnames = "
        query += "(SELECT del_null_from_arr ("
        query += "(SELECT colnames FROM blanks WHERE contest_id =" + id + ")))"
        query += "WHERE contest_id =" + id

        stmt, err = db.Prepare(query)
        defer connect.DBClose(stmt)
        utils.HandleErr("[Handler.Index->editBlank] Prepare: ", err)
        _, err = stmt.Exec()
        utils.HandleErr("[Handler.Index->editBlank] Exec: ", err)

        query = "UPDATE blanks SET types = "
        query += "(SELECT del_null_from_arr ("
        query += "(SELECT types FROM blanks WHERE contest_id =" + id + ")))"
        query += "WHERE contest_id =" + id

        stmt, err = db.Prepare(query)
        defer connect.DBClose(stmt)
        utils.HandleErr("[Handler.Index->editBlank] Prepare: ", err)
        _, err = stmt.Exec()
        utils.HandleErr("[Handler.Index->editBlank] Exec: ", err)
    }
}
