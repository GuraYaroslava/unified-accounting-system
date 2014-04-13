package controllers

import (
    "crypto/md5"
    "database/sql"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "github.com/uas/connect"
    "github.com/uas/utils"
    "regexp"
    "strconv"
    "time"
)

func MatchRegexp(pattern, str string) bool {
    result, _ := regexp.MatchString(pattern, str)
    return result
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func IsExist(login string) (bool, string, string, string) {
    db := connect.DBConnect()
    var id, hash, salt string
    query := connect.DBSelect("users", []string{"login"}, "id", "password", "salt")
    stmt, err := db.Prepare(query)
    utils.HandleErr("[IsExist] Prepare: ", err)
    defer connect.DBClose(db, stmt)
    err = stmt.QueryRow(login).Scan(&id, &hash, &salt)
    return err != sql.ErrNoRows, id, hash, salt
}

func (this *Handler) HandleRegister(login string, password string) string {
    result := map[string]string{"result": "ok"}
    salt := time.Now().Unix()
    fmt.Println("register salt: ", salt)
    hash := GetMD5Hash(password + strconv.Itoa(int(salt)))
    passHasInvalidChars := false
    for i := 0; i < len(password); i++ {
        if strconv.IsPrint(rune(password[i])) == false {
            passHasInvalidChars = true
            break
        }
    }
    isExist, _, _, _ := IsExist(login)
    if isExist == true {
        result["result"] = "loginExists"
    } else if !MatchRegexp("^[a-zA-Z0-9]{2,36}$", login) {
        result["result"] = "badLogin"
    } else if !MatchRegexp("^.{6,36}$", password) && !passHasInvalidChars {
        result["result"] = "badPassword"
    } else {
        db := connect.DBConnect()
        query := connect.DBInsert("users", []string{"login", "password", "salt"})
        stmt, err := db.Prepare(query)
        utils.HandleErr("[HandleRegister] Prepare error :", err)
        defer stmt.Close()
        _, err = stmt.Exec(login, hash, salt)
        utils.HandleErr("[HandleRegister] Query error :", err)
    }
    response, err := json.Marshal(result)
    utils.HandleErr("[HandleRegister] json.Marshal: ", err)
    return string(response)
}

func (this *Handler) HandleLogin(login, password string) string {

    result := map[string]interface{}{"result": "invalidCredentials"}
    isExist, id, hash, salt := IsExist(login)
    fmt.Println("login salt: ", salt)
    if isExist && hash == GetMD5Hash(password+salt) {

        sess := this.Session.SessionStart(this.Response, this.Request)
        sess.Set("createTime", time.Now().Unix())
        sess.Set("login", login)
        sess.Set("id", id)
        createTime := sess.Get("createTime")
        if createTime == nil {
            sess.Set("createTime", time.Now().Unix())
        } else if createTime.(int64)+this.Session.Maxlifetime < time.Now().Unix() {
            this.Session.GC()
        }
        result["id"] = id
        result["result"] = "ok"
    }
    response, err := json.Marshal(result)
    utils.HandleErr("[HandleLogin] json.Marshal: ", err)
    return string(response)
}

func (this *Handler) HandleLogout() string {
    result := map[string]string{"result": "ok"}
    this.Session.SessionDestroy(this.Response, this.Request)
    fmt.Println("session destroy", this.Session)
    response, err := json.Marshal(result)
    utils.HandleErr("[HandleLogout] json.Marshal: ", err)
    return string(response)
}
