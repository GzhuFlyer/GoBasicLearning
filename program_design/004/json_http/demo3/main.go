package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const temp1 = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(temp1))
	var data struct {
		A string        //不受信任的纯文本
		B template.HTML //受信任的 HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>world!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
