package repository

import (
	"context"
	"errors"
	"sync"
	"GoLang_Tutorial/internal/models"
)

type memoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*models.User
}

func NewMemoryUserRepository() UserRepository {
	return &memoryUserRepository{
		users: make(map[string]*models.User),
	}
}

func (m *memoryUserRepository) Create(ctx context.Context, user *models.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, u := range m.users {
		if u.Username == user.Username {
			return errors.New("username đã tồn tại")
		}
	}

	m.users[user.ID] = user
	return nil
}

func (m *memoryUserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, u := range m.users {
		if u.Username == username {
			return u, nil
		}
	}

	return nil, errors.New("không tìm thấy user")
}
