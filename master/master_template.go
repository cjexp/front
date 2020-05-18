package master

import (
	"html/template"

	"github.com/cjexp/base/utility/embedder"
	"github.com/cjexp/front/master/internal"
	"github.com/cjexp/front/master/util"
	"github.com/cjtoolkit/ctx"
)

func CloneMasterTemplate(context ctx.BackgroundContext) *template.Template {
	return template.Must(getMasterTemplate(context).Clone())
}

func getMasterTemplate(context ctx.BackgroundContext) *template.Template {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		return buildMasterTemplate(context), nil
	}).(*template.Template)
}

func buildMasterTemplate(context ctx.BackgroundContext) *template.Template {
	maps := template.FuncMap{}

	util.RegisterFlashBag(context, maps)
	util.RegisterCsrf(context, maps)

	name, tpl := "Master", embedder.DecodeValue(internal.Master)
	return template.Must(template.New(name).Funcs(maps).Parse(string(tpl)))
}
