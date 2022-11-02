package search

import (
	"bufio"
	"fmt"
	"os"
)

const dataFile = "data/data.txt"

type Feed struct {
	URI string
}

// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once the function returns.
	defer file.Close()

	// Decode the file into a slice of pointers to Feed values.
	var feeds []*Feed

	//err = json.NewDecoder(file).Decode(&feeds)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		uri := scanner.Text()
		if uri != "" {
			feeds = append(feeds, &Feed{URI: scanner.Text()})
		}

	}
	// for _, n := range feeds {
	// 	fmt.Printf("%s, %+v\n", n.URI, n)
	// }

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}
