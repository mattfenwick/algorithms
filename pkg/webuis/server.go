package webuis

import (
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"

	_ "modernc.org/sqlite"
)

const (
	rootTemplateText = `
<!DOCTYPE html>
	<head>
		<script src="https://unpkg.com/htmx.org@2.0.4"></script>
	</head>

<body>

<table class="table delete-row-example">
    <thead>
      <tr>
        <th>Id</th>
        <th>Name</th>
        <th>Notes</th>
        <th></th>
      </tr>
    </thead>
    <tbody hx-confirm="Are you sure?" hx-target="closest tr" hx-swap="outerHTML swap:1s">
	  {{ range . }}
      <tr>
	    <td>{{ .Id }}</td>
		<td>{{ .Name }}</td>
		<td>{{ .Notes }}</td>
        <td>
            <button class="btn danger" hx-delete="/chord/{{ .Id }}">
              Delete
            </button>
        </td>
      </tr>
	  {{ end }}
    </tbody>
  </table>

</body>
</html>
`
)

var (
	rootTemplate = template.Must(template.New("root").Parse(rootTemplateText))
)

func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("/ handling %s to %s", r.Method, r.URL.Path)
		if r.URL.Path != "/" {
			logrus.Errorf("not found -- %s to %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		chords := getChords()
		rootTemplate.Execute(w, chords)
	})
	http.HandleFunc("/chord/{id}", func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("/chord/ handling %s to %s", r.Method, r.URL.Path)
		if deleteChord(r.PathValue("id")) {
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	logrus.Info("starting server on :8081")
	logrus.Fatal(http.ListenAndServe(":8081", nil))
}
