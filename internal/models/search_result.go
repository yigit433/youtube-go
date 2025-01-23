package models

import "time"

type SearchResult struct {
	Video    *Video   
	Playlist *Playlist 
	Channel  *Channel 
}

type SearchConfig struct {
	SearchType string
	Timeout    time.Duration
	MaxResults int
}