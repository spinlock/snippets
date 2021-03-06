package strings

import "math"

func StrStr(s, sep string) int {
	var i, j int
	for {
		if j == len(sep) {
			return i
		} else if i+j == len(s) {
			return -1
		} else if s[i+j] == sep[j] {
			j++
		} else {
			i++
			j = 0
		}
	}
}

func KmpStr(s, sep string) int {
	prefix := make([]int, len(sep))
	for j := 2; j < len(sep); j++ {
		l := prefix[j-1]
		for l != 0 && sep[l] != sep[j-1] {
			l = prefix[l]
		}
		if sep[l] == sep[j-1] {
			prefix[j] = l + 1
		}
	}
	var i, j int
	for {
		if j == len(sep) {
			return i - j
		} else if i == len(s) {
			return -1
		} else if s[i] == sep[j] {
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = prefix[j]
		}
	}
}

func Atoi32(s string) int32 {
	var idx int
	for idx < len(s) && s[idx] == ' ' {
		idx++
	}
	if idx == len(s) {
		return 0
	}

	var minus = false
	switch s[idx] {
	case '-':
		idx++
		minus = true
	case '+':
		idx++
	}

	var out int32
	for ; idx < len(s); idx++ {
		if b := s[idx]; b >= '0' && b <= '9' {
			tmp := 10 * out
			if minus {
				tmp -= int32(b - '0')
			} else {
				tmp += int32(b - '0')
			}
			if tmp/10 == out {
				out = tmp
			} else if minus {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		} else {
			return out
		}
	}
	return out
}
