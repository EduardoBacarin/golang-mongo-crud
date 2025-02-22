package users

import (
	"fmt"
	"net/http"

	"github.com/EduardoBacarin/golang-mongo-crud/src/configuration/rest_err/validation"
	"github.com/EduardoBacarin/golang-mongo-crud/src/controller/model/request"
	"github.com/EduardoBacarin/golang-mongo-crud/src/model"
	"github.com/gin-gonic/gin"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	fmt.Println(userRequest)
	c.String(http.StatusOK, "")
}
