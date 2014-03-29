package controllers

import (
    "encoding/json"
    "fmt"
    "github.com/uas/connect"
    "github.com/uas/models"
    "github.com/uas/utils"
    "strconv"
    "strings"
)

func GetUserData(id string) string {
    db := connect.DBConnect()
    var where string = "id = $1"
    var user = new(models.User)
    var base *models.ModelManager
    user.TableData = base.Users()
    query := connect.DBSelect("users", where, strings.Join(user.TableData.Columns, ", "))
    stmt, err := db.Prepare(query)
    utils.HandleErr("[GetUserData] Prepare: ", err)
    result := stmt.QueryRow(id)
    err = result.Scan(
        &user.Id,
        &user.Login,
        &user.Password,
        &user.Salt,
        &user.Sid,
        &user.FName,
        &user.LName,
        &user.PName,
        &user.EMail,
        &user.Phone,
        &user.Address)
    utils.HandleErr("[GetUserData] Scan: ", err)
    defer connect.DBClose(db, stmt)
    response, err := json.Marshal(user)
    utils.HandleErr("[GetUserData] json.Marshal: ", err)
    return string(response)
}

func UpdateUserData(id string, data map[string]interface{}) string {
    fmt.Println(data)
    result := map[string]string{"result": "ok"}
    i := 0
    var (
        fields = make([]string, len(data))
        params = make([]string, len(data))
        values = make([]string, len(data)+1)
    )
    for key, _ := range data {
        fields[i] = key
        params[i] = "$" + strconv.Itoa(i)
        values[i] = data[key].(string)
        i++
    }
    values[i] = id
    db := connect.DBConnect()
    query := connect.DBUpdate("users", fields, "id = $"+strconv.Itoa(len(data)+1))
    println(query)
    stmt, err := db.Prepare(query)
    utils.HandleErr("[UpdateUserData] Prepare: ", err)
    defer connect.DBClose(db, stmt)
    _, err = stmt.Query(values[0], values[1], values[2], values[3], values[4], values[5], values[6], values[7])
    utils.HandleErr("[UpdateUserData] Update: ", err)
    response, err := json.Marshal(result)
    utils.HandleErr("[UpdateUserData] json.Marshal: ", err)
    return string(response)
}
