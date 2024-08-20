package services

import (
	"errors"
	"sync"

	"zhiming.cool/go/internal/models"
)

type UserService struct {
	mu    sync.Mutex
	users map[string]*models.User
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*models.User),
	}
}

func (s *UserService) GetUserByName(name string) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[name]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) CreateUser(name, email string) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[name]; exists {
		return nil, errors.New("user already exists")
	}

	user := &models.User{
		Name:  name,
		Email: email,
	}
	s.users[name] = user
	return user, nil
}

func (s *UserService) DeleteUser(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[name]; !exists {
		return errors.New("user not found")
	}

	delete(s.users, name)
	return nil
}

func (s *UserService) UpdateUser(name, email string) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[name]
	if !exists {
		return nil, errors.New("user not found")
	}

	user.Email = email
	return user, nil
}
