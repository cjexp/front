package controller

import (
	"github.com/cjexp/front/vueJsExp/view"
	"github.com/cjtoolkit/ctx/v2"
)

type AddressController struct {
	view view.AddressView
}

func NewAddressController(context ctx.Context) AddressController {
	return AddressController{
		view: view.NewAddressView(context),
	}
}

func (c AddressController) Index(context ctx.Context) {
	c.view.ExecAddressView(context)
}
