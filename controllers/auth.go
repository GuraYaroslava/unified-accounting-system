package controllers

import (
    "database/sql"
    "encoding/json"
    "github.com/nu7hatch/gouuid"
    "github.com/uas/connect"
    "github.com/uas/utils"
    "regexp"
    "strconv"
)

func MatchRegexp(pattern, str string) bool {
    result, _ := regexp.MatchString(pattern, str)
    return result
}

func IsExist(login string) (bool, string) {
    db := connect.DBConnect()
    var where string = "login = $1"
    var id string
    query := connect.DBSelect("users", where, "id")
    stmt, err := db.Prepare(query)
    utils.HandleErr("[IsExist] Prepare: ", err)
    defer connect.DBClose(db, stmt)
    err = stmt.QueryRow(login).Scan(&id)
    return err != sql.ErrNoRows, id
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
    isExist, _ := IsExist(login)
    if isExist == true {
        result["result"] = "loginExists"
    } else if !MatchRegexp("^[a-zA-Z0-9]{2,36}$", login) {
        result["result"] = "badLogin"
    } else if !MatchRegexp("^.{6,36}$", password) && !passHasInvalidChars {
        result["result"] = "badPassword"
    } else {
        db := connect.DBConnect()
        query := connect.DBInsert("users", []string{"login", "password"})
        stmt, err := db.Prepare(query)
        utils.HandleErr("[HandleRegister] Prepare error :", err)
        defer stmt.Close()
        _, err = stmt.Exec(login, password)
        utils.HandleErr("[HandleRegister] Query error :", err)
    }
    response, err := json.Marshal(result)
    utils.HandleErr("[HandleRegister] json.Marshal: ", err)
    return string(response)
}

func HandleLogin(login, password string) string {
    result := map[string]interface{}{"result": "invalidCredentials"}
    isExist, id := IsExist(login)
    if isExist {
        db := connect.DBConnect()
        u4, _ := uuid.NewV4()
        query := connect.DBUpdate("users", []string{"sid"}, "id = $2")
        stmt, err := db.Prepare(query)
        utils.HandleErr("[HandleLogin] Prepare: ", err)
        defer connect.DBClose(db, stmt)
        _, err = stmt.Exec(u4.String(), id)
        utils.HandleErr("[HandleLogin] Update sid: ", err)
        result["id"] = id
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
    query := connect.DBUpdate("users", []string{"sid"}, "sid = $2")
    stmt, err := db.Prepare(query)
    utils.HandleErr("[HandleLogout] Prepare: ", err)
    rows, err := stmt.Exec("", u4)
    defer connect.DBClose(db, stmt)
    utils.HandleErr("[HandleLogout] Prepare: ", err)
    if amount, _ := rows.RowsAffected(); amount != 1 {
        result["result"] = "badSid"
    }
    response, err := json.Marshal(result)
    utils.HandleErr("[HandleLogout] json.Marshal: ", err)
    return string(response)
}
