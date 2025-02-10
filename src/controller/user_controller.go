package controller

import (
	"api/src/model/service"

	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	DeleteById(c *gin.Context)
	CreateUser(c *gin.Context)
	EditUserById(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
