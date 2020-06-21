//go:generate gobox tools/easymock

package view

import (
	"html/template"

	"github.com/cjtoolkit/ctx/v2/ctxHttp"

	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/homePage/view/internal"
	"github.com/cjtoolkit/ctx/v2"
)

type HomeView interface {
	ExecIndexView(context ctx.Context)
}

func NewHomeView(context ctx.Context) HomeView {
	return &homeView{
		indexTpl:     internal.BuildIndexTemplate(context),
		errorService: loggers.GetErrorService(context),
	}
}

type homeView struct {
	indexTpl     *template.Template
	errorService loggers.ErrorService
}

func (h *homeView) ExecIndexView(context ctx.Context) {
	ctxHttp.SetTitle(context, "Hello")
	h.errorService.CheckErrorAndLog(h.indexTpl.Execute(ctxHttp.Response(context), context))
}
