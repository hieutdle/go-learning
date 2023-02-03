package kata

import "strings"

// "GCAT"  =>  "GCAU"

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}
