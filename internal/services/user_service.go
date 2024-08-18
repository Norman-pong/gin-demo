package services

import (
	"errors"
	"fmt"
	"sync"

	"zhiming.cool/go/pkg/utils" // Add this line to import the package containing GenerateRandomString
)

type User struct {
	Name  string
	Email string
}

type UserService struct {
	mu    sync.Mutex
	users map[string]*User
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*User),
	}
}

func (s *UserService) GetUserByName(name string) (*User, error) {
	str, err := utils.GenerateRandomString(2)

	if(err != nil){
		return nil, errors.New("error generating random string")
	}
	fmt.Printf("Random string: %s\n", str)

	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[name]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) CreateUser(name, email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[name]; exists {
		return nil, errors.New("user already exists")
	}

	user := &User{
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

func (s *UserService) UpdateUser(name, email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[name]
	if !exists {
		return nil, errors.New("user not found")
	}

	user.Email = email
	return user, nil
}