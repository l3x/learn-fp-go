package i18n

import (
	"github.com/nicksnyder/go-i18n/i18n"
	"os"
	"text/template"
)

var funcMap = map[string]interface{}{
	"T": i18n.IdentityTfunc,
}

var tmplIllegalBase64Data = template.Must(template.New("").Funcs(funcMap).Parse(`
{{T "illegal_base64_data" .}}
`))
var tmplUnexpectedEndOfJson= template.Must(template.New("").Funcs(funcMap).Parse(`
{{T "unexpected_end_of_json_input"}}
`))
var tmplJsonUnsupportedValue = template.Must(template.New("").Funcs(funcMap).Parse(`
{{T "json_unsupported_value" .}}
`))

func illegalBase64(T i18n.TranslateFunc, bytePos string) {
	tmplIllegalBase64Data.Execute(os.Stdout, map[string]interface{}{
		"BytePos":    bytePos,
	})
}
func unexpectedEndOfJson(T i18n.TranslateFunc) {
	tmplUnexpectedEndOfJson.Execute(os.Stdout, map[string]interface{}{
	})
}
func jsonUnsupportedValue(T i18n.TranslateFunc, bytePos string) {
	tmplJsonUnsupportedValue.Execute(os.Stdout, map[string]interface{}{
		"Val":    bytePos,
	})
}
