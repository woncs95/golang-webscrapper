package urlchecker

import (
	"fmt"
	"net/http"
)

type urlStatus struct {
	url    string
	status string
}

func CheckURL() {

	results := make(map[string]string)
	c := make(chan urlStatus)
	urls := []string{
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.naver.com/",
		"https://www.reddit.com/",
		"https://www.soundcloud.com/",
		"https://www.instagram.com/",
		"https://www.airbnb.com",
	}

	for _, url := range urls {
		go hitURL(url, c)

	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

//chan<- means send only!
func hitURL(url string, c chan<- urlStatus) {
	// fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- urlStatus{url: url, status: status}
}

//we want to do it asynchronous
