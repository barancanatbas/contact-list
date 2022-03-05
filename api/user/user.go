package user

import (
	"contact-list/config"
	"contact-list/handler"
	"contact-list/repository"
	"contact-list/services"
	"log"
)

func Init() handler.HandlerUser {
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.User(db)
	repo.Migration()
	services := services.User(repo)
	handler := handler.User(services)

	return handler
}
