//go:generate gobox tools/easymock

package view

import (
	"encoding/json"
	"html/template"

	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/vueJsExp/view/internal"
	"github.com/cjtoolkit/ctx/v2"
	"github.com/cjtoolkit/ctx/v2/ctxHttp"
)

type RandomHashView interface {
	ExecRandomHashView(context ctx.Context)
	ExecJsonString(context ctx.Context, data []string)
}

func NewRandomHashView(context ctx.Context) RandomHashView {
	return randomHashView{
		hashViewTemplate: internal.BuildViewRandomHash(context),
		errorService:     loggers.GetErrorService(context),
	}
}

type randomHashView struct {
	hashViewTemplate *template.Template
	errorService     loggers.ErrorService
}

func (r randomHashView) ExecRandomHashView(context ctx.Context) {
	ctxHttp.SetTitle(context, "Random Hash using VueJs")
	r.errorService.CheckErrorAndLog(r.hashViewTemplate.Execute(ctxHttp.Response(context), context))
}

func (r randomHashView) ExecJsonString(context ctx.Context, data []string) {
	r.errorService.CheckErrorAndLog(json.NewEncoder(ctxHttp.Response(context)).Encode(data))
}
