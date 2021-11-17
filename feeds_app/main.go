package main

import (
	_ "gotour/feeds_app/matchers" // Blank identifier, Compiler accept it & call any init functions within that package.
	"gotour/feeds_app/search"
	"log"
	"os"
)

// prior to main
func init() {
	log.SetOutput(os.Stdout)
}

// entry point
func main() {
	search.Run("fan")
}
