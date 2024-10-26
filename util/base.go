package util

import "strings"

func IsArrayContains(array []string, wanted string) bool {
	for _, v := range array {
		if v[len(v)-1] == '*' {
			if strings.Contains(wanted, v[0:len(v)-1]) {
				return true
			}
		} else {
			if v == wanted {
				return true
			}
		}
	}
	return false
}
