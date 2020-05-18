package urlScope

import (
	"html/template"

	"github.com/cjexp/front/urls"
	"github.com/cjexp/front/urls/urlScope/internal"
)

type Files struct{}

func (_ Files) Fonts(filepath string) template.HTMLAttr {
	return internal.ReplaceFilePath(urls.FontsFiles, filepath)
}
func (_ Files) Javascript(filepath string) template.HTMLAttr {
	return internal.ReplaceFilePath(urls.JavascriptFiles, filepath)
}
func (_ Files) Stylesheet(filepath string) template.HTMLAttr {
	return internal.ReplaceFilePath(urls.StylesheetFiles, filepath)
}
