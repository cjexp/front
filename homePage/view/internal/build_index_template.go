package internal

import (
	"html/template"

	"github.com/cjexp/base/utility/embedder"
	"github.com/cjexp/front/homePage/view/internal/internal"
	"github.com/cjexp/front/master"
	"github.com/cjtoolkit/ctx/v2"
)

func BuildIndexTemplate(context ctx.Context) *template.Template {
	return template.Must(master.CloneMasterTemplate(context).Parse(embedder.DecodeValueStr(internal.Index)))
}
