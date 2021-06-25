package main

type searcher func(string) []string

func main() {
	f := func(s string) []string {
		return []string{s}
	}
	searchers := []searcher{f}
	searchData("", searchers)
}
func searchData(data string, processors []searcher) []string {
	done := make(chan struct{})
	result := make(chan []string)
	for _, processor := range processors {
		go func(s searcher) {
			select {
			case result <- s(data):
			case <-done:
			}
		}(processor)
	}
	r := <-result
	close(done)
	return r
}
