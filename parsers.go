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
		var url string
		if navEndpoint.(map[string]interface{})["browseEndpoint"].(map[string]interface{})["canonicalBaseUrl"] {
			url = fmt.Sprintf("https://www.youtube.com%s", navEndpoint.(map[string]interface{})["browseEndpoint"].(map[string]interface{})["canonicalBaseUrl"])
		} else if navEndpoint.(map[string]interface{})["commandMetadata"].(map[string]interface{})["webCommandMetadata"].(map[string]interface{})["url"] {
			url = fmt.Sprintf("https://www.youtube.com%s", navEndpoint.(map[string]interface{})["commandMetadata"].(map[string]interface{})["webCommandMetadata"].(map[string]interface{})["url"])
		}

		thumbnails := data.(map[string]interface{})["thumbnail"].(map[string]interface{})["thumbnails"]
		thumbnail := thumbnails.([]interface{})[len(thumbnails.([]interface{}))-1]

		var out ChannelParser
		out = ChannelParser{
			Id:   data.(map[string]interface{})["channelId"].(string),
			Url:  url,
			Name: data.(map[string]interface{})["title"].(map[string]interface{})["simpleText"].(string),
			Icon: Thumbnail{
				Url:    fmt.Sprintf("%v", thumbnail["url"]),
				Width:  fmt.Sprintf("%v", thumbnail["width"]),
				Height: fmt.Sprintf("%v", thumbnail["height"]),
			},
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
