package search

import (
	"fmt"
	"log"
)

// Result contains the result of a search
type Result struct {
	Field   string
	Content string
}

// Matcher defiens the behavior required by types that want to implement a new search type
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for each individual feed to run search concurrently
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// Perform the search against the specified matcher
	searchResult, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channels
	for _, result := range searchResult {
		results <- result
	}
}

// Display writes results to the terminal as they are received by the individual goroutine
func Display(results chan *Result) {
	// The channel blocks until a result is written to the channel
	// Once the channel is closed the for loop terminates
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
