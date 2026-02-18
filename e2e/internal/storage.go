package internal

import "sync"

type Storage struct {
	// In-memory artifact store for passing data between tests.
	artifacts map[string]string
	mu        sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		artifacts: make(map[string]string),
	}
}

func (s *Storage) SetArtifact(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.artifacts == nil {
		s.artifacts = make(map[string]string)
	}
	s.artifacts[key] = value
}

func (s *Storage) GetArtifact(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, exists := s.artifacts[key]
	return value, exists
}
