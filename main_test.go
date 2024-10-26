package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestRemoveStopWords(t *testing.T) {
	exp := []string{
		"Applications", "robotics",
	}
	words := []string{
		"Applications", "of", "robotics",
	}

	res := removeStopwords(words)

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %s, got %s instead.", exp, res)
	}

}

func TestScrape(t *testing.T) {
	var wg sync.WaitGroup
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics"}
	sites := make([]Site, 0, 1)

	for _, URL := range urls {
		wg.Add(1)
		go Scrape(URL, &wg, &sites)
	}

}
