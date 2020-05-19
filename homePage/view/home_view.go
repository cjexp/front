//go:generate gobox tools/easymock

package view

import (
	"html/template"

	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/homePage/view/internal"
	"github.com/cjtoolkit/ctx"
)

type HomeView interface {
	ExecIndexView(context ctx.Context)
}

func NewHomeView(context ctx.BackgroundContext) homeView {
	return homeView{
		indexTpl:     internal.BuildIndexTemplate(context),
		errorService: loggers.GetErrorService(context),
	}
}

type homeView struct {
	indexTpl     *template.Template
	errorService loggers.ErrorService
}

func (h homeView) ExecIndexView(context ctx.Context) {
	context.SetTitle("Hello")
	h.errorService.CheckErrorAndLog(h.indexTpl.Execute(context.ResponseWriter(), context))
}
