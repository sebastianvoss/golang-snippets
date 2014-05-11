package main

import (
	"os"
	"text/template"
)

func main() {

	funcMap := template.FuncMap{
		"slice": func(i int) []int { return make([]int, i) },
	}

	tmpl := `{{$x := .}}{{range slice 10}}<p>{{$x}}</p>{{end}}`
	t, _ := template.New("template").Funcs(funcMap).Parse(tmpl)
	t.Execute(os.Stdout, "42")

}
