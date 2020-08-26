package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

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
		log.Println(err)
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
							if valStr == "" {
								break
							} else {
								firstLetter := valStr[0:1]
								description = strings.Title(string(firstLetter)) + valStr[1:]
							}
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
