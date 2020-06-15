package main

import (
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/go-macaron/binding"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

type reportForm struct {
	Category    string `form:"category" binding:"Required"`
	Email       string `form:"email"`
	Description string `form:"description" binding:"Required"`
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner())
	m.Use(csrf.Csrfer())

	m.Get("/", homeHandler)
	m.Get("/projects", projectsHandler)
	m.Get("/blog", blogHandler)
	m.Get("/pics", picsHandler)
	m.Get("/report", reportHandler)
	m.Post("/report", csrf.Validate, binding.Bind(reportForm{}), reportHandlerPOST)
	m.Get("/thankyou", thankyouHandler)
	m.Get("/reporterror", reporterrorHandler)
	m.Get("/credits", creditsHandler)

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

func reportHandler(ctx *macaron.Context, x csrf.CSRF) {
	ctx.Data["Title"] = "Report Issue"
	ctx.Data["csrf_token"] = x.GetToken()
	ctx.HTML(http.StatusOK, "report")
}

func thankyouHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Thank You"
	ctx.HTML(http.StatusOK, "thankyou")
}

func reporterrorHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Report Error"
	ctx.HTML(http.StatusOK, "reporterror")
}

func reportHandlerPOST(ctx *macaron.Context, form reportForm) {
	form.Category = ctx.Query("category")
	form.Email = ctx.Query("email")
	form.Description = ctx.Query("description")

	if form.Email == "" {
		form.Email = "[NO EMAIL PROVIDED]"
	}

	form.Description = "Category: " + form.Category + "\n" + "Email: " + form.Email + "\n" + "Description: " + form.Description

	from := "me@alak.bar"
	pass := os.Getenv("GPSSWRD")
	to := "me@alak.bar"
	msg := "From: " + from + "\n" + "To: " + to + "\n" + "Subject: Website Report: " + form.Category + "\n\n" + form.Description

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		ctx.Redirect("/reporterror")
		return
	} else {
		log.Print("Report email sent")
		ctx.Redirect("/thankyou")
	}
}

func creditsHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Credits"
	ctx.HTML(http.StatusOK, "credits")
}
