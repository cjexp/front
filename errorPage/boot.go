package errorPage

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/errorPage/controller"
	"github.com/cjtoolkit/ctx/v2"
)

func Boot(context ctx.Context) {
	bootError(controller.NewErrorController(context), router.GetRouter(context))
}
