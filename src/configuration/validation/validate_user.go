package validation

import (
	"encoding/json"
	"errors"

	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationErr error) *Error.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return Error.NewBadRequestError("invalid field type")
	} else if errors.As(validationErr, &jsonValidationError) {
		errorsCauses := []Error.Cause{}

		for _, e := range validationErr.(validator.ValidationErrors) {

			cause := Error.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)

		}

		return Error.NewBadRequestValidationError("Some field are invalid", errorsCauses)
	} else {
		return Error.NewBadRequestError("Erro trying to convert fields")
	}
}
