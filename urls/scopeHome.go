package urls

import "html/template"

type HomeUrlScope struct{}

func (_ HomeUrlScope) Index() template.HTMLAttr { return HomeIndex }
