package youtubego

import (
	"errors"
	"io"
	"log"

	"github.com/yigit433/youtube-go/models"
	"github.com/yigit433/youtube-go/internal/httpClient"
	"github.com/yigit433/youtube-go/internal/parser"
	"github.com/yigit433/youtube-go/internal/utils"
)

func Search(query string, config models.SearchConfig) ([]models.SearchResult, error) {
	searchType := utils.GetSearchType(config.SearchType)

	params := map[string]string{
		"search_query": query,
		"sp":           searchType,
	}

	url, err := utils.BuildURLWithQueryParams("http://youtube.com/results", params)
	if err != nil {
		log.Fatalf("Cannot build the URL: %v", err)

		return nil, errors.New("cannot build the URL")
	}

	client := httpclient.New(config.Timeout)

	response, err := client.Get(url)
	if err != nil {
		log.Fatalf("Cannot send the request to the URL: %v", err)

		return nil, errors.New("cannot send the request to the URL")
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)

		return nil, errors.New("status code error: " + response.Status)
	}
	
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Cannot read the body stream.")

		return nil, errors.New("cannot read the body stream")
	}

	results := parser.ParseHTML(string(body), config.MaxResults)

	return results, nil
}