package vueJsExp

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/urls"
	"github.com/cjexp/front/vueJsExp/controller"
	"github.com/cjtoolkit/ctx"
)

type bootAlertController struct {
	alertController controller.AlertController
	router          router.Router
}

func (b bootAlertController) boot() {
	b.router.GET(urls.VueAlert, func(context ctx.Context, params router.Params) {
		b.alertController.Index(context)
	})
}
