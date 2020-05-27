package internal

import (
	"html/template"

	"github.com/cjexp/base/utility/embedder"
	"github.com/cjexp/front/master"
	"github.com/cjexp/front/vueJsExp/view/internal/internal"
	"github.com/cjtoolkit/ctx/v2"
)

func BuildViewRandomHash(context ctx.Context) *template.Template {
	return template.Must(master.CloneMasterTemplate(context).Parse(embedder.DecodeValueStr(internal.VueRandomHash)))
}
