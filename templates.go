package main

import (
	"text/template"
)

func getTemplate() *template.Template {
	return template.Must(template.New("gen").Parse(tmpl))
}

const tmpl = `// {{.Command}}
// this file was auto-generated using github.com/clipperhouse/gen
// {{.Generated}}

package {{.Package}}

type {{.Plural}} []{{.Pointer}}{{.Singular}}

func ({{.Receiver}} {{.Plural}}) AggregateInt(fn func({{.Pointer}}{{.Singular}}, int) int) (result int) {
	for _, {{.Loop}} := range {{.Receiver}} {
		result = fn({{.Loop}}, result)
	}
	return result
}

func ({{.Receiver}} {{.Plural}}) All(fn func({{.Pointer}}{{.Singular}}) bool) bool {
	for _, {{.Loop}} := range {{.Receiver}} {
		if !fn({{.Loop}}) {
			return false
		}
	}
	return true
}

func ({{.Receiver}} {{.Plural}}) Any(fn func({{.Pointer}}{{.Singular}}) bool) bool {
	for _, {{.Loop}} := range {{.Receiver}} {
		if fn({{.Loop}}) {
			return true
		}
	}
	return false
}

func ({{.Receiver}} {{.Plural}}) Count(fn func({{.Pointer}}{{.Singular}}) bool) (result int) {
	for _, {{.Loop}} := range {{.Receiver}} {
		if fn({{.Loop}}) {
			result++
		}
	}
	return result
}

func ({{.Receiver}} {{.Plural}}) Each(fn func({{.Pointer}}{{.Singular}})) {
	for _, {{.Loop}} := range {{.Receiver}} {
		fn({{.Loop}})
	}
}

func ({{.Receiver}} {{.Plural}}) SumInt(fn func({{.Pointer}}{{.Singular}}) int) (result int) {
	var sum = func({{.Loop}} {{.Pointer}}{{.Singular}}, acc int) int {
		return acc + fn({{.Loop}})
	}
	return {{.Receiver}}.AggregateInt(sum)
}

func ({{.Receiver}} {{.Plural}}) Where(fn func({{.Pointer}}{{.Singular}}) bool) (result {{.Plural}}) {
	for _, {{.Loop}} := range {{.Receiver}} {
		if fn({{.Loop}}) {
			result = append(result, {{.Loop}})
		}
	}
	return result
}
`
