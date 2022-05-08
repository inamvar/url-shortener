package store

import "sync"

type InMemoryStore struct {
	sync.RWMutex
	data map[string]string
}

// RetrieveOriginalUrl implements StorageService
func (s *InMemoryStore) RetrieveOriginalUrl(shortUrl string) string {
	s.RLock()
	defer s.RUnlock()
	return s.data[shortUrl]
}

// SaveUrlMapping implements StorageService
func (s *InMemoryStore) SaveUrlMapping(shortUrl string, originalUrl string) {
	s.Lock()
	defer s.Unlock()
	s.data[shortUrl] = originalUrl
}

func NewInMemoryStore() StoreService {
	store := &InMemoryStore{}
	store.data = make(map[string]string)
	return store
}
