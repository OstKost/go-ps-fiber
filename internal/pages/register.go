package pages

import (
	"ostkost/go-ps-hw-fiber/views"

	"github.com/a-h/templ"
)

func RegisterComponent() templ.Component {
	return views.Register()
}
