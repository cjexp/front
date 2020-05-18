package urlScope

import (
	"html/template"
	"strings"

	"github.com/cjexp/front/urls"
)

type Files struct{}

func (f Files) f(src, replacement string) template.HTMLAttr {
	return template.HTMLAttr(strings.Replace(src, "*filepath", replacement, 1))
}

func (f Files) Fonts(filepath string) template.HTMLAttr { return f.f(urls.FontsFiles, filepath) }
func (f Files) Javascript(filepath string) template.HTMLAttr {
	return f.f(urls.JavascriptFiles, filepath)
}
func (f Files) Stylesheet(filepath string) template.HTMLAttr {
	return f.f(urls.StylesheetFiles, filepath)
}
