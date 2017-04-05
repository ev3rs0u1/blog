package controllers

import (
    "net/http"

    "gopkg.in/macaron.v1"
)

func Index(ctx *macaron.Context) {
    ctx.Data["PageTitle"] = "Home"
    ctx.HTML(http.StatusOK, "home")
}
