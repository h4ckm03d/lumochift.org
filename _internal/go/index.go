package handler

import (
	"html/template"
	"net/http"
)

const tplHtml = `
<!DOCTYPE html>
<html>
<head>
	<meta name="go-import" content="{{ .CanonicalURL }} git https://{{ .Repo }}">
	<meta name="go-source" content="{{ .CanonicalURL }} https://{{ .Repo }} https://{{ .Repo }}/tree/master{/dir} https://{{ .Repo }}/tree/master{/dir}/{file}#L{line}">
	<meta http-equiv="refresh" content="0; url={{ .GodocURL }}">
</head>
<body>
	Nothing to see here. Please <a href="{{ .GodocURL }}">move along</a>.
</body>
</html>
`

var tpl, _ = template.New("Package").Parse(tplHtml)
var repos = []string{"golang-playground", "goworker"}

func has(name string) bool {
	for i := range repos {
		if repos[i] == name {
			return true
		}
	}

	return false
}

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if !has(name) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := struct {
		CanonicalURL string
		Repo         string
		GodocURL     string
	}{
		CanonicalURL: "lumochift.org/go/" + name,
		Repo:         "github.com/h4ckm03d/" + name,
		GodocURL:     "https://godoc.org/lumochift.org/go/" + name,
	}

	tpl.Execute(w, data)
}