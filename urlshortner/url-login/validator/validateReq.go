package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateRequestBody(requestBody any) string {
	var (
		err    error
		errMsg string
	)

	validate := validator.New()
	err = validate.Struct(requestBody)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed on the '%s' rule\n", e.Field(), e.Tag())
			errMsg = fmt.Sprintf("Field '%s' failed on the '%s' rule", e.Field(), e.Tag())
			break
		}
		return errMsg
	}

	return ""
}
