package utils

import "strings"

const (
	SearchTypeVideoId    string = "EgIQAQ%253D%253D"
	SearchTypePlaylistId string = "EgIQAw%253D%253D"
	SearchTypeChannelId  string = "EgIQAg%253D%253D"
)

func GetSearchType(searchType string) string {
	switch strings.ToLower(searchType) {
	case "video":
		return SearchTypeVideoId
	case "playlist":
		return SearchTypePlaylistId
	case "channel":
		return SearchTypeChannelId
	default:
		return ""
	}
}