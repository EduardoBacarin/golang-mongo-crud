package validation

import (
	"encoding/json"
	"errors"

	"github.com/EduardoBacarin/golang-mongo-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_error error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_error, &jsonErr) {
		return rest_err.NewBadRequestErr("Invalid field type")
	} else if errors.As(validation_error, &jsonValidationError) {
		errorCauses := []rest_err.Causes{}

		for _, e := range validation_error.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}
		return rest_err.NewBadRequestValidationErr("Some fields are invalid", errorCauses)
	} else {
		return rest_err.NewBadRequestErr("Error trying to translate fields")
	}

}
