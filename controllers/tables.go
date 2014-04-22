package controllers

import (
    "fmt"
    "github.com/uas/models"
    "github.com/uas/utils"
    "strings"
    "text/template"
    "time"
)

type Model struct {
    Id        string
    TableName string
    Caption   string
    Table     []interface{}
    Columns   []interface{}
    ColNames  []interface{}
    Types     []interface{}
}

func (this *Handler) SelectById(tableName string) {
    sess := this.Session.SessionStart(this.Response, this.Request)
    createTime := sess.Get("createTime")
    life := this.Session.Maxlifetime
    if createTime == nil || createTime.(int64)+life < time.Now().Unix() {
        this.Session.SessionDestroy(this.Response, this.Request)
        fmt.Println("SelectById - Destroy")
        tmp, err := template.ParseFiles(
            "view/index.html",
            "view/header.html",
            "view/footer.html")
        utils.HandleErr("[handler.select] ParseFiles: ", err)
        err = tmp.ExecuteTemplate(this.Response, "index", nil)
        utils.HandleErr("[handler.select] ExecuteTemplate: ", err)
        return
    }
    sess.Set("createTime", time.Now().Unix())
    id := sess.Get("id")
    base := new(models.ModelManager)
    var answer []interface{}
    switch tableName {
    case "Users":
        model := base.Users()
        answer = model.Select(map[string]interface{}{"id": id}, model.UserColumns...)
        tmp, err := template.ParseFiles(
            "../uas/view/card.html",
            "../uas/view/header.html",
            "../uas/view/footer.html")
        utils.HandleErr("[Handler.SelectById] template.ParseFiles: ", err)
        err = tmp.ExecuteTemplate(this.Response, "card", Model{
            Table:    answer,
            ColNames: utils.ArrayStringToInterface(model.UserColNames),
            Columns:  utils.ArrayStringToInterface(model.UserColumns)})
        utils.HandleErr("[Handler.SelectById] tmp.Execute: ", err)
        break
    }
}

func (this *Handler) Select(tableName string) {
    sess := this.Session.SessionStart(this.Response, this.Request)
    createTime := sess.Get("createTime")
    life := this.Session.Maxlifetime
    if createTime == nil || createTime.(int64)+life < time.Now().Unix() {
        this.Session.SessionDestroy(this.Response, this.Request)
        fmt.Println("Select - Destroy")
        tmp, err := template.ParseFiles(
            "view/index.html",
            "view/header.html",
            "view/footer.html")
        utils.HandleErr("[handler.select] ParseFiles: ", err)
        err = tmp.ExecuteTemplate(this.Response, "index", nil)
        utils.HandleErr("[handler.select] ExecuteTemplate: ", err)
        return
    }
    sess.Set("createTime", time.Now().Unix())
    base := new(models.ModelManager)
    var model models.Entity
    switch tableName {
    case "Users":
        model = base.Users().Entity
        break
    case "Contests":
        model = base.Contests().Entity
        break
    case "Blanks":
        model = base.Blanks().Entity
        break
    }
    answer := model.Select(nil, model.Columns...)
    tmp, err := template.ParseFiles(
        "../uas/view/table.html",
        "../uas/view/header.html",
        "../uas/view/footer.html")
    utils.HandleErr("[Handler.Select] template.ParseFiles: ", err)
    err = tmp.ExecuteTemplate(this.Response, "edit", Model{
        Table:     answer,
        TableName: model.TableName,
        ColNames:  utils.ArrayStringToInterface(model.ColNames),
        Columns:   utils.ArrayStringToInterface(model.Columns),
        Caption:   model.Caption})
    utils.HandleErr("[Handler.Select] tmp.Execute: ", err)
}

func (this *Handler) Edit(tableName string) {
    sess := this.Session.SessionStart(this.Response, this.Request)
    createTime := sess.Get("createTime")
    life := this.Session.Maxlifetime
    if createTime.(int64)+life < time.Now().Unix() {
        this.Session.SessionDestroy(this.Response, this.Request)
        fmt.Println("Edit - Destroy")
        tmp, err := template.ParseFiles(
            "view/index.html",
            "view/header.html",
            "view/footer.html")
        utils.HandleErr("[handler.select] ParseFiles: ", err)
        err = tmp.ExecuteTemplate(this.Response, "index", nil)
        utils.HandleErr("[handler.select] ExecuteTemplate: ", err)
        return
    }
    sess.Set("createTime", time.Now().Unix())
    oper := this.Request.FormValue("oper")
    base := new(models.ModelManager)
    var model models.Entity
    switch tableName {
    case "Users":
        model = base.Users().Entity
        break
    case "Contests":
        model = base.Contests().Entity
        break
    case "Blanks":
        model = base.Blanks().Entity
        break
    }

    params := make([]interface{}, len(model.Columns)-1)
    for i := 0; i < len(model.Columns)-1 && this.Request.FormValue(model.Columns[i+1]) != ""; i++ {
        if model.Columns[i+1] == "date" {
            params[i] = this.Request.FormValue(model.Columns[i+1])[0:10]
        } else {
            params[i] = this.Request.FormValue(model.Columns[i+1])
        }
    }

    switch oper {
    case "edit":
        model.Update(model.Columns[1:], params, fmt.Sprintf("id=%s", this.Request.FormValue("id")))
        break
    case "add":
        model.Insert(model.Columns[1:], params)
        break
    case "del":
        ids := strings.Split(this.Request.FormValue("id"), ",")
        p := make([]interface{}, len(ids))
        for i, v := range ids {
            p[i] = interface{}(v)
        }
        model.Delete("id", p)
        break
    }
}
