package leetcode

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	l := len(s1)

	var diffA []uint8
	var diffB []uint8
	diffCount := 0
	for i := 0; i < l; i++ {
		if s1[i] != s2[i] {
			diffCount++
			diffA = append(diffA, s1[i])
			diffB = append(diffB, s2[i])
		}

		if diffCount > 2 {
			return false
		}
	}

	if diffCount == 1 {
		return false
	}

	if diffA[0] == diffB[1] && diffA[1] == diffB[0] {
		return true
	}
	return false

}
