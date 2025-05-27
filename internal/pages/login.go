package pages

import (
	"ostkost/go-ps-hw-fiber/views"

	"github.com/a-h/templ"
)

func LoginComponent() templ.Component {
	return views.Login()
}
