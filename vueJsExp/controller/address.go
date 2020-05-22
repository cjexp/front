package controller

import (
	"github.com/cjexp/front/vueJsExp/view"
	"github.com/cjtoolkit/ctx"
)

type AddressController struct {
	view view.AddressView
}

func NewAddressController(context ctx.BackgroundContext) AddressController {
	return AddressController{
		view: view.NewAddressView(context),
	}
}

func (c AddressController) Index(context ctx.Context) {
	c.view.ExecAddressView(context)
}
