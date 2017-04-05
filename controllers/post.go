package controllers

import (
    "blog/models"
    "net/http"

    "gopkg.in/macaron.v1"
)

func GetPost(ctx *macaron.Context) {
    id := ctx.Params(":id")
    post, err := models.GetPostById(id)
    if err != nil {
        ctx.HTML(http.StatusNotFound, "errors.404")
    } else {
        ctx.Data["PageTitle"] = post.Title
        ctx.Data["Post"] = post
        ctx.HTML(http.StatusOK, "post")
    }
}
