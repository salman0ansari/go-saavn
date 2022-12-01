package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var (
	// BaseURL is the base URL for the Saavn API
	BaseURL        string = "https://jiosaavn-api-privatecvc.vercel.app"
	SearchEndpoint string = "/search/songs?query="
	Bitrate        string = "320"
	limit          string = "10"
)

func SearchSong(query string) SearchResponse {
	// return func() tea.Msg {
	var url string = (BaseURL + SearchEndpoint + query + "&limit=" + limit)
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)
	if err != nil {
		// return errMsg{err}
	}
	defer res.Body.Close()
	var out SearchResponse
	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		// return errMsg{err}
	}
	return out

	// return resultMsg(out)
	// }
}

type resultMsg SearchResponse

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

type SearchResponse struct {
	Status  string `json:"status"`
	Results []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Album struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"album"`
		Year             string      `json:"year"`
		ReleaseDate      interface{} `json:"releaseDate"`
		Duration         string      `json:"duration"`
		Label            string      `json:"label"`
		PrimaryArtists   string      `json:"primaryArtists"`
		PrimaryArtistsID string      `json:"primaryArtistsId"`
		ExplicitContent  int         `json:"explicitContent"`
		PlayCount        string      `json:"playCount"`
		Language         string      `json:"language"`
		HasLyrics        string      `json:"hasLyrics"`
		Artist           string      `json:"artist"`
		Image            []struct {
			Quality string `json:"quality"`
			Link    string `json:"link"`
		} `json:"image"`
		URL         string `json:"url"`
		Copyright   string `json:"copyright"`
		DownloadURL []struct {
			Quality string `json:"quality"`
			Link    string `json:"link"`
		} `json:"downloadUrl"`
	} `json:"results"`
}
