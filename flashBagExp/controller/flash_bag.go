package controller

import (
	"github.com/cjexp/base/utility/cookie"
	"github.com/cjexp/base/utility/httpError"
	"github.com/cjexp/front/flashBagExp/view"
	"github.com/cjexp/front/urls"
	"github.com/cjtoolkit/ctx/v2"
)

type FlashBagController struct {
	view     view.FlashBagView
	flashBag cookie.FlashBag
}

func NewFlashBagController(context ctx.Context) *FlashBagController {
	return &FlashBagController{
		view:     view.NewFlashBagView(context),
		flashBag: cookie.GetFlashBag(context),
	}
}

func (c *FlashBagController) Index(context ctx.Context) {
	c.view.ExecIndex(context)
}

func (c *FlashBagController) TestFlashBag(context ctx.Context) {
	fb := c.flashBag.GetFlashBag(context)

	fb.Set("success", "Success 1")
	fb.Add("success", "Success 2")

	fb.Set("info", "Info 1")
	fb.Add("info", "Info 2")

	fb.Set("warning", "Warning 1")
	fb.Add("warning", "Warning 2")

	fb.Set("error", "Error 1")
	fb.Add("error", "Error 2")

	c.flashBag.SaveFlashBagToSession(context)
	httpError.HaltSeeOther(urls.FlashBagIndex)
}
