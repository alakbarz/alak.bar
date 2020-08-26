package main

import (
	"io/ioutil"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/metakeule/fmtdate"
)

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

func updateDescriptions() {
	for i, link := range linksArr {
		linksArr[i].Description = getDescription(link.URL)
	}
}
