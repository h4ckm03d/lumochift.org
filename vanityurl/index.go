package handler

import (
	"html/template"
	"log"
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

func has(name string) bool {
	resp, err := http.Get("https://api.github.com/repos/h4ckm03d/" + name)
	if err != nil {
		log.Println(err)
		return false
	}

	if resp.StatusCode != 200 {
		return false
	}

	return true
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
		CanonicalURL: "go.lumochift.org/" + name,
		Repo:         "github.com/h4ckm03d/" + name,
		GodocURL:     "https://godoc.org/go.lumochift.org/" + name,
	}

	_ = tpl.Execute(w, data)
}
