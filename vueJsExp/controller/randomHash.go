package controller

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/cjexp/base/utility/loggers"
	"github.com/cjexp/front/vueJsExp/view"
	"github.com/cjtoolkit/ctx/v2"
)

type RandomHashController struct {
	view         view.RandomHashView
	errorService loggers.ErrorService
}

func NewRandomHashController(context ctx.Context) RandomHashController {
	return RandomHashController{
		view:         view.NewRandomHashView(context),
		errorService: loggers.GetErrorService(context),
	}
}

func (c RandomHashController) Index(context ctx.Context) {
	c.view.ExecRandomHashView(context)
}

func (c RandomHashController) FetchData(context ctx.Context) {
	data := make([]string, 5)
	for id, _ := range data {
		b := make([]byte, 32)
		_, err := rand.Read(b)
		c.errorService.CheckErrorAndPanic(err)
		data[id] = base64.URLEncoding.EncodeToString(b)
	}
	c.view.ExecJsonString(context, data)
}
