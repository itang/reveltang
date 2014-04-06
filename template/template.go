package template

import (
	"fmt"
	"html"
	"html/template"

	"github.com/revel/revel"
)

var ExtTemplateFuncs = map[string]interface{}{
	"radiox": func(f *revel.Field, val string, rval string) template.HTML {
		checked := ""
		if f.Flash() == val {
			checked = " checked"
		} else if rval == val {
			checked = " checked"
		}
		return template.HTML(fmt.Sprintf(`<input type="radio" name="%s" value="%s"%s>`,
			html.EscapeString(f.Name), html.EscapeString(val), checked))
	},
	"checkboxx": func(f *revel.Field, val string, rval string) template.HTML {
		checked := ""
		if f.Flash() == val {
			checked = " checked"
		} else if rval == val {
			checked = " checked"
		}
		return template.HTML(fmt.Sprintf(`<input type="checkbox" name="%s" value="%s"%s>`,
			html.EscapeString(f.Name), html.EscapeString(val), checked))
	},
	"optionx": func(f *revel.Field, val, label string, rval string) template.HTML {
		selected := ""
		if f.Flash() == val {
			selected = " selected"
		} else if rval == val {
			selected = " selected"
		}
		return template.HTML(fmt.Sprintf(`<option value="%s"%s>%s</option>`,
			html.EscapeString(val), selected, html.EscapeString(label)))
	},
	"flash": func(renderArgs map[string]interface{}, name string) string {
		v, _ := renderArgs["flash"].(map[string]string)[name]
		return v
	},
	"renderArgs": func(key string, renderArgs map[string]interface{}) interface{} {
		v, ok := renderArgs[key]
		if !ok {
			return ""
		}
		return v
	},
	"notEq": func(a, b interface{}) bool {
		return !revel.Equal(a, b)
	},
}
