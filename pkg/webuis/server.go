package webuis

import (
	"context"
	"html/template"
	"net/http"

	"database/sql"

	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/sirupsen/logrus"

	_ "modernc.org/sqlite"
)

const (
	dbFilePath = "database.sql"
)

var (
	dbHandle *sql.DB
)

func init() {
	logrus.Infof("opening sqlite %s", dbFilePath)
	db := utils.Die(sql.Open("sqlite", dbFilePath))
	logrus.Infof("opened sqlite %s", dbFilePath)
	dbHandle = db
}

type Chord struct {
	Id    string
	Name  string
	Notes string // TODO []string
}

func getChords() []*Chord {
	rows := utils.Die(dbHandle.QueryContext(context.TODO(), "select * from chords"))
	var chords []*Chord
	for rows.Next() {
		var c Chord
		utils.Die0(rows.Scan(&c.Id, &c.Name, &c.Notes))
		chords = append(chords, &c)
	}
	return chords
}

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
        <th>Name</th>
        <th>Category</th>
        <th>Cost</th>
        <th></th>
      </tr>
    </thead>
    <tbody hx-confirm="Are you sure?" hx-target="closest tr" hx-swap="outerHTML swap:1s">
	  {{ range . }}
      <tr>
	    <td>{{ .Name }}</td>
		<td>{{ .Category }}</td>
		<td>{{ .Cost }}</td>
        <td>
            <button class="btn danger" hx-delete="/tool/{{ .Name }}">
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

type Tool struct {
	Name     string
	Category string
	Cost     string
}

func Server() {
	tools := []*Tool{
		{"Hammer", "Banging", "$10"},
		{"Pliers", "Gripping", "$5"},
		{"Saw", "Cutting", "$15"},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("handling %s to %s", r.Method, r.URL.Path)
		rootTemplate.Execute(w, tools)
	})
	http.HandleFunc("/tool/{name}", func(w http.ResponseWriter, r *http.Request) {
		rows := getChords()
		logrus.Infof("found %d rows", len(rows))
		logrus.Infof("handling %s to %s", r.Method, r.URL.Path)
		name := r.PathValue("name")
		for i := 0; i < len(tools); i++ {
			if tools[i].Name == name {
				tools = append(tools[:i], tools[i+1:]...)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	})

	logrus.Info("starting server on :8081")
	logrus.Fatal(http.ListenAndServe(":8081", nil))
}
