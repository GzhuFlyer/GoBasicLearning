package myTemplate

import (
	"html/template"
	"time"
)

const temp1 = `{{.TotalCount}} issues:

{{range .Items}}---------------------------------------
	Number:{{.Number}}
	User:  {{.User.Login}}
	Title: {{.Title | printf "%.64s"}}
	Age:	{{.CreatedAt | daysAgo}} days
	{{end}}`

var Report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(temp1))

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
