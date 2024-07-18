package taxes

import (
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Table struct {
	Headers []string
	Rows    [][]string
}

func (t *Table) ToFormattedTable() string {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)

	table.SetHeader(t.Headers)
	table.AppendBulk(t.Rows)

	table.Render()
	return tableString.String()
}
