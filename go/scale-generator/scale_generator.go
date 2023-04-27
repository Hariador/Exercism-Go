package scale

import (
	"strings"
	"unicode"
)

func Scale(tonic, interval string) []string {
	sharps := scale{notes: []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}}
	flates := scale{notes: []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}}

	var workingSet scale
	switch tonic {
	case "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "df#", "C", "a":
		{
			workingSet = sharps
		}
	default:
		{
			workingSet = flates
		}
	}
	return workingSet.getDiatonic(tonic, interval)

}

type scale struct {
	notes []string
}

func (s *scale) getOffset(tonic string) int {
	if len(tonic) > 1 {
		tonic = string(unicode.ToUpper(rune(tonic[0]))) + string(tonic[1])
	} else {
		tonic = strings.ToUpper(tonic)
	}
	for k, v := range s.notes {
		if v == tonic {
			return k
		}
	}

	return 0
}

func (s *scale) get(index int) string {
	index = index % s.Len()
	return s.notes[index]
}

func (s *scale) getDiatonic(tonic string, interval string) []string {
	index := s.getOffset(tonic)
	var output []string

	if interval == "" {
		for i := 0; i < s.Len(); i++ {
			output = append(output, s.get(i+index))
		}
		return output
	}
	output = append(output, s.get(index))
	for _, v := range interval {
		switch v {
		case 'm':
			{
				index += 1
				output = append(output, s.get(index))
			}
		case 'M':
			{
				index += 2
				output = append(output, s.get(index))
			}
		case 'A':
			{
				index += 3
				output = append(output, s.get(index))
			}
		}

	}
	return output
}

func (s *scale) Len() int { return len(s.notes) }
