package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

type match struct {
	line     string
	lineNr   int
	fileName string
}

func Search(pattern string, flags, files []string) []string {
	if slices.Contains(flags, "-x") {
		pattern = "^" + pattern + "$"
	}
	if slices.Contains(flags, "-i") {
		pattern = "(?i)" + pattern
	}
	re := regexp.MustCompile(pattern)
	multiFileMode := len(files) > 1
	matches := []match{}
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			panic("could not read file")
		}
		scanner := bufio.NewScanner(file)
		for i := 1; scanner.Scan(); i++ {
			// -x
			if maybeInvert(re.MatchString(scanner.Text()), slices.Contains(flags, "-v")) {
				matches = append(matches, match{scanner.Text(), i, fileName})
			}
		}
	}
	// handle the output flags here
	if slices.Contains(flags, "-l") {
		result := []string{}
		for _, match := range matches {
			result = append(result, match.fileName)
		}
		return slices.Compact(result)
	}
	result := []string{}
	for _, match := range matches {
		line := maybeAddFileName(match, multiFileMode) +
			maybeAddLineNumber(match, slices.Contains(flags, "-n")) +
			match.line
		result = append(result, line)
	}

	return result
}

func maybeAddLineNumber(m match, add bool) string {
	if add {
		return fmt.Sprint(m.lineNr) + ":"
	}
	return ""
}

func maybeAddFileName(m match, add bool) string {
	if add {
		return m.fileName + ":"
	}
	return ""
}

func maybeInvert(b bool, invert bool) bool {
	if invert {
		return !b
	}
	return b
}
