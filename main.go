package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/go-macaron/binding"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"github.com/gomarkdown/markdown"
	"github.com/metakeule/fmtdate"
	"github.com/robfig/cron"
	"golang.org/x/net/html"
	"gopkg.in/macaron.v1"
)

type link struct {
	Name        string
	Description string
	FileName    string
	ShortURL    string
	URL         string
}

var linksArr = []link{
	{Name: "Akilan", Description: "Not boring developer", FileName: "akilan.jpg", URL: "https://akilan.io", ShortURL: "akilan.io"},
	{Name: "Amaan", Description: "", FileName: "amaan.jpg", URL: "https://amaanakram.tech", ShortURL: "amaanakram.tech"},
	{Name: "Abdel", Description: "", FileName: "abdel.jpeg", URL: "https://elkabbany.xyz", ShortURL: "elkabbany.xyz"},
	{Name: "Euan", Description: "", FileName: "euan.jpg", URL: "https://euangordon.me", ShortURL: "euangordon.me"},
	{Name: "Humaid", Description: "", FileName: "humaid.jpg", URL: "https://humaidq.ae", ShortURL: "humaidq.ae"},
	{Name: "Hutchie", Description: "", FileName: "hutchie.png", URL: "https://hutchie.scot", ShortURL: "hutchie.scot"},
	{Name: "ReamSystems", Description: "A small tech startup company based in Scotland, that specializes in data analysis and web services.", FileName: "reamsystems.png", URL: "https://ream.systems", ShortURL: "ream.systems"},
	{Name: "Rikesh", Description: "Aspiring developer and designer", FileName: "rikesh.jpeg", URL: "http://rikeshmm.com", ShortURL: "rikeshmm.com"},
	{Name: "Rory", Description: "A collection of my projects and experiences", FileName: "rory.jpg", URL: "http://rorydobson.com/", ShortURL: "rorydobson.com"},
	{Name: "Ruaridh", Description: "A website showcasing who I am and some of my achievements over the last 3 years", FileName: "ruaridh.jpg", URL: "https://ruaridhmollica.com/", ShortURL: "ruaridhmollica.com"},
	{Name: "Shiva", Description: "", URL: "https://shiva-m.com/", FileName: "shiva.jpg", ShortURL: "shiva-m.com"},
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

type byDate []post

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner())
	m.Use(csrf.Csrfer())

	c := cron.New()
	c.AddFunc("@daily", func() { go updateDescriptions() })

	m.Get("/", homeHandler)
	m.Post("/", csrf.Validate, binding.Bind(contactForm{}), homeHandlerPOST)
	m.Get("/projects", projectsHandler)
	m.Get("/projects/:name", projectsFileHandler)
	m.Get("/blog", blogHandler)
	m.Get("/blog/:name", blogFileHandler)
	m.Get("/pics", picsHandler)
	m.Get("/report", reportHandler)
	m.Get("/thankyou", thankyouHandler)
	m.Get("/emailerror", emailerrorHandler)
	m.Get("/credits", creditsHandler)
	m.Get("/links", linksHandler)
	m.Get("/traffic", trafficHandler)
	m.Get("/alakbot", alakbotHandler)

	m.Post("/report", csrf.Validate, binding.Bind(reportForm{}), reportHandlerPOST)

	m.NotFound(func(ctx *macaron.Context) {
		ctx.Data["Title"] = "Not Found"
		ctx.HTML(http.StatusNotFound, "404")
	})

	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

func getPosts(directory string) {
	files, _ := ioutil.ReadDir(directory)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileContents, _ := ioutil.ReadFile(directory + file.Name())
		header := strings.Split(string(fileContents), "---")
		lines := strings.Split(header[1], "\n")

		var fields post
		fields.Markdown = header[2]

		html := markdown.ToHTML([]byte(header[2]), nil, nil)
		fields.HTML = string(html)

		filenameSplit := strings.Split(file.Name(), ".")
		fields.FileName = filenameSplit[0]

		for _, line := range lines {
			pair := strings.Split(line, ": ")

			if len(pair) == 2 {
				switch pair[0] {
				case "title":
					fields.Name = pair[1]
				case "description":
					fields.Description = pair[1]
				case "date":
					date, _ := fmtdate.Parse("YYYY-MM-DD", pair[1])

					dateFormatted := date.Format("January 2, 2006")
					fields.Date = date
					fields.DateFormatted = dateFormatted
				case "colourlight":
					fields.ColourLight = pair[1]
				case "colourdark":
					fields.ColourDark = pair[1]
				case "tags":
					// tags := strings.Split(pair[1], ",")
					fields.Tags = pair[1]
				}
			}
		}

		switch directory {
		case "public/projects/":
			projectsArr = append(projectsArr, fields)
			break
		case "public/blog/":
			blogsArr = append(blogsArr, fields)
		}
	}
}

func (d byDate) Len() int {
	return len(d)
}

func (d byDate) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d byDate) Less(i, j int) bool {
	return d[i].Date.After(d[j].Date)
}

func homeHandler(ctx *macaron.Context, x csrf.CSRF) {
	ctx.Data["Title"] = "Home"
	ctx.Data["csrf_token"] = x.GetToken()
	ctx.HTML(http.StatusOK, "index")
}

func homeHandlerPOST(ctx *macaron.Context, form contactForm) {
	form.Name = ctx.Query("name")
	form.Description = ctx.Query("description")

	if form.Name == "" {
		form.Name = "[NO NAME PROVIDED]"
	}

	form.Description = "Message From: " + form.Name + "\n\n" + "Description: " + form.Description
	sendMail("Message From: "+form.Name, form.Description, ctx)
}

func projectsHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Projects"
	projectsArr = nil
	getPosts("public/projects/")
	sort.Sort(byDate(projectsArr))
	ctx.Data["Projects"] = projectsArr
	ctx.HTML(http.StatusOK, "projects")
}

func projectsFileHandler(ctx *macaron.Context) {
	name := ctx.Params("name")

	for _, post := range projectsArr {
		if post.FileName == name {
			ctx.Data["HTML"] = template.HTML(post.HTML)
			ctx.Data["Name"] = post.FileName
			ctx.Data["ColourLight"] = post.ColourLight
			ctx.Data["ColourDark"] = post.ColourDark
			ctx.Data["Title"] = strings.Title(post.FileName)
		}
	}
	ctx.HTML(http.StatusOK, "post")
}

func blogHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Blog"
	blogsArr = nil
	getPosts("public/blog/")
	sort.Sort(byDate(blogsArr))
	ctx.Data["Blog"] = blogsArr
	ctx.HTML(http.StatusOK, "blog")
}

func blogFileHandler(ctx *macaron.Context) {
	name := ctx.Params("name")

	for _, post := range blogsArr {
		if post.FileName == name {
			ctx.Data["HTML"] = template.HTML(post.HTML)
			ctx.Data["Name"] = post.FileName
			ctx.Data["ColourLight"] = post.ColourLight
			ctx.Data["ColourDark"] = post.ColourDark
			ctx.Data["Title"] = strings.Title(post.FileName)
		}
	}
	ctx.HTML(http.StatusOK, "post")
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

func reportHandlerPOST(ctx *macaron.Context, form reportForm) {
	form.Category = "Website Report: " + ctx.Query("category")
	form.Email = ctx.Query("email")
	form.Description = ctx.Query("description")

	if form.Email == "" {
		form.Email = "[NO EMAIL PROVIDED]"
	}

	form.Description = "Category: " + form.Category + "\n\n" + "Email: " + form.Email + "\n\n" + "Description: " + form.Description
	sendMail(form.Category, form.Description, ctx)
}

func thankyouHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Thank You"
	ctx.HTML(http.StatusOK, "thankyou")
}

func emailerrorHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Report Error"
	ctx.HTML(http.StatusOK, "emailerror")
}

func creditsHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Credits"
	ctx.HTML(http.StatusOK, "credits")
}

func linksHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Links"
	ctx.Data["Links"] = linksArr
	ctx.HTML(http.StatusOK, "links")
}

func updateDescriptions() {
	for i, link := range linksArr {
		linksArr[i].Description = getDescription(link.URL)
	}
}

func trafficHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Traffic"
	ctx.HTML(http.StatusOK, "traffic")
}

func alakbotHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Alakbot"
	ctx.HTML(http.StatusOK, "alakbot")
}

func sendMail(subject, message string, ctx *macaron.Context) {
	from := "me@alak.bar"
	pass := os.Getenv("GPSSWRD")
	to := "me@alak.bar"
	msg := "From: " + from + "\n" + "To: " + to + "\n" + "Subject: " + subject + "\n\n" + message

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("[Gmail] SMTP ERR: %s", err)
		ctx.Redirect("/emailerror")
		return
	}

	log.Print("[Gmail] Email sent")
	ctx.Redirect("/thankyou")
}

func getDescription(url string) string {

	log.Println("Getting description for: " + url)

	// https://www.devdungeon.com/content/web-scraping-go
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", "Alakbot v0.1 (+http://alak.bar/alakbot)")

	// Make request
	response, err := client.Do(request)
	if err != nil {
		return ""
	}

	defer response.Body.Close()

	body := html.NewTokenizer(response.Body)

	for {
		tt := body.Next()
		if tt == html.ErrorToken {
			log.Println("HTML Error Token")
			return ""
		}

		switch tt {
		case html.ErrorToken:
			continue
		case html.StartTagToken, html.SelfClosingTagToken:
			tag, has := body.TagName()

			if !has || string(tag) != "meta" {
				continue
			} else {
				tagStr := string(tag)

				isDesc := false
				var description string

				for {
					key, val, has := body.TagAttr()

					keyStr := string(key)
					valStr := strings.ToLower(string(val))

					if tagStr == "meta" {
						if keyStr == "name" && valStr == "description" {
							isDesc = true
						} else if keyStr == "content" {
							description = valStr
						}
					}

					if !has {
						break
					}
				}

				if isDesc {
					return description
				}
			}
		}
	}
}
