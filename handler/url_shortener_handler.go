package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/inamvar/url-shortener/shortener"
	"github.com/inamvar/url-shortener/store"
)

func UrlShortenerHandler(s store.StoreService, router fiber.Router) {

	router.Post("/generate", createShortUrl(s))
	router.Get("/s/:shortUrl", handleShortUrlRedirect(s))
}

func createShortUrl(s store.StoreService) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		model := new(CreateShortUrlRequest)
		if err := c.BodyParser(model); err != nil {
			return c.Status(http.StatusBadRequest).JSON(err)
		}

		shortUrl, err := shortener.GenerateShortLink(model.LongUrl)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err)
		}
		s.SaveUrlMapping(shortUrl, model.LongUrl)
		shortUrl = c.BaseURL() + "/s/" + shortUrl
		return c.Status(http.StatusCreated).JSON(&CreateShortUrlResponse{
			LongUrl:  model.LongUrl,
			ShortUrl: shortUrl,
		})
	}
}

func handleShortUrlRedirect(s store.StoreService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		shortUrl := c.Params("shortUrl")
		lognUrl := s.RetrieveOriginalUrl(shortUrl)
		return c.Redirect(lognUrl, 302)
	}
}
