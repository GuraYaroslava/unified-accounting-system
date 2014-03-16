package controllers

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "github.com/nu7hatch/gouuid"
    "github.com/uas/connect"
    "github.com/uas/utils"
    "regexp"
    "strconv"
)

func (c *BaseController) Handler() *Handler {
    return new(Handler)
}

type Handler struct {
    Controller
}

func MatchRegexp(pattern, str string) bool {
    result, _ := regexp.MatchString(pattern, str)
    return result
}

func (this *Handler) Index() {
    var (
        request  interface{}
        response string
    )
    this.Response.Header().Set("Access-Control-Allow-Origin", "*")
    this.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    this.Response.Header().Set("Content-type", "application/json")

    decoder := json.NewDecoder(this.Request.Body)
    err := decoder.Decode(&request)
    utils.HandleErr("[Handler] Decode :", err)
    data := request.(map[string]interface{})

    if data["action"] == "logout" {
        response = HandleLogout(data["sid"].(string))

    } else {
        login, password := data["login"].(string), data["password"].(string)

        switch data["action"] {
        case "login":
            response = HandleLogin(login, password)
        case "register":
            response = HandleRegister(login, password)
        }
    }

    fmt.Fprintf(this.Response, "%s", response)
}

func HandleRegister(login string, password string) string {
    result := map[string]string{"result": "ok"}

    passHasInvalidChars := false
    for i := 0; i < len(password); i++ {
        if strconv.IsPrint(rune(password[i])) == false {
            passHasInvalidChars = true
            break
        }
    }

    if IsExist(login) {
        result["result"] = "loginExists"

    } else if !MatchRegexp("^[a-zA-Z0-9]{2,36}$", login) {
        result["result"] = "badLogin"

    } else if !MatchRegexp("^.{6,36}$", password) && !passHasInvalidChars {
        result["result"] = "badPassword"

    } else {
        db := connect.DBConnect()
        query, err := db.Prepare("INSERT INTO users(login, password) VALUES($1, $2)")
        utils.HandleErr("[HandleRegister] Prepare error :", err)
        defer query.Close()
        _, err = query.Exec(login, password)
        utils.HandleErr("[HandleRegister] Query error :", err)
    }

    response, _ := json.Marshal(result)
    return string(response)
}

func IsExist(login string) bool {
    db := connect.DBConnect()
    var where string = "login = $1"
    //fmt.Println(connect.DBSelect("users", where, "id"))
    stmt, err := db.Prepare(connect.DBSelect("users", where, "id"))
    utils.HandleErr("[IsExist] Prepare: ", err)
    defer connect.DBClose(db, stmt)
    return stmt.QueryRow(login).Scan() != sql.ErrNoRows
}

func HandleLogin(login, password string) string {
    result := map[string]interface{}{"result": "invalidCredentials"}
    if IsExist(login) {
        db := connect.DBConnect()
        u4, _ := uuid.NewV4()
        stmt, err := db.Prepare("INSERT INTO sessions (login, sid) VALUES ($1, $2)")
        utils.HandleErr("[HandleLogin] Prepare: ", err)
        defer connect.DBClose(db, stmt)
        _, err = stmt.Exec(login, u4.String())
        utils.HandleErr("[HandleLogin] Insert into sessions: ", err)
        result["sid"] = u4.String()
        result["result"] = "ok"
    }
    response, err := json.Marshal(result)
    utils.HandleErr("[HandleLogin] json.Marshal: ", err)
    return string(response)
}

func HandleLogout(u4 string) string {
    result := map[string]string{"result": "ok"}
    db := connect.DBConnect()
    stmt, err := db.Prepare("DELETE FROM sessions WHERE sid = $1")
    utils.HandleErr("[HandleLogout] Prepare: ", err)
    defer connect.DBClose(db, stmt)
    rows, err := stmt.Exec(u4)
    utils.HandleErr("[HandleLogout] Prepare: ", err)
    if amount, _ := rows.RowsAffected(); amount != 1 {
        result["result"] = "badSid"
    }
    response, _ := json.Marshal(result)
    return string(response)
}
