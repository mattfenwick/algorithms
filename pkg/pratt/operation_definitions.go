package pratt

var TestOperators = NewOperators(
	map[string]int{
		"!":       100,
		"+":       100,
		"-":       100,
		"^":       100,
		"*":       100,
		"&":       100,
		"<-":      100,
		"~":       100,
		"pre-low": 0,
		"pre-mid": 50,
	},
	map[string]int{
		";":        100,
		"post-low": 0,
		"post-mid": 50,
	},
	[]*BinOp{
		{"<-0", 0, false},
		{"+", 40, false},
		{"-", 40, false},
		{"bin-mid-left", 50, false},
		{"bin-mid-right", 50, true},
		{"*", 60, false},
		{"/", 60, false},
		{"%", 60, false},
		{"^", 70, false},
		{"**", 70, true},
	},
	[]*TernOp{
		{"?", ":", 0, true},
		{"if", "else", 0, true},
		{"<l", "<lb", 0, false},
		{"<l1", "<l1b", 1, false},
	},
	map[string]string{
		"(": ")",
	},
)
