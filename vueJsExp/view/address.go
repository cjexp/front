package view

import (
	"html/template"

	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/vueJsExp/view/internal"
	"github.com/cjtoolkit/ctx"
)

type AddressView interface {
	ExecAddressView(context ctx.Context)
}

func NewAddressView(context ctx.BackgroundContext) AddressView {
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
	context.SetTitle("Vue Address")
	a.errorService.CheckErrorAndLog(a.addressTemplate.Execute(context.ResponseWriter(), context))
}
