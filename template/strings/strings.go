package strings

func StrStr(s, sep string) int {
	i, j := 0, 0
	for {
		if j == len(sep) {
			return i
		} else if i+j == len(s) {
			return -1
		} else if sep[j] == s[i+j] {
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
	i, j := 0, 0
	for {
		if j == len(sep) {
			return i - j
		} else if i == len(s) {
			return -1
		} else if sep[j] == s[i] {
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = prefix[j]
		}
	}
}
