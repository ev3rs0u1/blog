package admin

import (
    "blog/models"
    "net/http"

    "github.com/go-macaron/session"
    "gopkg.in/macaron.v1"
)

func Index(ctx *macaron.Context) {
    ctx.Data["PageTitle"] = "Index"
    ctx.HTML(http.StatusOK, "admin.index")
}

func LoginGet(ctx *macaron.Context) {
    ctx.Data["PageTitle"] = "Login"
    ctx.HTML(http.StatusOK, "admin.login")
}

func LoginPost(ctx *macaron.Context, sess session.Store, flash *session.Flash) {
    username := ctx.QueryTrim("username")
    password := ctx.QueryTrim("password")
    if len(username) == 0 || len(password) == 0 {
        flash.Error("请填写用户名和密码")
        ctx.Redirect(ctx.URLFor("admin.login"), http.StatusFound)
        return
    }

    user := models.ValidateAccount(username, password)
    if user != nil {
        sess.Set("uid", user.Id)
        ctx.Redirect(ctx.URLFor("admin.index"), http.StatusSeeOther)
    } else {
        flash.Error("用户名或密码错误")
        ctx.Redirect(ctx.URLFor("admin.login"), http.StatusFound)
    }
}

func Logout(ctx *macaron.Context, sess session.Store) {
    sess.Delete("uid")
    ctx.Redirect(ctx.URLFor("admin.login"), http.StatusFound)
}

func ProfileGet(ctx *macaron.Context) {
    ctx.Data["PageTitle"] = "Profile"
    ctx.HTML(http.StatusOK, "admin.profile")
}

func ProfilePost(ctx *macaron.Context, sess session.Store, flash *session.Flash) {
    user := new(models.User)
    author := ctx.QueryTrim("author")
    url := ctx.QueryTrim("url")
    profile := ctx.QueryTrim("profile")
    if len(profile) > 0 {
        user.Author = author
        user.Url = url
        user.Update()
        flash.Success("资料更新成功")
        ctx.Redirect(ctx.URLFor("admin.profile"), http.StatusSeeOther)
        return
    }

    nPassword := ctx.QueryTrim("npassword")
    cPassword := ctx.QueryTrim("cpassword")
    if len(nPassword) == 0 {
        flash.Error("新密码不能为空")
    } else if len(nPassword) < 6 {
        flash.Error("密码长度不能少于6个字符")
    } else if nPassword != cPassword {
        flash.Error("两次输入的密码不一致")
    } else {
        user.Password = models.EncryptPwd(nPassword)
        user.Update()
        flash.Success("密码修改成功")
    }
    ctx.Redirect(ctx.URLFor("admin.profile"), http.StatusSeeOther)
}
