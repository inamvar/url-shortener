package handler

type CreateShortUrlRequest struct {
	LongUrl string `json:"long_url"`
}

type CreateShortUrlResponse struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}
