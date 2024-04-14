package repository

import (
	"api-survey-go/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: Repository[entity.User]{DB: db},
		DB:         db,
	}
}

func (repo *UserRepository) GetUsers() ([]entity.User, error) {
	var users []entity.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetUserByEmail(user *entity.User, email string) error {
	return repo.DB.Joins("Role").Joins("Client").Where("email = ?", email).First(&user).Error
}

func (repo *UserRepository) GetUserById(user *entity.User, id string) error {
	return repo.DB.Joins("Role").Joins("Client").First(&user, id).Error
}
