package sample

import (
	"fmt"
	"testing"
)

func convert_base(n int, base int) string {
	var t = make([]byte, 16)
	for i := 0; i < 16; i++ {
		if i < 10 {
			t[i] = '0' + byte(i)
		} else {
			t[i] = 'a' + byte(i-10)
		}
	}
	var s []byte
	for n != 0 {
		m := n % base
		s = append(s, t[m])
		n = n / base
	}
	if len(s) == 0 {
		s = append(s, t[0])
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func TestConvert(t *testing.T) {
	const n = 123456789
	for i := 2; i <= 16; i++ {
		fmt.Printf("%2d: %s\n", i, convert_base(n, i))
	}
}
