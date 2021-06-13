package main

import (
	"html/template"
	"net/http"
	"sort"
	"strings"

	"github.com/go-macaron/csrf"
	"gopkg.in/macaron.v1"
)

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

	postFound := false

	for _, post := range projectsArr {
		if post.FileName == name {
			ctx.Data["HTML"] = template.HTML(post.HTML)
			ctx.Data["Name"] = post.FileName
			ctx.Data["ColourLight"] = post.ColourLight
			ctx.Data["ColourDark"] = post.ColourDark
			ctx.Data["Title"] = strings.Title(post.FileName)
			ctx.Data["ProjDesc"] = post.Description
			postFound = true
		}
	}

	if postFound {
		ctx.HTML(http.StatusOK, "post")
	} else {
		ctx.Data["Title"] = "Not Found"
		ctx.HTML(http.StatusNotFound, "404")
	}
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

	postFound := false

	for _, post := range blogsArr {
		if post.FileName == name {
			ctx.Data["HTML"] = template.HTML(post.HTML)
			ctx.Data["Name"] = post.FileName
			ctx.Data["ColourLight"] = post.ColourLight
			ctx.Data["ColourDark"] = post.ColourDark
			ctx.Data["Title"] = strings.Title(post.FileName)
			ctx.Data["BlogDesc"] = post.Description
			postFound = true
		}
	}

	if postFound {
		ctx.HTML(http.StatusOK, "post")
	} else {
		ctx.Data["Title"] = "Not Found"
		ctx.HTML(http.StatusNotFound, "404")
	}
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

func trafficHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Traffic"
	ctx.HTML(http.StatusOK, "traffic")
}

func alakbotHandler(ctx *macaron.Context) {
	ctx.Data["Title"] = "Alakbot"
	ctx.HTML(http.StatusOK, "alakbot")
}
