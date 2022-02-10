package youtubego

type Thumbnail struct {
	Id, Width, Height, Url string
}

type Video struct {
	Thumbnail
	Id, Title, Url string
}

type VideoParser struct {
	Video
	IsSuccess bool
}

type Channel struct {
	Thumbnail
	Id, Url, Name, Subscribers string
	Verified                   bool
}

type ChannelParser struct {
	Channel
	IsSuccess bool
}

type Playlist struct {
	Thumbnail
	Channel
	Id, title string
	Videos    int
}

type PlaylistParser struct {
	Playlist
	IsSuccess bool
}

type SearchResult struct {
	Video
	Playlist
	Channels
}

type SearchOptions struct {
	Limit int
	Type  string
}
