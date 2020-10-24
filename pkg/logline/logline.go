package logline

import (
	"bufio"
	"os"
)

// LogLines .
type LogLines struct {
	size  int
	lines []string
}

//NewLogLines use internal lines
func NewLogLines() *LogLines {

	return &LogLines{
		lines: populateLines(),
		size:  len(populateLines()),
	}
}

//NewLogLinesFromFile use supplied file
func NewLogLinesFromFile(file string) (*LogLines, error) {
	lines, err := readLines(file)
	size := len(lines)
	return &LogLines{
		lines: lines,
		size:  size,
	}, err
}

// GetLine return line text corresponding to iterator
func (ll LogLines) GetLine(lineNumber int) string {
	return ll.lines[lineNumber%ll.size]
}

func populateLines() []string {
	return []string{
		"info Lorem ipsum dolor sit amet",
		"consectetur adipiscing error elit",
		"critical sed do eiusmod tempor",
		"incididunt ut fatal labore et dolore magna aliqua",
		"Ut enim ad minim veniam error ",
		"quis nostrud exercitation debug ",
		"ullamco info laboris nisi ut aliquip",
		"ex ea fatal commodo consequat",
		"Duis aute irure dolor in info reprehenderit",
		"critical in voluptate velit esse",
		"cillum critical dolore eu fugiat nulla pariatur",
		"error Excepteur sint occaecat",
		"cupidatat fatal non proident",
		"error sunt in culpa qui officia deserunt",
		"mollit anim id est critical laborum",
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
