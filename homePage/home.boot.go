package homePage

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/homePage/controller"
	"github.com/cjexp/front/urls"
	"github.com/cjtoolkit/ctx/v2"
)

type homeBoot struct {
	homeController controller.HomeController
	router         router.Router
}

func (b homeBoot) boot() {
	b.router.GET(urls.HomeIndex, func(context ctx.Context, _ router.Params) {
		b.homeController.Index(context)
	})
}
