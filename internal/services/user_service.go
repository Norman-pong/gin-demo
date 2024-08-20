package services

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"zhiming.cool/go/internal/models"
)

type UserService struct {
	DB *gorm.DB
}

// todo: Manage the database through middleware
func NewUserService(dsn string) *UserService {
	var db *gorm.DB
	var err error

	if dsn == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic("failed to get working directory: " + err.Error())
		}

		dsn = filepath.Join(wd, "db", "dev.db")

		if err := os.MkdirAll(filepath.Dir(dsn), os.ModePerm); err != nil {
			log.Fatalf("failed to create database directory: %v", err)
		}
	}

	// connect database
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: %v")
	}

	// Migrate
	db.AutoMigrate(&models.User{})

	return &UserService{DB: db}
}

func (s *UserService) GetUserById(cid string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("id = ?", cid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CreateUser(name, email string) (*models.User, error) {
	user := &models.User{
		Name:  name,
		Email: email,
	}
	if err := s.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(cid string) (*models.User, error) {
	var user models.User

	if err := s.DB.Where("id = ?", cid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	if err := s.DB.Delete(&user).Error; err != nil {
		return nil, errors.New("user delete error")
	}
	return &user, nil
}

func (s *UserService) UpdateUser(cid, email string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("id = ?", cid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	user.Email = email
	if err := s.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
