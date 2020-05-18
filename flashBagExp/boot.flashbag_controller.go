package flashBagExp

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/flashBagExp/controller"
	"github.com/cjexp/front/urls"
	"github.com/cjtoolkit/ctx"
)

type bootFlashBagController struct {
	controller controller.FlashBagController
	router     router.Router
}

func (b bootFlashBagController) boot() {
	b.router.GET(urls.FlashBagIndex, func(context ctx.Context, params router.Params) {
		b.controller.Index(context)
	})

	b.router.GET(urls.FlashBagTest, func(context ctx.Context, params router.Params) {
		b.controller.TestFlashBag(context)
	})
}
