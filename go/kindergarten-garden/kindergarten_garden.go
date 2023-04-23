package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	row1, row2 string
	children   []string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	rows := strings.Split(diagram, "\n")
	if len(rows) < 3 {
		return nil, errors.New("invalid diagram")
	}
	if len(rows[1]) != len(rows[2]) {
		return nil, errors.New("invalid diagram")
	}

	if len(rows[1])%2 != 0 {
		return nil, errors.New("invalid diagram")
	}
	if !validPlants(rows[1]) {
		return nil, errors.New("invalid plant")
	}
	if !validPlants(rows[2]) {
		return nil, errors.New("invalid plant")
	}
	if !checkNames(children) {
		return nil, errors.New("duplicate names")
	}

	var kids []string
	for _, kid := range children {
		kids = append(kids, kid)
	}
	sort.Strings(kids)

	garden := Garden{row1: rows[1], row2: rows[2], children: kids}

	return &garden, nil

}

func (g *Garden) Plants(child string) ([]string, bool) {

	plantNames := map[string]string{"V": "violets", "R": "radishes", "C": "clover", "G": "grass"}
	var offset int
	found := false

	for k, v := range g.children {
		if v == child {
			offset = k
			found = true
		}
	}
	if !found {
		return []string{}, false
	}
	output := []string{"", "", "", ""}

	output[0] = plantNames[string(g.row1[0+offset*2])]
	output[1] = plantNames[string(g.row1[1+offset*2])]
	output[2] = plantNames[string(g.row2[0+offset*2])]
	output[3] = plantNames[string(g.row2[1+offset*2])]
	return output, true
}

func validPlants(plants string) bool {
	for _, v := range plants {
		if v == 'V' || v == 'R' || v == 'G' || v == 'C' {

		} else {
			return false
		}
	}
	return true
}

func checkNames(children []string) bool {
	names := make(map[string]bool)
	for _, name := range children {
		_, exists := names[name]
		if exists {
			return false
		} else {
			names[name] = true
		}
	}
	return true
}
