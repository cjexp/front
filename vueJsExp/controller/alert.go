package controller

import (
	"github.com/cjexp/front/vueJsExp/view"
	"github.com/cjtoolkit/ctx/v2"
)

type AlertController struct {
	view view.AlertView
}

func NewAlertController(context ctx.Context) AlertController {
	return AlertController{
		view: view.NewAlertView(context),
	}
}

func (c AlertController) Index(context ctx.Context) {
	c.view.ExecAlertView(context)
}
