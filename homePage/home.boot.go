package homePage

import (
	"github.com/cjtoolkit/ctx"
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/homePage/controller"
	"github.com/cjexp/front/urls"
)

type homeBoot struct {
	homeController controller.HomeController
	router         router.Router
}

func (b homeBoot) boot() {
	b.router.GET(urls.Index, func(context ctx.Context, _ router.Params) {
		b.homeController.Index(context)
	})
}
