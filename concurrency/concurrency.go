package concurrency

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var URLs = []string{
	"https://gitlab.com/",
	"https://www.bnf.fr/fr",
	"https://www.gov.uk/",
	"https://www.francemusique.fr/",
}

type Result struct {
	Url     string
	Elapsed time.Duration
	Status  string
	Size    int
	Err     error
}

func Run() {
	ch := make(chan Result)

	// fire one goroutine per URL
	for _, Url := range URLs {
		go Get(ch, Url)
		// result for each URL goes here ...
	}
}

func Get(ch chan Result, url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{Err: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		ch <- Result{
			Url:     url,
			Elapsed: time.Since(start),
			Err:     fmt.Errorf("HTTP status code %d", resp.StatusCode),
			Status:  resp.Status,
		}
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- Result{
			Url:     url,
			Elapsed: time.Since(start),
			Err:     err,
			Status:  resp.Status,
		}
		return
	}

	// send the result back to the main goroutine via the channel
	ch <- Result{
		Url:     url,
		Elapsed: time.Since(start),
		Status:  resp.Status,
		Size:    len(body),
	}

}
