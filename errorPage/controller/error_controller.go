package controller

import (
	"runtime/debug"

	"github.com/cjtoolkit/ctx"
	"github.com/cjexp/base/utility/command/param"
	"github.com/cjexp/front/errorPage/model"
	"github.com/cjexp/front/errorPage/view"
)

type ErrorController struct {
	production bool
	view       view.ErrorView
}

func NewErrorController(context ctx.BackgroundContext) ErrorController {
	return ErrorController{
		production: param.GetParam(context).Production,
		view:       view.NewErrorView(context),
	}
}

func (c ErrorController) ShowError(context ctx.Context, code int, status, message string) {
	stackTrace := []byte{}
	if !c.production {
		stackTrace = debug.Stack()
	}

	c.view.ErrorTemplate(context, code, status, model.ErrorTemplateModel{
		Production: c.production,
		StackTrace: string(stackTrace),
		Message:    message,
	})
}
