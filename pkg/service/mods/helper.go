package mods

import (
	"strings"
)

func sortOrder(val string) string {
	if strings.ToLower(val) == "asc" {
		return "ASC"
	}

	return "DESC"
}

// func searchCut(r rune) bool {
// 	return !unicode.IsLetter(r) &&
// 		!unicode.IsNumber(r) &&
// 		r != ' ' &&
// 		r != '\t' &&
// 		r != '_' &&
// 		r != ',' &&
// 		r != '-' &&
// 		r != '.' &&
// 		r != ':' &&
// 		r != '*'
// }
