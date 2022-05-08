package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inamvar/url-shortener/handler"
	"github.com/inamvar/url-shortener/store"
)

func main() {
	app := fiber.New()

	inMemoryStore := store.NewInMemoryStore()
	router := app.Group("")
	handler.UrlShortenerHandler(inMemoryStore, router)
	app.Listen(":3000")
}
