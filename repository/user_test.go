package repository

import (
	"contact-list/config"
	"contact-list/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepo(t *testing.T) {
	_, err := config.Connect()
	assert.Equal(t, err, nil)
}

func TestUserCreate(t *testing.T) {
	db, err := config.Connect()
	assert.Equal(t, err, nil)

	repo := User(db)

	user := models.User{
		Name:        "Test1",
		Phone:       "1",
		Description: "Test1 Des",
	}

	err = repo.Create(&user)
	assert.Equal(t, err, nil)
}

func TestUserGet(t *testing.T) {
	db, err := config.Connect()
	assert.Equal(t, err, nil)
	t.Log("pass db test ...")

	repo := User(db)

	user := models.User{
		Name:        "Test1",
		Phone:       "1",
		Description: "Test1 Des",
	}

	err = repo.Create(&user)
	assert.Equal(t, err, nil)
	t.Log("pass create user test ...")

	getUser, err := repo.Get(int(user.ID))
	assert.Equal(t, err, nil)
	assert.Equal(t, getUser.Name, user.Name)
	t.Log("pass get user test ...")
}

func TestUserGets(t *testing.T) {
	db, err := config.Connect()
	assert.Equal(t, err, nil)
	t.Log("pass db test ...")

	repo := User(db)

	user := models.User{
		Name:        "Test2",
		Phone:       "2",
		Description: "Test2 Des",
	}

	err = repo.Create(&user)
	assert.Equal(t, err, nil)
	t.Log("pass create user test ...")

	_, err = repo.Gets()
	assert.Equal(t, err, nil)
	t.Log("pass get user test ...")
}

func TestUserDelete(t *testing.T) {
	db, err := config.Connect()
	assert.Equal(t, err, nil)
	t.Log("pass db test ...")

	repo := User(db)

	user := models.User{
		Name:        "Test2",
		Phone:       "2",
		Description: "Test2 Des",
	}

	err = repo.Create(&user)
	assert.Equal(t, err, nil)
	t.Log("pass create user test ...")

	err = repo.Delete(&user)
	assert.Equal(t, err, nil)
	t.Log("pass delete user test ...")
}

func TestUserUpdate(t *testing.T) {
	db, err := config.Connect()
	assert.Equal(t, err, nil)
	t.Log("pass db test ...")

	repo := User(db)

	user := models.User{
		Name:        "Test2",
		Phone:       "2",
		Description: "Test2 Des",
	}

	err = repo.Create(&user)
	assert.Equal(t, err, nil)
	t.Log("pass create user test ...")

	user.Name = "update name"
	user.Phone = "3"
	user.Description = "update des"

	err = repo.Update(&user)
	assert.Equal(t, err, nil)
	t.Log("pass update user test ...")
}
