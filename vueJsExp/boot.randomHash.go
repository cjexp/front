package vueJsExp

import (
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/urls"
	"github.com/cjexp/front/vueJsExp/controller"
	"github.com/cjtoolkit/ctx/v2"
)

type bootRandomHash struct {
	controller controller.RandomHashController
	router     router.Router
}

func (b bootRandomHash) boot() {
	b.router.GET(urls.VueRandomHash, func(context ctx.Context) {
		b.controller.Index(context)
	})

	b.router.PUT(urls.VueRandomHash, func(context ctx.Context) {
		b.controller.FetchData(context)
	})
}
