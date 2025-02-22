package model

import (
	"fmt"

	"github.com/EduardoBacarin/golang-mongo-crud/src/configuration/rest_err"
)

func (ud *UserDomain) DeleteUser(string) *rest_err.RestErr {
	ud.EncryptPassword()
	fmt.Println(ud)
	return nil
}
