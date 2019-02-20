package strand

func ToRNA(dna string) string {
	rna := map[byte]byte{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	result := ""
	if len(dna) == 0 {
		return result
	}

	for i := 0; i < len(dna); i++ {
		result += string(rna[dna[i]])
	}
	return result
}
