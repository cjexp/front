package urls

import "html/template"

type FlashBagUrlScope struct{}

func (_ FlashBagUrlScope) Index() template.HTMLAttr { return FlashBagIndex }
func (_ FlashBagUrlScope) Test() template.HTMLAttr  { return FlashBagTest }
