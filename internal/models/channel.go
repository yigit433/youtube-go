package models

type Channel struct {
	ID          string   
	Name        string    
	Url         string    
	Subscribers string       
	Verified    bool      
	Icon        Thumbnail 
}

type ChannelParser struct {
	Channel   Channel 
	IsSuccess bool    
}