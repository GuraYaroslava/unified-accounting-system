package controllers

import (
    "github.com/uas/utils"
    "html/template"
)

func (c *BaseController) Index() *IndexController {
    return new(IndexController)
}

type IndexController struct {
    Controller
}

func (this *IndexController) Index() {
    this.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
    tmp, err := template.ParseFiles("view/index.html", "view/header.html", "view/footer.html")
    utils.HandleErr("[IndexController] ParseFiles: ", err)
    err = tmp.ExecuteTemplate(this.Response, "index", nil)
    utils.HandleErr("[IndexController] ExecuteTemplate: ", err)
}
