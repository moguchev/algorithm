package decode

import "strconv"

// NumDecodings :
// A message containing letters from A-Z is being encoded to numbers using the following mapping:
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// Given a non-empty string containing only digits, determine the total number of ways to decode it.
// Input: s = "226"
// Output: 3
// Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
func NumDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	runes := make([]rune, 0, len(s))
	for _, ch := range s {
		runes = append(runes, ch)
	}

	m := make(map[int]int)
	return recursiveWithMemo(0, runes, m)
}

func recursiveWithMemo(index int, s []rune, m map[int]int) int {
	if index == len(s) {
		return 1
	}

	if s[index] == '0' {
		return 0
	}

	if index == len(s)-1 {
		return 1
	}

	if _, ok := m[index]; ok {
		return m[index]
	}

	res := recursiveWithMemo(index+1, s, m) // 0-1
	if n, _ := strconv.Atoi(string(s[index : index+2])); n <= 26 {
		res += recursiveWithMemo(index+2, s, m) // 0-1
	}
	m[index] = res // 0-1-2
	return res
}
