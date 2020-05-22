//go:generate gobox tools/easymock

package view

import (
	"html/template"

	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/flashBagExp/view/internal"
	"github.com/cjtoolkit/ctx"
)

type FlashBagView interface {
	ExecIndex(ctx ctx.Context)
}

func NewFlashBagView(context ctx.BackgroundContext) FlashBagView {
	return flashBagView{
		indexTpl:     internal.BuildIndexPage(context),
		errorService: loggers.GetErrorService(context),
	}
}

type flashBagView struct {
	indexTpl     *template.Template
	errorService loggers.ErrorService
}

func (f flashBagView) ExecIndex(context ctx.Context) {
	context.SetTitle("Flash Bag Test")

	f.errorService.CheckErrorAndLog(f.indexTpl.Execute(context.ResponseWriter(), context))
}
