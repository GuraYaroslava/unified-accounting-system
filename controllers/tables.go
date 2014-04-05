package controllers

import (
    "fmt"
    "github.com/uas/models"
    "github.com/uas/utils"
    "strings"
    "text/template"
)

type Model struct {
    Table     []interface{}
    TableName string
    Columns   []string
    ColNames  []string
    Caption   string
}

func (this *Handler) SelectById(tableName, id string) {
    fmt.Println("SelectById: ", tableName)
    base := new(models.ModelManager)
    var answer []interface{}
    switch tableName {
    case "Users":
        model := base.Users()
        answer = model.Select(map[string]interface{}{"id": id}, model.UserColumns...)
        //answer = append(answer, map[string]interface{}{"columns": model.UserColumns})
        //answer = append(answer, map[string]interface{}{"colNames": model.UserColNames})
        tmp, err := template.ParseFiles("../uas/view/card.html", "../uas/view/header.html", "../uas/view/footer.html")
        utils.HandleErr("[Handler.SelectById] template.ParseFiles: ", err)
        err = tmp.ExecuteTemplate(this.Response, "card", Model{Table: answer, ColNames: model.UserColNames})
        utils.HandleErr("[Handler.SelectById] tmp.Execute: ", err)
        break
        /*case "Contests":
        model := base.Contests()
        answer = model.Select(map[string]interface{}{"id": id}, model.Columns...)
        break*/
    }
    /*response, err := json.Marshal(answer)
    utils.HandleErr("[GetUserData] json.Marshal: ", err)
    fmt.Fprintf(this.Response, "%s", string(response))*/
}

func (this *Handler) Select(tableName string) {
    fmt.Println("Select: ", tableName)
    base := new(models.ModelManager)
    var model models.Entity
    switch tableName {
    case "Users":
        model = base.Users().Entity
        break
    case "Contests":
        model = base.Contests().Entity
        break
    }
    answer := model.Select(nil, model.Columns...)
    tmp, err := template.ParseFiles("../uas/view/table.html", "../uas/view/header.html", "../uas/view/footer.html")
    utils.HandleErr("[Handler.Select] template.ParseFiles: ", err)
    err = tmp.ExecuteTemplate(this.Response, "edit", Model{Table: answer, TableName: model.TableName, ColNames: model.ColNames, Columns: model.Columns, Caption: model.Caption})
    utils.HandleErr("[Handler.Select] tmp.Execute: ", err)
}

func (this *Handler) Edit(tableName string) {
    fmt.Println("Edit: ", tableName)
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
    }

    params := make([]interface{}, len(model.Columns)-1)
    for i := 0; i < len(model.Columns)-1 && this.Request.FormValue(model.Columns[i+1]) != ""; i++ {
        fmt.Println(this.Request.FormValue(model.Columns[i+1]))
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
