package homePage

import (
	"github.com/cjtoolkit/ctx"
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/homePage/controller"
)

func Boot(context ctx.BackgroundContext) {
	homeBoot{
		homeController: controller.GetHomeController(context),
		router:         router.GetRouter(context),
	}.boot()
}
