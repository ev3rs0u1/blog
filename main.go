package main

import (
    "blog/routers"
)

func main() {
    app := routers.New()
    app.ListenAndServe()
}
