package controllers

import (
    "net/http"
)

type BaseController struct{}

type Controller struct {
    Request  *http.Request
    Response http.ResponseWriter
}
