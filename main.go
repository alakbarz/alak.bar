package main

import (
    "log"
    "net/http"
    "gopkg.in/macaron.v1"
)

func main() {
    m := macaron.Classic()
    m.Use(macaron.Renderer())

    m.Get("/", homeHandler)
    m.Get("/about", aboutHandler)

    log.Println("Server is running...")
    log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

func homeHandler(ctx *macaron.Context) {
    ctx.Data["Title"] = "Alakbar"
    ctx.HTML(http.StatusOK, "index")
}

func aboutHandler(ctx *macaron.Context) {
    ctx.Data["Title"] = "About"
    ctx.HTML(http.StatusOK, "about")
}