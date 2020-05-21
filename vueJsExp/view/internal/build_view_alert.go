package internal

import (
	"html/template"

	"github.com/cjexp/base/utility/embedder"
	"github.com/cjexp/front/master"
	"github.com/cjexp/front/vueJsExp/view/internal/internal"
	"github.com/cjtoolkit/ctx"
)

func BuildViewAlertTemplate(context ctx.BackgroundContext) *template.Template {
	return template.Must(master.CloneMasterTemplate(context).Parse(embedder.DecodeValueStr(internal.VueAlert)))
}
