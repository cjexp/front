package fileServer

import (
	"net/http"
	"os"

	"github.com/cjexp/base/utility/command/param"
	"github.com/cjexp/base/utility/router"
	"github.com/cjexp/front/urls"
	"github.com/cjtoolkit/ctx/v2"
	"github.com/cjtoolkit/zipfs"
)

func Boot(context ctx.Context) {
	r := router.GetRouter(context)

	if param.GetParam(context).Production == false {
		if _, err := os.Stat("asset/live"); err == nil {
			serveFiles(r, http.Dir("asset/live"))
			return
		}
	}

	serveFiles(r, zipfs.InitZipFs("asset.zip"))
}

func serveFiles(r router.Router, fs http.FileSystem) {
	r.ServeFiles(urls.FontsFiles, zipfs.Prefix("/fonts", fs))
	r.ServeFiles(urls.JavascriptFiles, zipfs.Prefix("/javascript", fs))
	r.ServeFiles(urls.StylesheetFiles, zipfs.Prefix("/stylesheets", fs))
}
