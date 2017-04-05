package routers

import (
    "blog/controllers"
    "blog/controllers/admin"
    "blog/middlewares"
)

var (
    adminRoute = "/" + conf.Admin.Route
)

func initRouter() {
    router.Get("/", controllers.Index).Name("index")
    router.Get("/:id", controllers.GetPost).Name("post")

    router.Group(adminRoute, func() {
        router.Combo("/login").
            Get(admin.LoginGet).
            Post(admin.LoginPost).
            Name("admin.login")
    }, middlewares.AuthLogin)

    router.Group(adminRoute, func() {
        router.Get("/", admin.Index).Name("admin.index")
        router.Get("/logout", admin.Logout).Name("admin.logout")
        router.Combo("/profile").
            Get(admin.ProfileGet).
            Post(admin.ProfilePost).
            Name("admin.profile")

        router.Group("/post", func() {
            router.Get("/new", admin.NewPost).Name("post.new")
            router.Get("/list", admin.ListPost).Name("post.list")
            router.Get("/edit/:id", admin.EditPost).Name("post.edit")
            router.Post("/save", admin.SavePost).Name("post.save")
            router.Post("/delete", admin.DeletePost).Name("post.delete")
        })
    }, middlewares.AuthPermissions)

    router.NotFound(controllers.NotFound)
}
