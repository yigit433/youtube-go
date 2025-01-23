package parser

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/yigit433/youtube-go/models"
	"github.com/yigit433/youtube-go/internal/utils"
)

func ParseHTML(html string, maxResults int) []models.SearchResult {
	parsedJSON, err := extractJSONFromHTML(html)
	if err != nil {
		log.Fatalf("Cannot extract JSON from HTML: %v", err)
	}

	contents, err := getContents(parsedJSON)
	if err != nil {
		log.Fatalf("Cannot get contents from parsed JSON: %v", err)
	}

	results := parseContents(contents, maxResults)
	// for _, result := range results {
	// 	if result.Video != nil {
	// 		fmt.Printf("Video: %v\n", result.Video)
	// 	} else if result.Playlist != nil {
	// 		fmt.Printf("Playlist: %v\n", result.Playlist)
	// 	} else if result.Channel != nil {
	// 		fmt.Printf("Channel: %v\n", result.Channel)
	// 	}
	// }

	return results
}

func extractJSONFromHTML(html string) (map[string]interface{}, error) {
	index := len(strings.Split(html, `{"itemSectionRenderer":`)) - 1
	items := strings.Split(html, `{"itemSectionRenderer":`)[index]
	parsedJSON := strings.Split(items, `},{"continuationItemRenderer":{`)[0]

	var out map[string]interface{}
	if err := json.Unmarshal([]byte(parsedJSON), &out); err != nil {
		return nil, errors.New("unable to unmarshal HTML into JSON")
	}

	return out, nil
}

func getContents(parsedJSON map[string]interface{}) ([]interface{}, error) {
	contents, exists := parsedJSON["contents"]
	if !exists {
		return nil, errors.New("'contents' not found in parsed JSON")
	}

	arr, ok := contents.([]interface{})
	if !ok {
		return nil, errors.New("'contents' is not an array")
	}

	return arr, nil
}

func parseContents(contents []interface{}, limit int) []models.SearchResult {
	var results []models.SearchResult

	for i := 0; i < limit; i++ {
		item := contents[i].(map[string]interface{})

		// Try to parse as Video, Playlist, or Channel
		if videoRenderer, ok := item["videoRenderer"]; ok {
			if video := utils.ParseVideo(videoRenderer); video.IsSuccess {
				results = append(results, models.SearchResult{Video: &video.Video})
			}
		} else if playlistRenderer, ok := item["playlistRenderer"]; ok {
			if playlist := utils.ParsePlaylist(playlistRenderer); playlist.IsSuccess {
				results = append(results, models.SearchResult{Playlist: &playlist.Playlist})
			}
		} else if channelRenderer, ok := item["channelRenderer"]; ok {
			if channel := utils.ParseChannel(channelRenderer); channel.IsSuccess {
				results = append(results, models.SearchResult{Channel: &channel.Channel})
			}
		}
	}

	return results
}