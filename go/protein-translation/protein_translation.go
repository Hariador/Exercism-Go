package protein

import "errors"

var (
	ErrStop        = errors.New("stop codon")
	ErrInvalidBase = errors.New("invalid base")
)

func FromRNA(rna string) ([]string, error) {
	var codons []string
	if len(rna) < 1 {
		return codons, errors.New("no RNA provided")
	}
	processing := true
	pos := 0
	for processing {
		codon := rna[pos : pos+3]
		protein, err := FromCodon(codon)
		if err == ErrInvalidBase {
			return codons, err
		}
		if err == ErrStop {
			break
		}
		codons = append(codons, protein)
		pos = pos + 3
		if pos+3 > len(rna) {
			break
		}
	}

	return codons, nil
}

func FromCodon(codon string) (string, error) {

	switch codon {
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	}

	return "", ErrInvalidBase

}
