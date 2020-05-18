//go:generate gobox tools/easymock

package view

import (
	"html/template"

	"github.com/cjtoolkit/ctx"
	"github.com/cjexp/base/utility/embedder"
	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/homePage/model"
	"github.com/cjexp/front/homePage/view/internal"
	"github.com/cjexp/front/master"
)

type HomeView interface {
	ExecIndexView(context ctx.Context, data model.Index)
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
		indexTpl:     buildIndexTemplate(context),
		errorService: loggers.GetErrorService(context),
	}
}

func (h homeView) ExecIndexView(context ctx.Context, data model.Index) {
	context.SetTitle("Hello World")

	type m struct {
		ctx.Context
		Local model.Index
	}

	h.errorService.CheckErrorAndLog(h.indexTpl.Execute(context.ResponseWriter(), m{
		Context: context,
		Local:   data,
	}))
}

func buildIndexTemplate(context ctx.BackgroundContext) *template.Template {
	return template.Must(master.CloneMasterTemplate(context).Parse(embedder.DecodeValueStr(internal.Index)))
}
