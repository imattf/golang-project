package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

type tweet struct {
	Name     string
	Username string
	Message  string
}

func main() {
	fmt.Println("Doing crawl stuff...")

	c := colly.NewCollector()

	messages := []tweet{}

	// Find and visit all links
	// c.OnHTML("a", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })

	// Find all the tweets
	//c.OnHTML(".tweet", func(e *colly.HTMLElement) {
	c.OnHTML(".css-1dbjc4n", func(e *colly.HTMLElement) {
		messages = append(messages, tweet{
			Name:     e.ChildText(".css-901oao .css-16my406 r-1qd0xha r-ad9z0x r-bcqeeo r-qvutc0"),
			Username: e.ChildText(".css-901oao"),
			Message:  e.ChildText(".css-16my406"),
		})
		fmt.Println("Doing tweet stuff...")
	})

	fmt.Println(messages)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//c.Visit("http://go-colly.org/")
	err := c.Visit("https://twitter.com/Todd_McLeod/status/1169751640926146560")
	if err != nil {
		panic(err)
	}

	bs, err := json.MarshalIndent(messages, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
	fmt.Println("Number of tweets:", len(messages))
}
