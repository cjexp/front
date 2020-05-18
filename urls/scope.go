package urls

import (
	"html/template"
	"strings"
)

type UrlScope struct {
	Home     Home
	FlashBag FlashBag
	Files    Files
}

func RegisterUrlScope(m template.FuncMap) {
	m["urls"] = func() UrlScope { return UrlScope{} }
}

type Home struct{}

func (_ Home) Index() template.HTMLAttr { return HomeIndex }

type FlashBag struct{}

func (_ FlashBag) Index() template.HTMLAttr { return FlashBagIndex }
func (_ FlashBag) Test() template.HTMLAttr  { return FlashBagTest }

type Files struct{}

func replaceFilePath(src, replacement string) template.HTMLAttr {
	return template.HTMLAttr(strings.Replace(src, "*replace", replacement, -1))
}

func (_ Files) Fonts(filepath string) template.HTMLAttr { return replaceFilePath(FontsFiles, filepath) }
func (_ Files) Javascript(filepath string) template.HTMLAttr {
	return replaceFilePath(JavascriptFiles, filepath)
}
func (_ Files) Stylesheet(filepath string) template.HTMLAttr {
	return replaceFilePath(StylesheetFiles, filepath)
}
