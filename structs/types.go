package structs

type ResultStruct struct {
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
