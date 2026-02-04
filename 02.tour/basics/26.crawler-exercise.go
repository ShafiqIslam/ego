package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// ---- NEW: concurrency-safe visited cache ----

type SafeVisited struct {
	mu      sync.Mutex
	visited map[string]bool
}

func (s *SafeVisited) CheckAndMark(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.visited[url] {
		return true
	}
	s.visited[url] = true
	return false
}

// ---- MODIFIED Crawl ----

func Crawl(url string, depth int, fetcher Fetcher, visited *SafeVisited, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	// don't fetch same URL twice
	if visited.CheckAndMark(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, visited, wg)
	}
}

func main() {
	visited := &SafeVisited{
		visited: make(map[string]bool),
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go Crawl("https://golang.org/", 4, fetcher, visited, &wg)

	wg.Wait()
}

// ---- fake fetcher (unchanged) ----

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": {
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": {
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": {
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": {
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
