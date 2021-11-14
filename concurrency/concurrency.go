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

		// get one result back for each of the URLs
		var results []Result
		for range URLs {
			// receive a result from whichever of the pawned goroutines
			// that's finished; blocks until any one finishes.
			res := <-ch
			results = append(results, res)
		}

		// print phase
		for _, res := range results {
			if res.Err != nil {
				fmt.Println(res.Url, res.Elapsed, res.Status, res.Err)
			} else {
				fmt.Println(res.Url, res.Elapsed, res.Status, res.Size)
			}
		}
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
