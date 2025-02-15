package routes

import (
	"github.com/EduardoBacarin/golang-mongo-crud/src/controller/users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/id/:userId", users.FindUserById)
	r.GET("/user/email/:userEmail", users.FindUserByEmail)
	r.POST("/user", users.CreateUser)
	r.PATCH("/user/:userId", users.UpdateUser)
	r.DELETE("/user/:userId", users.DeleteUser)
}
