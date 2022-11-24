package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func WriteToFile(songId string, songData *http.Response, ext string) {

	if _, err := os.Stat("temp"); os.IsNotExist(err) {
		os.Mkdir("temp", 0755)
	}

	songFile, err := os.Create("temp/" + songId + "." + ext)
	if err != nil {
		fmt.Println("Can not create file")
	}
	defer songFile.Close()
	io.Copy(songFile, songData.Body)
}

func GetResp(url string) (result *http.Response, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
