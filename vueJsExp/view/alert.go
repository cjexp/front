//go:generate gobox tools/easymock

package view

import (
	"html/template"

	"github.com/cjtoolkit/ctx/v2/ctxHttp"

	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/vueJsExp/view/internal"
	"github.com/cjtoolkit/ctx/v2"
)

type AlertView interface {
	ExecAlertView(context ctx.Context)
}

func NewAlertView(context ctx.Context) AlertView {
	return alertView{
		alertTemplate: internal.BuildViewAlertTemplate(context),
		errorService:  loggers.GetErrorService(context),
	}
}

type alertView struct {
	alertTemplate *template.Template
	errorService  loggers.ErrorService
}

func (a alertView) ExecAlertView(context ctx.Context) {
	ctxHttp.SetTitle(context, "Alert using VueJs")
	a.errorService.CheckErrorAndLog(a.alertTemplate.Execute(ctxHttp.Response(context), context))
}
