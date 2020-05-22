package urlScope

import (
	"html/template"

	"github.com/cjexp/front/urls"
)

type Vue struct{}

func (_ Vue) Alert() template.HTMLAttr   { return urls.VueAlert }
func (_ Vue) Address() template.HTMLAttr { return urls.VueAddress }
