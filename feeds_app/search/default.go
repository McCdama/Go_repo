package search

type defaultMatcher struct{}

// init registers the def. mat. with the prog.
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the def. mat.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
