package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-macaron/binding"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"github.com/robfig/cron"
	"gopkg.in/macaron.v1"
)

type byDate []post

type link struct {
	Name        string
	Description string
	FileName    string
	ShortURL    string
	URL         string
}

type reportForm struct {
	Category    string `form:"category" binding:"Required"`
	Email       string `form:"email"`
	Description string `form:"description" binding:"Required"`
}

type contactForm struct {
	Name        string `form:"name"`
	Description string `form:"description" binding:"Required"`
}

type post struct {
	ColourDark    string
	ColourLight   string
	Date          time.Time
	DateFormatted string
	Description   string
	FileName      string
	HTML          string
	Markdown      string
	Name          string
	Tags          string
}

var projectsArr []post
var blogsArr []post
var linksArr = []link{
	{Name: "Akilan", Description: "Not boring developer", FileName: "akilan.jpg", URL: "https://akilan.io", ShortURL: "akilan.io"},
	{Name: "Amaan", Description: "", FileName: "amaan.jpg", URL: "https://amaanakram.tech", ShortURL: "amaanakram.tech"},
	{Name: "Abdel", Description: "", FileName: "abdel.jpeg", URL: "https://elkabbany.xyz", ShortURL: "elkabbany.xyz"},
	{Name: "Euan", Description: "", FileName: "euan.jpg", URL: "https://euangordon.me", ShortURL: "euangordon.me"},
	{Name: "Humaid", Description: "", FileName: "humaid.jpg", URL: "https://humaidq.ae", ShortURL: "humaidq.ae"},
	{Name: "Hutchie", Description: "", FileName: "hutchie.png", URL: "https://hutchie.scot", ShortURL: "hutchie.scot"},
	{Name: "ReamSystems", Description: "", FileName: "reamsystems.png", URL: "https://ream.systems", ShortURL: "ream.systems"},
	{Name: "Rikesh", Description: "Aspiring developer and designer", FileName: "rikesh.jpeg", URL: "http://rikeshmm.com", ShortURL: "rikeshmm.com"},
	{Name: "Rory", Description: "A collection of my projects and experiences", FileName: "rory.jpg", URL: "http://rorydobson.com/", ShortURL: "rorydobson.com"},
	{Name: "Ruaridh", Description: "", FileName: "ruaridh.jpg", URL: "https://ruaridhmollica.com/", ShortURL: "ruaridhmollica.com"},
	{Name: "Shiva", Description: "", URL: "https://shiva-m.com/", FileName: "shiva.jpg", ShortURL: "shiva-m.com"},
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner())
	m.Use(csrf.Csrfer())
	c := cron.New()
	c.AddFunc("@daily", func() { go updateDescriptions() })
	c.Start()
	go updateDescriptions()
	m.Get("/", homeHandler)
	m.Post("/", csrf.Validate, binding.Bind(contactForm{}), homeHandlerPOST)
	m.Get("/projects", projectsHandler)
	m.Get("/projects/:name", projectsFileHandler)
	m.Get("/blog", blogHandler)
	m.Get("/blog/:name", blogFileHandler)
	m.Get("/pics", picsHandler)
	m.Get("/report", reportHandler)
	m.Post("/report", csrf.Validate, binding.Bind(reportForm{}), reportHandlerPOST)
	m.Get("/thankyou", thankyouHandler)
	m.Get("/emailerror", emailerrorHandler)
	m.Get("/credits", creditsHandler)
	m.Get("/links", linksHandler)
	m.Get("/traffic", trafficHandler)
	m.Get("/alakbot", alakbotHandler)
	m.NotFound(func(ctx *macaron.Context) {
		ctx.Data["Title"] = "Not Found"
		ctx.HTML(http.StatusNotFound, "404")
	})
	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}
