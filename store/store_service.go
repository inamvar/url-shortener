package store

// save and retrieve urls
type StoreService interface {
	SaveUrlMapping(shortUrl string, originalUrl string)
	RetrieveOriginalUrl(shortUrl string) string
}
