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
	m.Get("/projects", projectsHandler)
	m.Get("/blog", blogHandler)
	m.Get("/pics", picsHandler)

	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

func homeHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Home"
	ctx.HTML(http.StatusOK, "index")
}

func projectsHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Projects"
	ctx.HTML(http.StatusOK, "projects")
}

func blogHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Blog"
	ctx.HTML(http.StatusOK, "blog")
}

func picsHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Pictures"
	ctx.HTML(http.StatusOK, "pics")
}
