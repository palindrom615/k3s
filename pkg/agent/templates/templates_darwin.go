//go:build darwin

package templates

import (
	"encoding/json"
	"path/filepath"
	"text/template"
)

var templateFuncs = template.FuncMap{
	"deschemify": func(s string) string {
		return s
	},
	"toJson": func(v any) string {
		output, _ := json.Marshal(v)
		return string(output)
	},
	"filepathjoin": filepath.Join,
}
