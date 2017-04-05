package routers

import (
    "blog/helpers"
    "blog/models"
    "html/template"
    "log"
    "net/http"

    "github.com/go-macaron/gzip"
    "github.com/go-macaron/session"
    "gopkg.in/macaron.v1"
)

var (
    router *macaron.Macaron
    conf   = helpers.GetConfig()
)

type Router struct {
    *macaron.Macaron
}

func New() *Router {
    new := func() *macaron.Macaron {
        initServer()
        iniStatic()
        initTmpl()
        initSession()
        initRouter()
        return router
    }
    return &Router{new()}
}

func (router *Router) ListenAndServe() {
    if !conf.Debug {
        macaron.Env = "production"
    }
    log.Printf("Server is running on %s... (%s)\n", conf.Addr, macaron.Env)
    http.ListenAndServe(conf.Addr, router)
}

func initServer() {
    router = macaron.New()
    router.Use(macaron.Logger())
    router.Use(macaron.Recovery())
}

func iniStatic() {
    router.Use(gzip.Gziper(gzip.Options{CompressionLevel: 2}))
    router.Use(macaron.Static(conf.Public))
}

func initTmpl() {
    router.Use(macaron.Renderer(macaron.RenderOptions{
        Directory: "views",
        Funcs: []template.FuncMap{map[string]interface{}{
            "URL":         router.URLFor,
            "User":        models.GetUser,
            "AllPosts":    models.GetAllPosts,
            "RecentPosts": models.GetRecentPosts,
            "PostsCount":  models.GetPostsCount,
            "DateTime":    helpers.DateTime,
            "Markdown":    helpers.Markdown,
            "HTML":        func(x string) interface{} { return template.HTML(x) },
        }},
    }))
}

func initSession() {
    router.Use(session.Sessioner(session.Options{CookieName: "sid"}))
}
