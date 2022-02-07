package youtubego

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func CreateRequest(searchWord string) []Video {
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
	html := []byte(ParseHTML(string(bodyResp)))

	var out map[string]interface{}
	err = json.Unmarshal(html, &out)

	if err != nil {
		panic("Something went wrong, the problem was encountered while analyzing JSON!")
	}
	arr := out["contents"].([]interface{})
	output := []Video{}

	for i := 0; len(arr) > i; i++ {
		sdata := arr[i].(map[string]interface{})["videoRenderer"]
		parsedVideo := ParseVideo(sdata)

		if parsedVideo.IsSuccess {
			output = append(output, parsedVideo.Video)
		}
	}

	return output
}

func ParseHTML(html string) string {
	index := len(strings.Split(html, `{"itemSectionRenderer":`)) - 1
	items := strings.Split(html, `{"itemSectionRenderer":`)[index]

	return strings.Split(items, `},{"continuationItemRenderer":{`)[0]
}

func ParseVideo(data interface{}) VideoParser {
	if data != nil {
		thumbnail := data.(map[string]interface{})["thumbnail"].(map[string]interface{})["thumbnails"].([]interface{})

		var out VideoParser
		out = VideoParser{
			Video: Video{
				Id:    data.(map[string]interface{})["videoId"].(string),
				Title: data.(map[string]interface{})["title"].(map[string]interface{})["runs"].([]interface{})[0].(map[string]interface{})["text"].(string),
				Url:   fmt.Sprintf("https://www.youtube.com/watch?v=%s", data.(map[string]interface{})["videoId"].(string)),
				Thumbnail: Thumbnail{
					Id:     data.(map[string]interface{})["videoId"].(string),
					Url:    thumbnail[len(thumbnail)-1].(map[string]interface{})["url"].(string),
					Width:  fmt.Sprintf("%v", thumbnail[len(thumbnail)-1].(map[string]interface{})["width"]),
					Height: fmt.Sprintf("%v", thumbnail[len(thumbnail)-1].(map[string]interface{})["height"]),
				},
			},
			IsSuccess: true,
		}

		return out
	} else {
		return VideoParser{
			IsSuccess: false,
		}
	}
}
