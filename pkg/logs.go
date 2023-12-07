package pkg

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Query struct {
	Id    int
	Raw   string
	Terms []string
}

func (q *Query) Match(l *Log) bool {
	for _, queryTerm := range q.Terms {
		found := false
		for _, logTerm := range l.Terms {
			if queryTerm == logTerm {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

type Log struct {
	Raw   string
	Terms []string
}

// type Match

type Parser struct {
	NextQueryId int
}

func (p *Parser) ParseInput(in string) (*Query, *Log) {
	terms := strings.Split(strings.ToLower(in[3:]), " ")
	switch in[0] {
	case 'Q':
		queryId := p.NextQueryId
		p.NextQueryId++
		return &Query{Id: queryId, Raw: in[3:], Terms: terms}, nil
	case 'L':
		return nil, &Log{Raw: in[3:], Terms: terms}
	default:
		panic(errors.Errorf("invalid input"))
	}
}

func LogsMain() {
	logs := []string{
		"Q: abc DEF",
		"Q: Abc ghi",
		"Q: aBc",
		"L: mno pqr",
		"L: ABC",
		"L: abc def GHi mno pqr",
		"L: abc def mno",
		"Q: defg",
		"L: abcd defgh",
		"Q: mno abc",
		"L: mno def abc",
	}

	parser := &Parser{NextQueryId: 1}

	queries := []*Query{}
	index := map[string][]*Query{}
	fmt.Printf("%+v\n%+v\n", index, map[*Query]int{})

	for _, log := range logs {
		q, l := parser.ParseInput(log)
		//fmt.Printf("%+v, %+v\n", q, l)
		if q != nil {
			queries = append(queries, q)
			for _, term := range q.Terms {
				if _, ok := index[term]; !ok {
					index[term] = []*Query{}
				}
				index[term] = append(index[term], q)
			}
			fmt.Printf("ACK: '%s' (%+v); Id:%d\n", q.Raw, q.Terms, q.Id)
		} else if l != nil {
			// using list of queries
			queryMatches := []*Query{}
			for _, query := range queries {
				if query.Match(l) {
					queryMatches = append(queryMatches, query)
				}
			}
			if len(queryMatches) > 0 {
				var ids []int
				for _, query := range queryMatches {
					ids = append(ids, query.Id)
				}
				fmt.Printf("FOUND: (%s, %+v); [%+v]\n", l.Raw, l.Terms, ids)
			} else {
				fmt.Printf("NOT found: (%s, %+v)\n", l.Raw, l.Terms)
			}

			// using index
			indexCounts := map[*Query]int{}
			for _, logTerm := range l.Terms {
				if queries, ok := index[logTerm]; ok {
					for _, query := range queries {
						indexCounts[query]++
					}
				}
			}
			foundType2 := []int{}
			for query, count := range indexCounts {
				if len(query.Terms) == count {
					foundType2 = append(foundType2, query.Id)
				}
			}
			if len(foundType2) > 0 {
				fmt.Printf("FOUND (2): (%s, %+v); [%+v]\n", l.Raw, l.Terms, foundType2)
			} else {
				fmt.Printf("NOT found (2): (%s, %+v)\n", l.Raw, l.Terms)
			}

			fmt.Println()
		} else {
			panic("unexpected")
		}
	}
}
