package flashBagExp

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/flashBagExp/controller"
	"github.com/cjtoolkit/ctx"
)

func Boot(context ctx.BackgroundContext) {
	bootFlashBagController{
		controller: controller.NewFlashBagController(context),
		router:     router.GetRouter(context),
	}.boot()
}
