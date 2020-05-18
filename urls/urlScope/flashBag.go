package urlScope

import (
	"html/template"

	"github.com/cjexp/front/urls"
)

type FlashBag struct{}

func (_ FlashBag) Index() template.HTMLAttr { return urls.FlashBagIndex }
func (_ FlashBag) Test() template.HTMLAttr  { return urls.FlashBagTest }
