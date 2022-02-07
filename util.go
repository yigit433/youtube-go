package youtubego

import (
	"encoding/json"
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
	arr := out["contents"]

	output := make([]Video, len(arr.([]interface{})))

	for i := 0; len(arr.([]interface{})) > i; i++ {
		pureData := arr.([]interface{})[i].(map[string]interface{})["videoRenderer"]

		output[i] = ParseVideo(pureData)
	}

	return output
}

func ParseHTML(html string) string {
	index := len(strings.Split(html, `{"itemSectionRenderer":`)) - 1
	items := strings.Split(html, `{"itemSectionRenderer":`)[index]

	return strings.Split(items, `},{"continuationItemRenderer":{`)[0]
}

func ParseVideo(data map[string]interface{}) Video {
	var out Video

	out = Video{
		Id: data["videoId"].(string),
	}

	return out
}
