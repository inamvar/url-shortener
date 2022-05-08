package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionAndRetrieval(t *testing.T) {
	originalUrl := "long-url-goes-here"
	shortURL := "tiny-url"

	store := NewInMemoryStore()
	// Persist data mapping
	store.SaveUrlMapping(shortURL, originalUrl)
	// Retrieve initial URL
	retrievedUrl := store.RetrieveOriginalUrl(shortURL)

	assert.Equal(t, originalUrl, retrievedUrl)
}
