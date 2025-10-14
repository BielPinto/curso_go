package storage

import (
	"context"
	"fmt"
	"time"
)

// MockStorage for testing
type MockStorage struct {
	data map[string]string
}

func NewMockStorage() *MockStorage {
	return &MockStorage{data: make(map[string]string)}
}

func (m *MockStorage) Increment(ctx context.Context, key string) (int64, error) {
	val, exists := m.data[key]
	if !exists {
		m.data[key] = "1"
		return 1, nil
	}
	// Simple increment
	count := 0
	for _, c := range val {
		if c >= '0' && c <= '9' {
			count = count*10 + int(c-'0')
		}
	}
	count++
	m.data[key] = string(rune(count + '0')) // rough
	return int64(count), nil
}

func (m *MockStorage) SetExpiration(ctx context.Context, key string, expiration time.Duration) error {
	// Mock, do nothing
	return nil
}

func (m *MockStorage) Get(ctx context.Context, key string) (string, error) {
	if val, exists := m.data[key]; exists {
		return val, nil
	}
	return "", fmt.Errorf("key not found") // simulate redis.Nil
}

func (m *MockStorage) Delete(ctx context.Context, key string) error {
	delete(m.data, key)
	return nil
}
