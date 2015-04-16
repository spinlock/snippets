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

func Atoi(s string) (out int, ok bool) {
	var idx = 0
	for idx < len(s) && s[idx] == ' ' {
		idx++
	}
	if idx == len(s) {
		return
	}
	var minus = false
	if s[idx] == '-' {
		idx++
		minus = true
	}
	for idx < len(s) && s[idx] == ' ' {
		idx++
	}
	if idx == len(s) {
		return
	}
	var v int
	for idx < len(s) {
		switch b := s[idx]; {
		default:
			return
		case b >= '0' && b <= '9':
			var n int
			if minus {
				n = v*10 - int(b-'0')
				if n > v {
					return
				}
			} else {
				n = v*10 + int(b-'0')
				if n < v {
					return
				}
			}
			v = n
		}
		idx++
	}
	return v, true
}
