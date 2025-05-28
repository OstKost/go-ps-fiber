package pages

import (
	"ostkost/go-ps-hw-fiber/views"

	"github.com/a-h/templ"
)

func NotFoundComponent() templ.Component {
	return views.NotFound()
}
