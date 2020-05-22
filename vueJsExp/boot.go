package vueJsExp

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/vueJsExp/controller"
	"github.com/cjtoolkit/ctx"
)

func Boot(context ctx.BackgroundContext) {
	bootAddress{
		addressController: controller.NewAddressController(context),
		router:            router.GetRouter(context),
	}.boot()

	bootAlert{
		alertController: controller.NewAlertController(context),
		router:          router.GetRouter(context),
	}.boot()
}
