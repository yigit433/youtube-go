package main

import (
	"fmt"
	"net/url"
)

type Options struct {
	limit      int
	safeSearch bool
}

func Search(searchWord string, options Options) {
	Url, err := url.Parse("https://youtube.com/results")

	if err != nil {
		panic("The URL is incorrect!")
	}

	query := url.Values{}
	query.Add("q", searchWord)
	query.Add("hl", "en")

	Url.RawQuery = query.Encode()

	fmt.Println(Url.String())
}

func main() {
	Search("Hello world", Options{
		limit: 5,
	})
}
