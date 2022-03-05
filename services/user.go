package services

import (
	"contact-list/models"
	"contact-list/repository"
	requestUser "contact-list/request/user"
	"strconv"
)

type ServiceUser interface {
	Create(req requestUser.Create) string
	Get(id string) (models.User, error)
	Delete(id int) error
	Gets() ([]models.User, error)
	Update(req requestUser.Update) error
}

type userService struct {
	repo repository.RepoUser
}

func User(repo repository.RepoUser) userService {
	return userService{repo: repo}
}

func (s userService) Create(req requestUser.Create) string {
	user := models.User{
		Name:        req.Name,
		Phone:       req.Phone,
		Description: req.Description,
	}

	err := s.repo.Create(&user)
	if err != nil {
		return err.Error()
	}
	return "başarılı"
}

func (s userService) Get(id string) (models.User, error) {
	Id, _ := strconv.Atoi(id)

	data, err := s.repo.Get(Id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s userService) Gets() ([]models.User, error) {
	data, err := s.repo.Gets()

	return data, err
}

func (s userService) Delete(id int) error {

	data, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	err = s.repo.Delete(&data)
	if err != nil {
		return err
	}
	return nil
}

func (s userService) Update(req requestUser.Update) error {
	data, err := s.repo.Get(req.ID)
	if err != nil {
		return err
	}

	data.Name = req.Name
	data.Description = req.Description
	data.Phone = req.Phone

	err = s.repo.Update(&data)
	if err != nil {
		return err
	}

	return nil
}
