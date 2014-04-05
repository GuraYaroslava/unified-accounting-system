package controllers

import (
    "encoding/json"
    "fmt"
    "github.com/uas/utils"
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
        response = HandleRegister(login, password)
        fmt.Fprintf(this.Response, "%s", response)
        break
    case "login":
        login, password := data["login"].(string), data["password"].(string)
        response = HandleLogin(login, password)
        fmt.Fprintf(this.Response, "%s", response)
        break
    case "logout":
        //response = HandleLogout(data["sid"].(string))
        fmt.Fprintf(this.Response, "%s", response)
        break
    }
}
