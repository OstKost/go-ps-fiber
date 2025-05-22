package validator

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/validate"
)

func FormatErrors(errors *validate.Errors) string {
	msg := "Ошибки при отправке формы:\n"
	for _, value := range errors.Errors {
		msg += fmt.Sprintf("%s\n", strings.Join(value, ", "))
	}
	return msg
}
