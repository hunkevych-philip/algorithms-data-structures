package leetcode

func FindFirstOccurrenceOfNeedle(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if haystack == needle {
		return 0
	}

	needleIndexes := len(needle)
	for i, _ := range haystack {
		if len(haystack[i:]) < needleIndexes {
			return -1
		}

		if haystack[i:i+needleIndexes] == needle {
			return i
		}
	}

	return -1
}
