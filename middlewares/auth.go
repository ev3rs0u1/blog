package middlewares

import (
    "net/http"

    "github.com/go-macaron/session"
    "gopkg.in/macaron.v1"
)

func AuthLogin(ctx *macaron.Context, sess session.Store) {
    if uid := sess.Get("uid"); uid != nil {
        ctx.Redirect(ctx.URLFor("admin.index"), http.StatusSeeOther)
    } else {
        ctx.Next()
    }
}

func AuthPermissions(ctx *macaron.Context, sess session.Store) {
    if uid := sess.Get("uid"); uid != nil {
        ctx.Next()
    } else {
        ctx.Redirect(ctx.URLFor("admin.login"), http.StatusFound)
    }
}
