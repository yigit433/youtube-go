package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Options struct {
	Limit      int
	SafeSearch bool
}

func SearchVideo(searchWord string, options Options) {
	Url, err := url.Parse("http://youtube.com/results")

	if err != nil {
		panic("The URL is incorrect!")
	}

	query := url.Values{}
	query.Add("search_query", searchWord)
	query.Add("sp", "EgIQAQ%253D%253D")

	Url.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", Url.String(), nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:78.0) Gecko/20100101 Firefox/78.0")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic("Something went wrong, the request cannot be sent to the URL!")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	bodyResp, err := io.ReadAll(res.Body)
	html := string(bodyResp)

	index := len(strings.Split(html, `{"itemSectionRenderer":`)) - 1
	items := strings.Split(html, `{"itemSectionRenderer":`)[index]
	var parsed = []byte(strings.Split(items, `},{"continuationItemRenderer":{`)[0])

	var out map[string]interface{}
	err = json.Unmarshal(parsed, &out)

	if err != nil {
		panic("Something went wrong, the problem was encountered while analyzing JSON!")
	}

	fmt.Println(out["contents"].([]interface{})[0].(map[string]interface{})["videoRenderer"])
}

func main() {
	SearchVideo("Duman eyvallah", Options{
		Limit: 10,
	})
}
