package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cjexp/base/utility/command/param"
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/errorPage"
	"github.com/cjexp/front/fileServer"
	"github.com/cjexp/front/flashBagExp"
	"github.com/cjexp/front/homePage"
	"github.com/cjexp/front/vueJsExp"
	"github.com/cjtoolkit/ctx"
)

var build = "Undefined"

func boot() (http.Handler, param.Param) {
	context := ctx.NewBackgroundContext()
	defer ctx.ClearBackgroundContext(context)

	_param := param.GetParam(context)

	fmt.Printf("Build: %q", build)
	fmt.Println()

	errorPage.Boot(context)
	fileServer.Boot(context)
	flashBagExp.Boot(context)
	vueJsExp.Boot(context)

	homePage.Boot(context)

	fmt.Println("Booted up successfully.")
	fmt.Println("")

	return router.GetRouter(context), _param
}

func main() {
	handler, _param := boot()

	param.CheckIfTestRun(_param)

	fmt.Println("Now listening on", _param.Address)
	log.Print(http.ListenAndServe(_param.Address, handler))
}
