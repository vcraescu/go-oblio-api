package token

import (
	"context"
	"errors"
	"sync"
	"time"
)

type InMemStorage struct {
	token     string
	expiresAt time.Time
	mu        sync.Mutex
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{}
}

func (s *InMemStorage) Set(_ context.Context, token string, ttl time.Duration) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.token = token
	s.expiresAt = time.Now().Add(ttl)

	return nil
}

func (s *InMemStorage) Get(_ context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.token == "" || s.expiresAt.Before(time.Now()) {
		return "", errors.New("token not set")
	}

	return s.token, nil
}
