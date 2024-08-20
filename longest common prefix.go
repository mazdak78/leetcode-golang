package main

import (
	"strings"
)

func main() {
}

func longestCommonPrefix(strs []string) string {
	if len(strs[0]) == 0 {
		return ""
	}

	isGoodPrefix := true
	prefix := ""

outer:
	for i := 0; i < len(strs[0]); i++ {
		prefix = strs[0][:i+1]
		for j := 1; j < len(strs); j++ {
			if strings.HasPrefix(strs[j], prefix) {
				continue
			}

			isGoodPrefix = false
			break outer
		}
	}

	if isGoodPrefix {
		return prefix
	}

	return prefix[0 : len(prefix)-1]
}
