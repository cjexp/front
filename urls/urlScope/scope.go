package urlScope

import (
	"html/template"
)

type UrlScope struct {
	Home     Home
	FlashBag FlashBag
	Files    Files
	Vue      Vue
}

func RegisterUrlScope(m template.FuncMap) {
	m["urls"] = func() UrlScope { return UrlScope{} }
}
