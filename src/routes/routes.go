package routes

import (
	"api/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/userid/:userId", controller.FindUserById)
	r.GET("/email/:userEmail", controller.FindUserByEmail)
	r.POST("/create", controller.CreateUser)
	r.PUT("/edit/:userId", controller.EditUserById)
	r.DELETE("/excluiruser/:userId", controller.DeleteById)

}
