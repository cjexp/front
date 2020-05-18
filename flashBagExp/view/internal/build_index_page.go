package internal

import (
	"html/template"

	"github.com/cjexp/base/utility/embedder"
	"github.com/cjexp/front/flashBagExp/view/internal/internal"
	"github.com/cjexp/front/master"
	"github.com/cjtoolkit/ctx"
)

func BuildIndexPage(context ctx.BackgroundContext) *template.Template {
	return template.Must(master.CloneMasterTemplate(context).Parse(embedder.DecodeValueStr(internal.FlashBagIndex)))
}
