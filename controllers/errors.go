package controllers

import (
    "net/http"

    "gopkg.in/macaron.v1"
)

func NotFound(ctx *macaron.Context) {
    ctx.HTML(http.StatusNotFound, "errors.404")
}
