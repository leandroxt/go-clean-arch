package outbound

type GiphyImagesOriginalResponse struct {
	URL string `json:"url"`
}

type GiphyImagesResponse struct {
	Original GiphyImagesOriginalResponse `json:"original"`
}

type GiphyDataResponse struct {
	ID     string              `json:"id"`
	Images GiphyImagesResponse `json:"images"`
}

type GiphyResponse struct {
	Data GiphyDataResponse `json:"data"`
}
