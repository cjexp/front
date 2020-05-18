package errorPage

import (
	"github.com/cjtoolkit/ctx"
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/errorPage/controller"
)

func Boot(context ctx.BackgroundContext) {
	bootError(controller.NewErrorController(context), router.GetRouter(context))
}
