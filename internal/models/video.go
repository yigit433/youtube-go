package models

type Video struct {
	ID        string   
	Title     string    
	Url       string    
	Thumbnail Thumbnail
}

type VideoParser struct {
	Video     Video  
	IsSuccess bool   
}