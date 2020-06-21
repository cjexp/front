package controller

import (
	"github.com/cjexp/front/homePage/view"
	"github.com/cjtoolkit/ctx/v2"
)

type HomeController struct {
	view view.HomeView
}

func GetHomeController(context ctx.Context) *HomeController {
	return &HomeController{view: view.NewHomeView(context)}
}

func (h *HomeController) Index(context ctx.Context) {
	h.view.ExecIndexView(context)
}
