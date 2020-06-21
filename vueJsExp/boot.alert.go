package vueJsExp

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/urls"
	"github.com/cjexp/front/vueJsExp/controller"
	"github.com/cjtoolkit/ctx/v2"
)

type bootAlert struct {
	alertController *controller.AlertController
	router          router.Router
}

func (b bootAlert) boot() {
	b.router.GET(urls.VueAlert, func(context ctx.Context) {
		b.alertController.Index(context)
	})
}
