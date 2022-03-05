package handler

import (
	requestUser "contact-list/request/user"
	"contact-list/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Gets(c *gin.Context)
	Update(c *gin.Context)
}

type handlerUser struct {
	services services.ServiceUser
}

func User(services services.ServiceUser) handlerUser {
	return handlerUser{services: services}
}

func (h handlerUser) Create(c *gin.Context) {
	var req requestUser.Create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := h.services.Create(req)

	c.JSON(http.StatusOK, gin.H{"message": data})
}

func (h handlerUser) Get(c *gin.Context) {
	id := c.Param("id")

	data, err := h.services.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	c.JSON(200, gin.H{
		"message": "ok",
		"data":    data,
	})
}

func (h handlerUser) Gets(c *gin.Context) {

	data, err := h.services.Gets()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(200, gin.H{
		"message": "pong",
		"data":    data,
	})
}

func (h handlerUser) Delete(c *gin.Context) {
	var req requestUser.Delete

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err := h.services.Delete(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(200, gin.H{
		"message": "Silme başarılı",
	})
}

func (h handlerUser) Update(c *gin.Context) {
	var req requestUser.Update
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err := h.services.Update(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(200, gin.H{
		"message": "Güncelleme başarılı",
	})
}
