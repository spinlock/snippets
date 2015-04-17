package sample

import (
	"fmt"
	"testing"
)

func hanoi(n int, f, m, t byte) {
	if n == 0 {
		return
	}
	hanoi(n-1, f, t, m)
	fmt.Printf("move %d, from %c to %c\n", n, f, t)
	hanoi(n-1, m, f, t)
}

func TestHanoi(t *testing.T) {
	hanoi(5, 'A', 'B', 'C')
}
