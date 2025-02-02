// https://go.dev/tour/concurrency/10

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Moving FakeResult definition here for readability
type fakeResult struct {
	body string
	urls []string
}

// Cache data structure with mutex
type Cache struct {
	// Internal map which stores cached URLs along with their content (part atleast)
	v map[string]*fakeResult
	// Mutex
	mu sync.Mutex
}

// Function to get the value from the URL in the cache
func (cache *Cache) Get(url string) (*fakeResult, bool) {
	cache.mu.Lock()
	// Defer coz we want to unlock after the function returns the value
	defer cache.mu.Unlock() 
	result, ok := cache.v[url]
	return result, ok
}

// Function to update a value in the cache
func (cache *Cache) Update(url string, body string, urls []string, err error) {
	cache.mu.Lock()
	
	// Initialize the internal result struct
	cache.v[url] = &fakeResult{}
	
	// If an error is there, just put the error message in body and leave urls empty
	if err != nil {
		cache.v[url].body = fmt.Sprint(err) 

	} else {
		// Otherwise copy the body and URLs as it is
		cache.v[url].body = body
		cache.v[url].urls = urls
	}
	cache.mu.Unlock()
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// Modified to take in the Cache data structure as an input
func Crawl(url string, depth int, fetcher Fetcher, cache *Cache) {
	if depth <= 0 {
		return
	}
	
	// Don't fetch the same URL twice.
	// First check the cache for existing values
	// If available there, then no need to fetch again
	// We are just interested in existance of value, so no need to store the actual result
	if _, ok := cache.Get(url); ok {
		return
	}
	
	body, urls, err := fetcher.Fetch(url)
	
	// Update the cache
	cache.Update(url, body, urls, err)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	// Fetch URLs in parallel
	// We use sync.WaitGroup for the same
	wg := &sync.WaitGroup{}
	for _, u := range urls {
    // Add the task to the WaitGroup
		wg.Add(1)
    // Need to use the actual WaitGroup and not its copies
		go func(wg *sync.WaitGroup) {
			Crawl(u, depth-1, fetcher, cache)
      // Indicate the task is finished when the above function returns 
			wg.Done()
		}(wg)
	}
	wg.Wait()
	return
}

func main() {
	// Initialize the cache
	cache := Cache {v: make(map[string]*fakeResult)}
	Crawl("https://golang.org/", 4, fetcher, &cache)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
