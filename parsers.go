package youtubego

import "fmt"

func ParsePlaylist(data interface{}) PlaylistParser {
	if data != nil {
		return PlaylistParser{
			IsSuccess: false,
		}
	} else {
		return PlaylistParser{
			IsSuccess: false,
		}
	}
}

func ParseChannel(data interface{}) ChannelParser {
	if data != nil {
		navEndpoint := data.(map[string]interface{})["navigationEndpoint"]
		url := fmt.Sprintf("https://www.youtube.com%s", navEndpoint.(map[string]interface{})["browseEndpoint"].(map[string]interface{})["canonicalBaseUrl"] || navEndpoint.(map[string]interface{})["commandMetadata"].(map[string]interface{})["webCommandMetadata"].(map[string]interface{})["url"])
		thumbnail := data.(map[string]interface{})["thumbnail"].(map[string]interface{})["thumbnails"]

		var out ChannelParser
		out = ChannelParser{
			Id:          data.(map[string]interface{})["channelId"].(string),
			Url:         url,
			Name:        data.(map[string]interface{})["title"].(map[string]interface{})["simpleText"].(string),
			Icon:        thumbnail.([]interface{})[len(thumbnail.([]interface{}))-1],
			Subscribers: data.(map[string]interface{})["subscriberCountText"].(map[string]interface{})["simpleText"],
		}
	} else {
		return ChannelParser{
			IsSuccess: false,
		}
	}
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
