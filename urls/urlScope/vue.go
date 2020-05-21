package urlScope

import (
	"html/template"

	"github.com/cjexp/front/urls"
)

type Vue struct{}

func (_ Vue) Alert() template.HTMLAttr { return urls.VueAlert }
