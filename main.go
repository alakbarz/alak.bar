package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"regexp"

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
	m.Get("/blog", blogHandler)
	m.Get("/pics", picsHandler)
	m.Get("/report", reportHandler)
	m.Post("/report", csrf.Validate, binding.Bind(reportForm{}), reportHandlerPOST)
	m.Get("/thankyou", thankyouHandler)
	m.Get("/emailerror", emailerrorHandler)
	m.Get("/credits", creditsHandler)
	m.Get("/links", linksHandler)
	m.Get("/traffic", trafficHandler)

	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

func homeHandler(ctx *macaron.Context, x csrf.CSRF) {
	ctx.Data["Title"] = "Home"
	ctx.Data["csrf_token"] = x.GetToken()
	ctx.HTML(http.StatusOK, "index")
	getDescription()
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
	ctx.HTML(http.StatusOK, "links")
}

func trafficHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Traffic"
	ctx.HTML(http.StatusOK, "traffic")
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

func getDescription() {
	// Make HTTP request
	response, err := http.Get("https://www.humaidq.ae")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read response data in to memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	// <meta name="description" content="An Emirati who likes writing open-source software." />

	// Create a regular expression to find comments
	re := regexp.MustCompile("content=")
	comments := re.FindAllString(string(body), -1)
	if comments == nil {
		fmt.Println("No matches.")
	} else {
		for _, comment := range comments {
			fmt.Println(comment)
		}
	}
}

// func getDescription() {
// 	// Make HTTP request
// 	response, err := http.Get("https://www.humaidq.ae")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer response.Body.Close()

// 	// Create a goquery document from the HTTP response
// 	document, err := goquery.NewDocumentFromReader(response.Body)
// 	if err != nil {
// 		log.Fatal("Error loading HTTP response body. ", err)
// 	}

// 	// Find all links and process them with the function
// 	// defined earlier
// 	document.Find("meta").Each(processElement)
// }

// func processElement(index int, element *goquery.Selection) {
// 	// See if the href attribute exists on the element
// 	desc, exists := element.Attr("content")
// 	if exists {
// 		fmt.Println(desc)
// 	}
// }
