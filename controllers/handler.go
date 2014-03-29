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
    if data["action"] == "logout" {
        response = HandleLogout(data["sid"].(string))
    } else if data["action"] == "identification" {
        response = GetUserData(data["id"].(string))
    } else if data["action"] == "updateUser" {
        userData := data["data"].(map[string]interface{})
        response = UpdateUserData(data["id"].(string), userData)
        fmt.Println(response)
    } else {
        login, password := data["login"].(string), data["password"].(string)
        if data["action"] == "login" {
            response = HandleLogin(login, password)
        }
        if data["action"] == "register" {
            response = HandleRegister(login, password)
        }
    }
    fmt.Fprintf(this.Response, "%s", response)
}
