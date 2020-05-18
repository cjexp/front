package urls

import (
	"html/template"
	"strings"
)

type FilesUrlScope struct{}

func (f FilesUrlScope) f(src, replacement string) template.HTMLAttr {
	return template.HTMLAttr(strings.Replace(src, "*replace", replacement, -1))
}

func (f FilesUrlScope) Fonts(filepath string) template.HTMLAttr { return f.f(FontsFiles, filepath) }
func (f FilesUrlScope) Javascript(filepath string) template.HTMLAttr {
	return f.f(JavascriptFiles, filepath)
}
func (f FilesUrlScope) Stylesheet(filepath string) template.HTMLAttr {
	return f.f(StylesheetFiles, filepath)
}
