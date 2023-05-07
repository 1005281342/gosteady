package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kkdai/youtube/v2"
)

//playlistID = "PLtTlf9YsVQqbiXojzwBJsxG75rOK3T64e"
func main() {
	router := gin.Default()
	router.GET("/", handleRequest)
	log.Fatal(router.Run(":8080"))
}

func handleRequest(c *gin.Context) {
	playlistID := c.Query("playlistID")
	if playlistID == "" {
		c.String(http.StatusBadRequest, "Missing playlistID parameter")
		return
	}

	proxyURL, err := url.Parse("http://127.0.0.1:20171")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	httpClient := &http.Client{Transport: &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}}

	ytClient := &youtube.Client{
		HTTPClient: httpClient,
	}

	playlist, err := ytClient.GetPlaylist(playlistID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	/* ----- Create playlist folder if it doesn't exist ----- */
	playlistFolder := playlist.Title
	if _, err := os.Stat(playlistFolder); os.IsNotExist(err) {
		err = os.Mkdir(playlistFolder, 0755)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	/* ----- Enumerating playlist videos ----- */
	header := fmt.Sprintf("Playlist %s by %s", playlist.Title, playlist.Author)
	fmt.Printf("%s\n%s\n\n", header, strings.Repeat("=", len(header)))

	go func() {
		for k, v := range playlist.Videos {
			fmt.Printf("(%d) %s - '%s'\n", k+1, v.Author, v.Title)
			for i := 0; i < 5; i++ {
				if err = downloadVideo(ytClient, v, playlistFolder); err != nil {
					log.Println(err)
					time.Sleep(time.Minute * time.Duration(i+1))
					continue
				}
				fmt.Printf("done (%d) %s - '%s'\n", k+1, v.Author, v.Title)
				break
			}
		}
	}()

	c.String(http.StatusOK, "后台下载中")
}

func downloadVideo(ytClient *youtube.Client, v *youtube.PlaylistEntry, folder string) error {
	video, err := ytClient.VideoFromPlaylistEntry(v)
	if err != nil {
		return err
	}

	var formats = video.Formats.WithAudioChannels()
	formats.Sort()
	// Now it's fully loaded.
	fmt.Printf("Downloading %s by '%s'!\n", video.Title, video.Author)

	stream, _, err := ytClient.GetStream(video, &formats[0])
	if err != nil {
		return err
	}

	filePath := folder + "/" + v.Title + ".mp4"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer func() {
		if file != nil {
			_ = file.Close()
		}
	}()
	_, err = io.Copy(file, stream)
	if err != nil {
		return err
	}

	return nil
}
