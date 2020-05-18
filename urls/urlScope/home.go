package urlScope

import (
	"html/template"

	"github.com/cjexp/front/urls"
)

type Home struct{}

func (_ Home) Index() template.HTMLAttr { return urls.HomeIndex }
