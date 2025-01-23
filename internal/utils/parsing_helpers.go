package utils

import (
	"fmt"
	"github.com/yigit433/youtube-go/models"
)

// ParsePlaylist parses the playlist data and returns a PlaylistParser
func ParsePlaylist(data interface{}) models.PlaylistParser {
	if data == nil {
		return models.PlaylistParser{IsSuccess: false}
	}

	// Extract thumbnail information
	thumbnails := getArrayValue(data, "thumbnail", "thumbnails")
	if len(thumbnails) == 0 {
		return models.PlaylistParser{IsSuccess: false}
	}

	thumbnail := thumbnails[len(thumbnails)-1]

	// Parse the rest of the data
	playlist := models.Playlist{
		ID:       getStringValue(data, "playlistId"),
		Title:    getStringValue(data, "title", "simpleText"),
		Thumbnail: models.Thumbnail{
			Url:    getStringValue(thumbnail, "url"),
			Width:  getFloat64Value(thumbnail, "width"),
			Height: getFloat64Value(thumbnail, "height"),
		},
	}

	return models.PlaylistParser{
		Playlist:  playlist,
		IsSuccess: true,
	}
}

// ParseChannel parses the channel data and returns a ChannelParser
func ParseChannel(data interface{}) models.ChannelParser {
	if data == nil {
		return models.ChannelParser{IsSuccess: false}
	}

	// Extract thumbnail information
	thumbnails := getArrayValue(data, "thumbnail", "thumbnails")
	if len(thumbnails) == 0 {
		return models.ChannelParser{IsSuccess: false}
	}

	thumbnail := thumbnails[len(thumbnails)-1]

	// Parse the rest of the data
	channel := models.Channel{
		ID:          getStringValue(data, "channelId"),
		Name:        getStringValue(data, "title", "simpleText"),
		Icon:        models.Thumbnail{
			Url:    getStringValue(thumbnail, "url"),
			Width:  getFloat64Value(thumbnail, "width"),
			Height: getFloat64Value(thumbnail, "height"),
		},
		Subscribers: getStringValue(data, "subscriberCountText", "simpleText"),
	}

	return models.ChannelParser{
		Channel:   channel,
		IsSuccess: true,
	}
}

// ParseVideo parses the video data and returns a VideoParser
func ParseVideo(data interface{}) models.VideoParser {
	if data == nil {
		return models.VideoParser{IsSuccess: false}
	}

	// Extract thumbnail information
	thumbnails := getArrayValue(data, "thumbnail", "thumbnails")
	if len(thumbnails) == 0 {
		return models.VideoParser{IsSuccess: false}
	}

	thumbnail := thumbnails[len(thumbnails)-1]

	// Parse the rest of the data
	video := models.Video{
		ID:    getStringValue(data, "videoId"),
		Title: getStringValue(data, "title", "runs", "0", "text"),
		Url:   fmt.Sprintf("https://www.youtube.com/watch?v=%s", getStringValue(data, "videoId")),
		Thumbnail: models.Thumbnail{
			ID:     getStringValue(data, "videoId"),
			Url:    getStringValue(thumbnail, "url"),
			Width:  getFloat64Value(thumbnail, "width"),
			Height: getFloat64Value(thumbnail, "height"),
		},
	}

	return models.VideoParser{
		Video:     video,
		IsSuccess: true,
	}
}
