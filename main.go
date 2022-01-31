package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
)

type Options struct {
	limit      int
	safeSearch bool
}

func SearchVideo(searchWord string, options Options) {
	Url, err := url.Parse("https://youtube.com/results")

	if err != nil {
		panic("The URL is incorrect!")
	}

	query := url.Values{}
	query.Add("search_query", searchWord)
	query.Add("hl", "en")
	query.Add("sp", "EgIQAQ%253D%253D")

	Url.RawQuery = query.Encode()

	res, err := http.Get(Url.String())

	if err != nil {
		panic("Something went wrong, the request cannot be sent to the URL!")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div#contents").Each(func(index int, s *goquery.Selection) {
		// The part of listing videos will be done
	})
}

func main() {
	SearchVideo("Hello world", Options{
		limit: 10,
	})
}
