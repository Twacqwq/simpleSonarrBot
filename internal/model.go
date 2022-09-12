package internal

//剧集信息
type Series struct {
	Title string `json:"title"`
	SortTitle string `json:"sortTitle"`
	IsEnded bool `json:"ended"`
	Overview string `json:"overview"`
	PreviousAiring string `json:"previousAiring"`
	Network string `json:"network"`
	AirTime string `json:"airTime"`
	Images []Image
	Seasons string
	Year int `json:"year"`
	TvdbId int `json:"tvdbId"`
	FirstAired string `json:"firstAired"`
	SeriesId int `json:"id"`
}

type Image struct {
	CoverType string `json:"coverType"`
	URL string `json:"url"`
	RemoteURL string `json:"remoteUrl"`
}

type Season struct {
	SeasonNumber int `json:"seasonNumber"`
	Monitored bool `json:"monitored"`
	Statistics SeasonInfo `json:"statistics"`
}

type SeasonInfo struct {
	PreviousAiring string `json:"previousAiring"`
	EpisodeFileCount int `json:"episodeFileCount"`
	EpisodeCount int `json:"episodeCount"`
	TotalEpisodeCount int `json:"totalepisodeCount"`
}