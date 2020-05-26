package homePage

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/homePage/controller"
	"github.com/cjtoolkit/ctx/v2"
)

func Boot(context ctx.Context) {
	homeBoot{
		homeController: controller.GetHomeController(context),
		router:         router.GetRouter(context),
	}.boot()
}
