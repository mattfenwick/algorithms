package webuis

import (
	"context"
	"encoding/json"

	"database/sql"

	"github.com/google/uuid"
	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/sirupsen/logrus"

	_ "modernc.org/sqlite"
)

const (
	dbFilePath = "database.sql"
)

const schema = `
create table chords (
	id text PRIMARY KEY,
	name text unique,
	notes text
)
`

var (
	dbHandle *sql.DB
)

func getTables() map[string]int {
	result := utils.Die(dbHandle.QueryContext(context.TODO(), `select name from sqlite_master where type = 'table'`))
	tables := map[string]int{}
	for result.Next() {
		var next string
		utils.Die0(result.Scan(&next))
		tables[next]++
	}
	return tables
}

// func isTableCreated(tableName string) bool {
// 	row := dbHandle.QueryRowContext(context.TODO(), `select name from sqlite_master WHERE type='table' and name=?`, tableName)
// 	// utils.Die0(row.Scan())
// 	return row.Scan() == nil
// }

func init() {
	logrus.Infof("opening sqlite %s", dbFilePath)
	db := utils.Die(sql.Open("sqlite", dbFilePath))
	logrus.Infof("opened sqlite %s", dbFilePath)
	dbHandle = db

	tables := getTables()
	if tables["chords"] > 0 {
		logrus.Infof("table 'chords' already exists, skipping 'create table' and seeding")
	} else {
		schemaResult := utils.Die(db.ExecContext(context.TODO(), schema))
		logrus.Infof("schema result: %+v", schemaResult)

		logrus.Infof("created %+v", createChord("a", []string{"b", "c", "d"}))
		logrus.Infof("created %+v", createChord("b", []string{"b", "b", "b"}))
	}
	logrus.Infof("tables? %d, %d", tables["chords"], tables["boards"])
}

type Chord struct {
	Id    string
	Name  string
	Notes []string
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

func createChord(name string, notes []string) *Chord {
	result := utils.Die(dbHandle.ExecContext(context.TODO(),
		`insert into chords (id, name, notes) values (?, ?, ?)`,
		uuid.New().String(),
		name,
		string(utils.Die(json.Marshal(notes)))))
	// TODO this part isn't necessary but it's a nice example of how to use sqlite
	lastInsertId := utils.Die(result.LastInsertId())
	row := dbHandle.QueryRowContext(context.TODO(), `select id, name, notes from chords where rowid = ?`, lastInsertId)
	var c Chord
	var notesFromDb string
	utils.Die0(row.Scan(&c.Id, &c.Name, &notesFromDb))
	utils.Die0(json.Unmarshal([]byte(notesFromDb), &c.Notes))
	return &c
}
