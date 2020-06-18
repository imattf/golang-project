// Hands-on exercise #2
// For this hands-on exercise:
// use gocolly http://go-colly.org/
// crawl https://www.reddit.com/r/golang/
// hint - take a look at: https://old.reddit.com/r/golang/

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type item struct {
	StoryURL  string
	Source    string
	comments  string
	CrawledAt time.Time
	Comments  string
	Title     string
}

func main() {
	stories := []item{}
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: reddit.com
		colly.AllowedDomains("old.reddit.com"),
		colly.Async(true),
	)

	// On every a element which has .top-matter attribute call callback
	// This class is unique to the div that holds all information about a story
	c.OnHTML(".top-matter", func(e *colly.HTMLElement) {
		temp := item{}
		temp.StoryURL = e.ChildAttr("a[data-event-action=title]", "href")
		temp.Source = "https://old.reddit.com/r/golang/"
		temp.Title = e.ChildText("a[data-event-action=title]")
		temp.Comments = e.ChildAttr("a[data-event-action=comments]", "href")
		temp.CrawledAt = time.Now()
		stories = append(stories, temp)
	})

	// On every span tag with the class next-button
	c.OnHTML("span.next-button", func(h *colly.HTMLElement) {
		t := h.ChildAttr("a", "href")
		c.Visit(t)
	})

	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())

	})

	// Crawl all reddits the user passes in
	reddits := os.Args[1:]
	for _, reddit := range reddits {
		c.Visit(reddit)

	}

	c.Wait()
	//fmt.Println(stories)

	sortedCounts := wordCount(stories)
	for _, v := range sortedCounts {
		fmt.Printf(" %s: %d\n", v.word, v.count)
	}

}

type wordInfo struct {
	word  string
	count int
}

func wordCount(stories []item) []wordInfo {
	wordMap := map[string]int{}

	for _, story := range stories {
		words := strings.Split(story.Title, " ")
		for _, w := range words {
			wordMap[strings.ToLower(w)]++
		}
	}

	// wis = word infos
	wis := []wordInfo{}
	for w, c := range wordMap {
		wis = append(wis, wordInfo{
			word:  w,
			count: c,
		})

	}
	//sort in descending order
	sort.Slice(wis, func(i, j int) bool {
		return wis[i].count > wis[j].count
	})

	return wis

}

//todd's solution...

// c.Wait()
// 	// fmt.Println(stories)

// 	m := map[string]int{}

// 	for _, story := range stories {
// 		words := strings.Split(strings.ToLower(story.Title), " ")
// 		for _, word := range words {
// 			m[word]++
// 		}
// 	}

// 	type wordCount struct {
// 		word  string
// 		count int
// 	}

// 	xwc := []wordCount{}

// 	for w, c := range m {
// 		xwc = append(xwc, wordCount{
// 			word:  w,
// 			count: c,
// 		})
// 	}

// 	sort.Slice(xwc, func(i, j int) bool {
// 		return xwc[i].count > xwc[j].count
// 	})

// 	for _, v := range xwc {
// 		fmt.Println(v.word, v.count)
// 	}
