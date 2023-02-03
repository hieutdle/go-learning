package kata

import "strings"

// Sam Harris => S.H
// patrick feeney => P.F

func AbbrevName(name string) string {
	var s []string
	for _, part := range strings.Fields(name) {
		s = append(s, strings.ToUpper(part[:1]))
	}
	return strings.Join(s, ".")
}
