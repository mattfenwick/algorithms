package utils

import (
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

type Table struct {
	Headers []string
	Rows    [][]string
}

func NewTable(headers []string, rows ...[]string) *Table {
	t := &Table{
		Headers: headers,
		Rows:    nil,
	}
	for _, r := range rows {
		t.AddRow(r)
	}
	return t
}

func (t *Table) AddRow(row []string) {
	if len(row) != len(t.Headers) {
		panic(errors.Errorf("len of row doesn't match length of headers (%d vs %d)", len(row), len(t.Headers)))
	}
	t.Rows = append(t.Rows, row)
}

func (t *Table) ToFormattedTable() string {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	// table.SetAutoMergeCells(true)

	table.SetHeader(t.Headers)
	table.AppendBulk(t.Rows)

	table.Render()
	return tableString.String()
}

/*
func (t *Table) ToFormattedTable(format *func(t *tablewriter.Table)) string {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	if format != nil {
		(*format)(table)
	}
	// table.SetAutoMergeCells(true)

	table.SetHeader(t.Headers)
	table.AppendBulk(t.Rows)

	table.Render()
	return tableString.String()
}
*/
