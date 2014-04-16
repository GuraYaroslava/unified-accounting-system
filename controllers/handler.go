package controllers

import (
    "encoding/json"
    "fmt"
    "github.com/uas/models"
    "github.com/uas/utils"
    "html/template"
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
    case "home":
        tmp, err := template.ParseFiles("view/index.html", "view/header.html", "view/footer.html")
        utils.HandleErr("[Handler.Index] ParseFiles: ", err)
        err = tmp.ExecuteTemplate(this.Response, "index", nil)
        utils.HandleErr("[Handler.Index] ExecuteTemplate: ", err)
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
        //fmt.Println("select data: ", answer)
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
    }
}
