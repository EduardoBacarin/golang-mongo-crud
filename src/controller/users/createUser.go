package users

import (
	"fmt"

	"github.com/EduardoBacarin/golang-mongo-crud/src/configuration/rest_err/validation"
	"github.com/EduardoBacarin/golang-mongo-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
