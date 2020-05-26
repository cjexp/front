package master

import (
	"html/template"

	"github.com/cjexp/front/urls/urlScope"

	"github.com/cjexp/base/utility/embedder"
	"github.com/cjexp/front/master/internal"
	"github.com/cjexp/front/master/util"
	"github.com/cjtoolkit/ctx/v2"
)

func CloneMasterTemplate(context ctx.Context) *template.Template {
	return template.Must(getMasterTemplate(context).Clone())
}

func getMasterTemplate(context ctx.Context) *template.Template {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		return buildMasterTemplate(context), nil
	}).(*template.Template)
}

func buildMasterTemplate(context ctx.Context) *template.Template {
	maps := template.FuncMap{}

	util.RegisterTitle(maps)
	util.RegisterFlashBag(context, maps)
	util.RegisterCsrf(context, maps)
	urlScope.RegisterUrlScope(maps)

	name, tpl := "Master", embedder.DecodeValue(internal.Master)
	return template.Must(template.New(name).Funcs(maps).Parse(string(tpl)))
}
