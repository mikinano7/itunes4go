package itunes4go

import "time"

type Request struct {
	Term      string `url:"term"`
	Country   string `url:"country,omitempty"`
	Media     string `url:"media,omitempty"`
	Entity    string `url:"entity,omitempty"`
	Attribute string `url:"attribute,omitempty"`
	Callback  string `url:"callback,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Lang      string `url:"lang,omitempty"`
	Version   int    `url:"version,omitempty"`
	Explicit  string `url:"explicit,omitempty"`
}

type Response struct {
	ResultCount int      `json:"resultCount"`
	Results     []Result `json:"results"`
}

type Result struct {
	WrapperType            string    `json:"wrapperType"`
	Kind                   string    `json:"kind"`
	ArtistId               int       `json:"artistId"`
	CollectionId           int       `json:"collectionId"`
	TrackId                int       `json:"trackId"`
	ArtistName             string    `json:"artistName"`
	CollectionName         string    `json:"collectionName"`
	TrackName              string    `json:"trackName"`
	CollectionCensoredName string    `json:"collectionCensoredName"`
	TrackCensoredName      string    `json:"trackCensoredName"`
	ArtistViewUrl          string    `json:"artistViewUrl"`
	CollectionViewUrl      string    `json:"collectionViewUrl"`
	TrackViewUrl           string    `json:"trackViewUrl"`
	PreviewUrl             string    `json:"previewUrl"`
	ArtworkUrl30           string    `json:"artworkUrl30"`
	ArtworkUrl60           string    `json:"artworkUrl60"`
	ArtworkUrl100          string    `json:"artworkUrl100"`
	CollectionPrice        float32   `json:"collectionPrice"`
	TrackPrice             float32   `json:"trackPrice"`
	ReleaseDate            time.Time `json:"releaseDate"`
	CollectionExplicitness string    `json:"collectionExplicitness"`
	TrackExplicitness      string    `json:"trackExplicitness"`
	DiscCount              int       `json:"discCount"`
	DiscNumber             int       `json:"discNumber"`
	TrackCount             int       `json:"trackCount"`
	TrackNumber            int       `json:"trackNumber"`
	TrackTimeMillis        int       `json:"trackTimeMillis"`
	Country                string    `json:"country"`
	Currency               string    `json:"currency"`
	PrimaryGenreName       string    `json:"primaryGenreName"`
	IsStreamable           bool      `json:"isStreamable"`
}
