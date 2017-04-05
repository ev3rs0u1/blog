package admin

import (
    "blog/helpers"
    "blog/models"
    "net/http"

    "gopkg.in/macaron.v1"
)

func NewPost(ctx *macaron.Context) {
    ctx.Data["PageTitle"] = "New Post"
    ctx.Data["New"] = true
    ctx.HTML(http.StatusOK, "post.edit")
}

func ListPost(ctx *macaron.Context) {
    ctx.Data["PageTitle"] = "Posts"
    ctx.HTML(http.StatusOK, "post.list")
}

func EditPost(ctx *macaron.Context) {
    id := ctx.Params(":id")
    post, err := models.GetPostById(id)
    if err != nil {
        ctx.HTML(http.StatusNotFound, "errors.404")
    } else {
        ctx.Data["PageTitle"] = "Edit Post"
        ctx.Data["Post"] = post
        ctx.HTML(http.StatusOK, "post.edit")
    }
}

func SavePost(ctx *macaron.Context) {
    id := ctx.QueryTrim("id")
    title := ctx.QueryTrim("title")
    content := ctx.QueryTrim("content")
    created := helpers.ParseTime(ctx.QueryTrim("created"))
    post, _ := models.GetPostById(id)
    post.Title = title
    post.Content = content
    post.Created = created
    post.Save()
    ctx.Redirect(ctx.URLFor("post.list"), http.StatusSeeOther)
}

func DeletePost(ctx *macaron.Context) {
    ids := ctx.QueryStrings("id")
    models.DeletePosts(ids)
    ctx.Redirect(ctx.URLFor("post.list"), http.StatusSeeOther)
}
