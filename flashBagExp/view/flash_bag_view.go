//go:generate gobox tools/easymock

package view

import (
	"html/template"

	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/flashBagExp/view/internal"
	"github.com/cjexp/front/urls"
	"github.com/cjtoolkit/ctx"
)

type FlashBagView interface {
	ExecIndex(ctx ctx.Context)
}

func GetFlashBagView(context ctx.BackgroundContext) FlashBagView {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		return initFlashBagView(context), nil
	}).(FlashBagView)
}

type flashBagView struct {
	indexTpl     *template.Template
	errorService loggers.ErrorService
}

func initFlashBagView(context ctx.BackgroundContext) FlashBagView {
	return flashBagView{
		indexTpl:     internal.BuildIndexPage(context),
		errorService: loggers.GetErrorService(context),
	}
}

func (f flashBagView) ExecIndex(context ctx.Context) {
	context.SetTitle("Flash Bag Test")

	type local struct {
		TestUrl template.HTMLAttr
	}

	type c struct {
		ctx.Context

		Local local
	}

	f.errorService.CheckErrorAndLog(f.indexTpl.Execute(context.ResponseWriter(), c{
		Context: context,
		Local: local{
			TestUrl: urls.FlashBagTest,
		},
	}))
}
