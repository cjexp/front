//go:generate gobox tools/easymock

package view

import (
	"html/template"

	"github.com/cjtoolkit/ctx/v2/ctxHttp"

	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/flashBagExp/view/internal"
	"github.com/cjtoolkit/ctx/v2"
)

type FlashBagView interface {
	ExecIndex(ctx ctx.Context)
}

func NewFlashBagView(context ctx.Context) FlashBagView {
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
	ctxHttp.SetTitle(context, "Flash Bag Test")

	f.errorService.CheckErrorAndLog(f.indexTpl.Execute(ctxHttp.Response(context), context))
}
