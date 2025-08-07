package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type option struct {
	lineNumber, fileName, insensitive, invert, fullMatch bool
}

func Search(pattern string, flags, files []string) []string {
	opt := parseOptions(flags)
	if opt.fullMatch {
		pattern = "^" + pattern + "$"
	}
	if opt.insensitive {
		pattern = "(?i)" + pattern
	}
	re := regexp.MustCompile(pattern)
	multiFile := len(files) > 1
	matches := []string{}
	for _, fileName := range files {
		file, err := os.Open(fileName)
		defer file.Close()
		if err != nil {
			panic("could not read file")
		}
		scanner := bufio.NewScanner(file)
		for i := 1; scanner.Scan(); i++ {
			match := re.MatchString(scanner.Text())
			if opt.invert {
				match = !match
			}
			if match {
				if opt.fileName {
					matches = append(matches, fileName)
					break
				}
				var line string
				if multiFile {
					line = fileName + ":"
				}
				if opt.lineNumber {
					line = line + fmt.Sprint(i) + ":"
				}
				line = line + scanner.Text()
				matches = append(matches, line)
			}
		}
	}
	return matches
}

func parseOptions(flags []string) option {
	option := option{}
	for _, flag := range flags {
		switch flag {
		case "-n":
			option.lineNumber = true
		case "-l":
			option.fileName = true
		case "-i":
			option.insensitive = true
		case "-v":
			option.invert = true
		case "-x":
			option.fullMatch = true
		}
	}
	return option
}
