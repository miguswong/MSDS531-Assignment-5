package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

// Initialize struct to hold wiki data
type Site struct {
	Url      string   `json:"name"`
	Title    string   `json:"title"`
	BodyText string   `json:"bodytext"`
	Tags     []string `json:"tags"`
}

// Function to remove stopwords, returns list of strings with stopwords removed
func removeStopwords(words []string) []string {
	stopwords := map[string]bool{
		"and": true, "the": true, "is": true, "in": true, "of": true, "a": true, "an": true, // Add more stopwords as needed
	}
	var result []string
	for _, word := range words {
		if !stopwords[word] {
			result = append(result, word)
		}
	}
	return result
}

// takes a url, and creates a struct as the returned value.
func Scrape(URL string, wg *sync.WaitGroup, sites *[]Site) {
	defer wg.Done()
	var targetSite Site
	//Create Empty struct targetSite
	targetSite.Url = URL

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
	// Grab Body Text
	c.OnHTML(".mw-body-content", func(e *colly.HTMLElement) {
		targetSite.BodyText = e.Text
	})

	//Parse URL
	parsedURL, err := url.Parse(URL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}
	// Split the URL path into segments
	pathSegments := strings.Split(parsedURL.Path, "/")

	// Initial tags list from URL segments
	tagsList := []string{pathSegments[1], pathSegments[2]}

	var moreTags []string
	// Generate more tags

	urlPart := pathSegments[2]
	words := strings.Split(urlPart, "_")
	moreTags = removeStopwords(words)

	// Clean and add more tags
	re := regexp.MustCompile(`[^a-zA-Z]`)
	for _, tag := range moreTags {
		cleanedTag := re.ReplaceAllString(tag, "")
		tagsList = append(tagsList, strings.ToLower(cleanedTag))
	}
	targetSite.Tags = tagsList

	c.Visit(URL)

	*sites = append(*sites, targetSite)
}

func main() {
	var wg sync.WaitGroup
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

	//Initialize Site slice called sites
	sites := make([]Site, 0, 500)

	//Iterate through URL list and collect data
	for _, URL := range urls {
		wg.Add(1)
		go Scrape(URL, &wg, &sites)
	}

	//
	wg.Wait()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	enc.Encode(sites)
}
