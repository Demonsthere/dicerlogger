package logline

// LogLines .
type LogLines struct {
	size  int
	lines []string
}

//NewLogLines .
func NewLogLines() *LogLines {
	return &LogLines{
		lines: populateLines(),
		size:  len(populateLines()),
	}
}

// GetLine return line text corresponding to iterator
func (ll LogLines) GetLine(lineNumber int) string {
	return ll.lines[lineNumber%ll.size]
}

func populateLines() []string {
	return []string{
		"Lorem ipsum dolor sit amet error",
		"consectetur critical adipiscing elit",
		"sed do eiusmod tempor",
		"incididunt fatal ut labore et dolore magna aliqua",
		"Ut enim ad minim veniam",
		"quis error nostrud exercitation",
		"ullamco laboris nisi ut aliquip",
		"ex ea fatal commodo consequat",
		"Duis aute irure dolor in reprehenderit",
		"in critical voluptate velit esse",
		"cillum dolore eu fugiat critical nulla pariatur",
		"Excepteur sint occaecat",
		"cupidatat error non proident",
		"sunt in culpa qui officia deserunt",
		"mollit fatal anim id est laborum",
	}
}
