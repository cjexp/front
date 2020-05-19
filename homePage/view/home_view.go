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

func GetHomeView(context ctx.BackgroundContext) HomeView {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		return initHomeView(context), nil
	}).(HomeView)
}

type homeView struct {
	indexTpl     *template.Template
	errorService loggers.ErrorService
}

func initHomeView(context ctx.BackgroundContext) homeView {
	return homeView{
		indexTpl:     internal.BuildIndexTemplate(context),
		errorService: loggers.GetErrorService(context),
	}
}

func (h homeView) ExecIndexView(context ctx.Context) {
	context.SetTitle("Hello")
	h.errorService.CheckErrorAndLog(h.indexTpl.Execute(context.ResponseWriter(), context))
}
