package youtubego

type Thumbnail struct {
	Id     string
	Width  string
	Height string
	Url    string
}

type Video struct {
	Thumbnail
	Id    string
	Title string
	Url   string
}

type VideoParser struct {
	Video
	IsSuccess bool
}
