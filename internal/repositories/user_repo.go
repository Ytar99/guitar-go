package repositories

import (
	"guitar-go/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (repo *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := repo.db.Find(&users).Error
	return users, err
}
