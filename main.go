package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/helper"
	"main/structs"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/zakaria-chahboun/cute"
)

var API_URL string = "https://jiosaavn-api-privatecvc.vercel.app"
var Search_EP string = "/search/songs?query="
var BITRATE string = "320"
var limit string = "5"

// var ALLOWED_BITRATE := [5]int8{12,48,96,160,320}

func fetchSearch(searchTerm string) (results []string, err error) {
	var URL string = (API_URL + Search_EP + searchTerm + "&limit=" + limit)
	// fmt.Println("GET " + URL)

	resp, err := helper.GetResp(URL)
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		var FetchedItems structs.ResultStruct
		if err := json.Unmarshal(body, &FetchedItems); err != nil {
			cute.Check("JSON Error", errors.New("can not unmarshal JSON"))
		}
		var i int = 0

		songList := cute.NewList(cute.BrightBlue, "Search Result")
		for _, item := range FetchedItems.Results {
			i = i + 1
			index := strconv.Itoa(i)
			songAlbum := string(item.Album.Name)
			songName := string(item.Name)
			songArtist := string(item.Artist)
			songList.Add(cute.BrightBlue, index+" | "+songName+" | "+songArtist+" | "+songAlbum)
		}
		songList.Print()

		cute.Printf("Enter Track Index", "")

		var songIndexes string
		_, err = fmt.Scanln(&songIndexes)
		cute.Check("Error Input", err)

		if !strings.ContainsAny(songIndexes, "0123456789") {
			cute.Check("Error Input", errors.New("invalid input"))
		}

		if songIndexes == "" || songIndexes == "0" || songIndexes > limit {
			cute.Check("Error Input", errors.New("invalid index"))
			return nil, nil
		}
		// split songIndexes by comma and store in array
		var songIndexesArray []string = strings.Split(songIndexes, ",")
		list := cute.NewList(cute.BrightBlue, "Downloading Songs")

		for _, songIndex := range songIndexesArray {
			i, _ := strconv.ParseInt(songIndex, 0, 8)
			item := FetchedItems.Results[i-1]
			songAlbum := item.Album.Name
			songName := item.Name
			list.Add(cute.BrightGreen, songName+" | "+songAlbum+" | "+item.Artist)
		}
		list.Print()

		for _, songIndex := range songIndexesArray {
			i, _ := strconv.ParseInt(songIndex, 0, 8)
			item := FetchedItems.Results[i-1]
			songAlbum := item.Album.Name
			songName := item.Name
			songYear := item.Year
			songArtist := item.Artist
			songImgUrl := item.Image[len(item.Image)-1].Link
			songCopyright := item.Copyright
			songPublisher := "go-saavn"
			songDownloadUrl := item.DownloadURL[len(item.DownloadURL)-1].Link

			// cute.Println("Downloading", ""+songName+" - "+songAlbum+" by "+songArtist+" ("+songYear+")")

			songData, _ := helper.GetResp(songDownloadUrl)
			defer songData.Body.Close()
			songid := []byte(item.ID)
			songId := b64.StdEncoding.EncodeToString([]byte(songid))

			songComment := "Downloaded using go-saavn by @salman0ansari \n https://github.com/salman0ansari/go-saavn"

			// write song to file
			helper.WriteToFile(songId, songData, "mp3")
			cute.Println("Downloaded", songName+" - "+songAlbum+" by "+songArtist+" ("+songYear+")")

			songData, _ = helper.GetResp(songImgUrl)
			defer songData.Body.Close()
			helper.WriteToFile(songId, songData, "jpg")

			output, _ := os.Getwd()
			currentWD := output + "/" + songName + "-" + songYear + ".mp3"

			arguments := []string{"-i", "./temp/" + songId + ".mp3", "-i", "./temp/" + songId + ".jpg",
				"-map", "0:0", "-map", "1:0", "-c", "copy", "-id3v2_version", "3",
				"-metadata", "title=" + songName,
				"-metadata", "album=" + songAlbum,
				"-metadata", "artist=" + songArtist,
				"-metadata", "date=" + songYear,
				"-metadata", "album_artist=" + songArtist,
				"-metadata", "copyright=" + songCopyright,
				"-metadata", "publisher=" + songPublisher,
				"-metadata", "comment=" + songComment,
				"-codec:a", "libmp3lame", "-b:a", "320k", "-hide_banner", "-y", currentWD}

			_, err := exec.Command("ffmpeg", arguments...).CombinedOutput()

			if err != nil {
				cute.Check("ffmpeg error", err)
			}

			// delete directory
			_ = os.RemoveAll("temp/")

		}
		return nil, err
	}
	return
}

func main() {
	cute.Printf("Enter Song Name", "")
	var searchQuery string
	fmt.Scanln(&searchQuery)
	fetchSearch(searchQuery)
}
