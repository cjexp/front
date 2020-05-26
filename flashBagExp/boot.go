package flashBagExp

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/flashBagExp/controller"
	"github.com/cjtoolkit/ctx/v2"
)

func Boot(context ctx.Context) {
	bootFlashBagController{
		controller: controller.NewFlashBagController(context),
		router:     router.GetRouter(context),
	}.boot()
}
