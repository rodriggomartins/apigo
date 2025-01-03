package controller

import (
	"api/src/configuration/rest_err/validation"
	"api/src/controller/model/request"
	"api/src/controller/model/response"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error tryping to marshal object, error=%s\n", err.Error())
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:       "test",
		Email:    userRequest.Email,
		Password: "testetete",
		Name:     "Rodrigo",
	}

}
