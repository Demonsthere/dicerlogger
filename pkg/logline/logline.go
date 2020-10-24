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
		"|info| Lorem ipsum dolor sit amet",
		"|error| consectetur adipiscing elit",
		"|critical| sed do eiusmod tempor",
		"|fatal| incididunt ut labore et dolore magna aliqua",
		"|error| Ut enim ad minim veniam",
		"|debug| quis nostrud exercitation",
		"|info| ullamco laboris nisi ut aliquip",
		"|fatal| ex ea commodo consequat",
		"|info| Duis aute irure dolor in reprehenderit",
		"|critical| in voluptate velit esse",
		"|critical| cillum dolore eu fugiat nulla pariatur",
		"|error| Excepteur sint occaecat",
		"|fatal| cupidatat non proident",
		"|error| sunt in culpa qui officia deserunt",
		"|critical| mollit anim id est laborum",
	}
}
