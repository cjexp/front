package vueJsExp

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/urls"
	"github.com/cjexp/front/vueJsExp/controller"
	"github.com/cjtoolkit/ctx/v2"
)

type bootAddress struct {
	addressController *controller.AddressController
	router            router.Router
}

func (b bootAddress) boot() {
	b.router.GET(urls.VueAddress, func(context ctx.Context) {
		b.addressController.Index(context)
	})
}
