package kindergarten

import (
	"errors"
	"slices"
	"strings"
)

type Garden map[string][]string

func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(diagram) != (len(children)*2*2)+2 {
		return nil, errors.New("invalid diagram")
	}
	uniqueChildren := map[string]struct{}{}
	for _, child := range children {
		uniqueChildren[child] = struct{}{}
	}
	if len(children) != len(uniqueChildren) {
		return nil, errors.New("duplicate name")
	}
	childrenCopy := make([]string, len(children))
	copy(childrenCopy, children)
	slices.Sort(childrenCopy)
	garden := Garden{}
	for _, line := range strings.Split(diagram, "\n") {
		if line == "" {
			continue
		}
		chars := []rune(line)
		for i, child := range childrenCopy {
			first := flowers[chars[i*2]]
			second := flowers[chars[i*2+1]]
			if first == "" || second == "" {
				return nil, errors.New("invalid code")
			}
			garden[child] = append(garden[child], first, second)
		}
	}
	return &garden, nil
}

var flowers = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if plants, ok := (*g)[child]; ok {
		return plants, true
	}
	return nil, false
}
