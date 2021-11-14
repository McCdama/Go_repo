package concurrency

import "time"

var URLs = []string{
	"https://gitlab.com/",
	"https://www.bnf.fr/fr",
	"https://www.gov.uk/",
	"https://www.francemusique.fr/",
}

type result struct {
	url     string
	elapsed time.Duration
	status  string
	size    int
	err     error
}
