package repository

import (
	"contact-list/models"

	"gorm.io/gorm"
)

type RepoUser interface {
	Create(user *models.User) error
	Get(id int) (models.User, error)
	Delete(user *models.User) error
	Gets() ([]models.User, error)
	Update(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func User(db *gorm.DB) userRepository {
	return userRepository{db: db}
}

func (u userRepository) Migration() {
	u.db.AutoMigrate(&models.User{})
}

func (u userRepository) Create(user *models.User) error {
	err := u.db.Save(user).Error

	return err
}

func (u userRepository) Get(id int) (models.User, error) {
	var user models.User

	err := u.db.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (u userRepository) Gets() ([]models.User, error) {
	var users []models.User

	err := u.db.Find(&users).Error

	return users, err
}

func (u userRepository) Delete(user *models.User) error {
	err := u.db.Delete(user).Error

	return err
}

func (u userRepository) Update(user *models.User) error {
	err := u.db.Save(user).Error

	return err
}
