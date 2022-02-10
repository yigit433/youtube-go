package youtubego

func Search(searchq string, options SearchOptions) []interface{} {
	res := CreateRequest(searchq, options)

	return res
}
