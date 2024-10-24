package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gocolly/colly"
)

type Site struct {
	Url      string `json:"name"`
	Title    string `json:"title"`
	BodyText string `json:"bodytext"`
}

// scraping functions
func scrape(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	//Initialize Scraper
	c := colly.NewCollector(
		// Visit only domains: en.wikipedia.org
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	//Find the tile of the Wiki Page
	c.OnHTML(".mw-page-title-main", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	// mw-content-text
	c.OnHTML(".mw-body-content", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	// c.OnHTML(".searchaux", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	c.Visit(url)
}

func main() {

	//Initialize item.json file
	fName := "item.jl"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	// Wikipedia URLs for topic of interest
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}
	//Initialize slice of courses
	sites := make([]Site, 0, 100)

	for _, url := range urls {

		//Create Empty struct targetSite
		var targetSite Site
		targetSite.Url = url

		//Initialize Scraper
		c := colly.NewCollector(
			// Visit only domains: en.wikipedia.org
			colly.AllowedDomains("en.wikipedia.org"),
		)

		// Before making a request print "Visiting ..."
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("visiting", r.URL.String())
		})

		//Find the tile of the Wiki Page
		c.OnHTML(".mw-page-title-main", func(e *colly.HTMLElement) {
			targetSite.Title = e.Text
		})
		// mw-content-text
		c.OnHTML(".mw-body-content", func(e *colly.HTMLElement) {
			targetSite.BodyText = e.Text
		})

		c.Visit(url)
		sites = append(sites, targetSite)
	}

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	enc.Encode(sites)
}
