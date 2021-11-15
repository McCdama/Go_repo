package search

import "log"

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
