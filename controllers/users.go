package controllers

import (
    "encoding/json"
    "fmt"
    "github.com/uas/connect"
    "github.com/uas/models"
    "github.com/uas/utils"
)

func GetUserData(id string) string {
    db := connect.DBConnect()
    var where string = "id = $1"
    var user = new(models.User)
    var base *models.ModelManager
    user.TableData = base.Users()
    query := connect.DBSelect("users", where, "id", "fname", "lname", "pname", "login", "salt", "hash", "email", "phone", "address")
    stmt, err := db.Prepare(query)
    utils.HandleErr("[GetUserData] Prepare: ", err)
    result := stmt.QueryRow(id)
    err = result.Scan(&user.Id, &user.FName, &user.LName, &user.PName, &user.Login, &user.Salt, &user.Hash, &user.EMail, &user.Phone, &user.Address)
    utils.HandleErr("[GetUserData] Scan: ", err)
    defer connect.DBClose(db, stmt)
    response, err := json.Marshal(user)
    fmt.Println(user)
    utils.HandleErr("[GetUserData] json.Marshal: ", err)
    return string(response)
}
