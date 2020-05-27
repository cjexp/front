package vueJsExp

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/vueJsExp/controller"
	"github.com/cjtoolkit/ctx/v2"
)

func Boot(context ctx.Context) {
	bootAddress{
		addressController: controller.NewAddressController(context),
		router:            router.GetRouter(context),
	}.boot()

	bootAlert{
		alertController: controller.NewAlertController(context),
		router:          router.GetRouter(context),
	}.boot()

	bootRandomHash{
		controller: controller.NewRandomHashController(context),
		router:     router.GetRouter(context),
	}.boot()
}
