package util

import (
	"html/template"

	"github.com/cjexp/base/utility/csrf"
	"github.com/cjtoolkit/ctx/v2"
)

func RegisterCsrf(context ctx.Context, m template.FuncMap) {
	_csrfControler := csrf.GetController(context)
	m["csrf"] = func(context ctx.Context) csrf.Data { return _csrfControler.GetCsrfData(context) }
}
