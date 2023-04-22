package strand

import "strings"

func ToRNA(dna string) string {
	rna := strings.Builder{}
	mapping := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
	for _, necleotide := range dna {
		rna.WriteRune(mapping[necleotide])
	}
	return rna.String()
}
