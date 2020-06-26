package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/go-macaron/binding"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"github.com/gomarkdown/markdown"
	"golang.org/x/net/html"
	"gopkg.in/macaron.v1"
)

type reportForm struct {
	Category    string `form:"category" binding:"Required"`
	Email       string `form:"email"`
	Description string `form:"description" binding:"Required"`
}

type contactForm struct {
	Name        string `form:"name"`
	Description string `form:"description" binding:"Required"`
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner())
	m.Use(csrf.Csrfer())

	m.Get("/", homeHandler)
	m.Post("/", csrf.Validate, binding.Bind(contactForm{}), homeHandlerPOST)
	m.Get("/projects", projectsHandler)
	m.Get("/projects/:name", projectsFileHandler)
	m.Get("/blog", blogHandler)
	m.Get("/pics", picsHandler)
	m.Get("/report", reportHandler)
	m.Get("/thankyou", thankyouHandler)
	m.Get("/emailerror", emailerrorHandler)
	m.Get("/credits", creditsHandler)
	m.Get("/links", linksHandler)
	m.Get("/traffic", trafficHandler)
	m.Get("/alakbot", alakbotHandler)

	m.Post("/report", csrf.Validate, binding.Bind(reportForm{}), reportHandlerPOST)

	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

func homeHandler(ctx *macaron.Context, x csrf.CSRF) {
	ctx.Data["Title"] = "Home"
	ctx.Data["csrf_token"] = x.GetToken()
	ctx.HTML(http.StatusOK, "index")
	// markdownToHTML()
}

func homeHandlerPOST(ctx *macaron.Context, form contactForm) {
	form.Name = ctx.Query("email")
	form.Description = ctx.Query("description")

	if form.Name == "" {
		form.Name = "[NO NAME PROVIDED]"
	}

	form.Description = "Message From: " + form.Name + "\n\n" + "Description: " + form.Description
	sendMail("Message From: "+form.Name, form.Description, ctx)
}

func projectsHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Projects"
	ctx.HTML(http.StatusOK, "projects")
}

func projectsFileHandler(ctx *macaron.Context) {

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
	ctx.Data["HumaidDesc"] = getDescription("https://www.humaidq.ae")
	ctx.HTML(http.StatusOK, "links")
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
		log.Fatal(err)
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
					valStr := string(val)

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

func markdownToHTML() {
	file := "public/projects/test.md"

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Println("Could not read file")
	}

	html := markdown.ToHTML(content, nil, nil)

	log.Print(string(html))
}
