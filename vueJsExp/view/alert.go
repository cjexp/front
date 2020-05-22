//go:generate gobox tools/easymock

package view

import (
	"html/template"

	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/vueJsExp/view/internal"
	"github.com/cjtoolkit/ctx"
)

type AlertView interface {
	ExecAlertView(context ctx.Context)
}

func NewAlertView(context ctx.BackgroundContext) AlertView {
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
	context.SetTitle("Vue Alert")
	a.errorService.CheckErrorAndLog(a.alertTemplate.Execute(context.ResponseWriter(), context))
}
