## Youtube GO

A Go package prepared for Video searching on youtube.

> This project was developed inspired by [youtube-sr](https://github.com/DevAndromeda/youtube-sr)

### Supported
- ✅ Regular YouTube Search (Video/Channel/Playlist)
- ✅ Get specific video
- 🛠 Get Playlist (including all videos)
- 🛠 YouTube safe search

### Installation
```bash
go get github.com/yigit433/youtube-go
```

### Example
```go
package main

import (
  "fmt"
  "time"

  "github.com/yigit433/youtube-go"
  "github.com/yigit433/youtube-go/models"
)

func main() {
  result, err := youtubego.Search("ABBA Money, Money, Money", models.SearchConfig{
    SearchType: "video", // channel , playlist , all
    Timeout: 10 * time.Second, 
    MaxResults: 15,
  })
  if err != nil {
	panic(err)
  }

  fmt.Println(result[0].Video.Title)
}
```
