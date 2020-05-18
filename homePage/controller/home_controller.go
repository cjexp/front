package controller

import (
	"github.com/cjtoolkit/ctx"
	"github.com/cjexp/front/homePage/model"
	"github.com/cjexp/front/homePage/view"
)

type HomeController struct {
	view view.HomeView
}

func GetHomeController(context ctx.BackgroundContext) HomeController {
	return HomeController{view: view.GetHomeView(context)}
}

func (h HomeController) Index(context ctx.Context) {
	h.view.ExecIndexView(context, model.Index{Message: "Hello World"})
}
