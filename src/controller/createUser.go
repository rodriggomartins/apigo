package controller

import (
	"api/src/configuration/logger"
	"api/src/configuration/validation"
	"api/src/controller/model/request"
	"api/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init Createuser Controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error para validacao do usuario", err,
			zap.String("journey", "createUser"),
		)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(userRequest.Name, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Creiado com sucesso",
		zap.String("jorney", "createUser"),
	)

	c.String(http.StatusOK, "")

}
