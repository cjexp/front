package urls

import (
	"html/template"
)

type UrlScope struct {
	Home     HomeUrlScope
	FlashBag FlashBagUrlScope
	Files    FilesUrlScope
}

func RegisterUrlScope(m template.FuncMap) {
	m["urls"] = func() UrlScope { return UrlScope{} }
}
