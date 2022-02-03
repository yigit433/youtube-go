package youtubego

func SearchVideos(searchq string) []Video {
	res := CreateRequest(searchq)

	return res
}
