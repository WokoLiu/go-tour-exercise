// https://tour.golang.org/concurrency/10
// https://tour.go-zh.org/concurrency/10

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

// 最终返回的数据
type Fetched struct {
	data map[string][]fakeResult //最终爬取到的数据
	urls map[string]int          // 这里存每个url访问的次数
	mux  sync.Mutex
	ch   chan string // 按顺序存储每次访问的url
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, fetched *Fetched) {
	// Fetch URLs in parallel.
	// Don't fetch the same URL twice.
	if depth <= 0 {
		return
	}
	fetched.mux.Lock()
	_, ok := fetched.urls[url]
	if ok {
		fetched.mux.Unlock()
		return
	}
	fetched.ch <- url
	body, urls, err := fetcher.Fetch(url)
	fetched.urls[url]++
	fetched.data[url] = append(fetched.data[url], fakeResult{body, urls})
	if err != nil {
		fmt.Println(err)
		fetched.mux.Unlock()
		return
	}
	fetched.mux.Unlock()
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, fetched)
	}
	return
}

func main() {
	ch := make(chan string)
	fetched := Fetched{data: make(map[string][]fakeResult), urls: make(map[string]int), ch: ch}
	go Crawl("https://golang.org/", 4, fetcher, &fetched)
	for i := 0; i < 5; i++ { // 因为一共有5个 url
		fmt.Println("ch", <-ch)
	}
	// 需要等待一段时间再print，否则看不到数据（print ch 的时间足够了）
	fmt.Println(fetched)
}

// fakeFetcher is Fetcher that returns canned results.
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
