package main

import (
	"os"
	"text/template"
)

func main() {

	funcMap := template.FuncMap{
		"calculate": func(i int) int { return 42 },
	}

	tmpl := `{{$check := eq (calculate 1) 42}}{{if $check}}Correct answer{{end}}{{if not $check}}Wrong answer{{end}}`

	t, _ := template.New("template").Funcs(funcMap).Parse(tmpl)
	t.Execute(os.Stdout, "x")

}
