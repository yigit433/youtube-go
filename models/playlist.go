package models

type Playlist struct {
	ID        string    
	Title     string    
	Thumbnail Thumbnail 
	Videos    int       
	Channel   Channel   
}

type PlaylistParser struct {
	Playlist  Playlist
	IsSuccess bool    
}