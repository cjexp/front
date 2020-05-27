//go:generate gobox tools/easymock

package view

import (
	"html/template"

	"github.com/cjtoolkit/ctx/v2/ctxHttp"

	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/vueJsExp/view/internal"
	"github.com/cjtoolkit/ctx/v2"
)

type AddressView interface {
	ExecAddressView(context ctx.Context)
}

func NewAddressView(context ctx.Context) AddressView {
	return addressView{
		addressTemplate: internal.BuildViewAddressTemplate(context),
		errorService:    loggers.GetErrorService(context),
	}
}

type addressView struct {
	addressTemplate *template.Template
	errorService    loggers.ErrorService
}

func (a addressView) ExecAddressView(context ctx.Context) {
	ctxHttp.SetTitle(context, "Address Label Generator using VueJs")
	a.errorService.CheckErrorAndLog(a.addressTemplate.Execute(ctxHttp.Response(context), context))
}
