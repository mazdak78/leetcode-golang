package main

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	needleBytes := needle[0:]
	for i := 0; i < len(haystack); i++ {
		if haystack[i:i+len(needle)] == needleBytes {
			return i
		}

		if len(haystack[i:]) <= len(needleBytes) {
			break
		}
	}

	return -1
}
