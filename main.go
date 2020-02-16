package main

import (
	"fmt"
	"sync"
)

//Fetcher defines an interface for any object that fetches a web url
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

//Crawl all urls upto a depth using a fetcher
func Crawl(url string, depth int, fetcher Fetcher) {
	crawler := NewCrawler(&CachingFetcher{
		fetcher:     fetcher,
		mutex:       &sync.Mutex{},
		visitedUrls: make(map[string]*result),
	})
	fmt.Printf("crawler: %#v\n", crawler)
	crawler.urlSource <- crawlable{url, depth}
	crawler.run()
}

