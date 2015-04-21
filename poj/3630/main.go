package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	next [10]*node
	tail bool
}

type TrieTree struct {
	root *node
}

func New() *TrieTree {
	return &TrieTree{}
}

func (t *TrieTree) lazyInit() {
	if t.root == nil {
		t.root = &node{}
	}
}

func (t *TrieTree) InsertNoPrefix(s string) bool {
	t.lazyInit()
	x := t.root
	for j := 0; j < len(s); j++ {
		b := s[j]
		if !(b >= '0' && b <= '9') {
			break
		}
		if x.tail {
			return false
		}
		i := int(b - '0')
		if x.next[i] == nil {
			x.next[i] = &node{}
		}
		x = x.next[i]
	}
	if x.tail {
		return false
	}
	for i := 0; i < len(x.next); i++ {
		if x.next[i] != nil {
			return false
		}
	}
	x.tail = true
	return true
}

func process(r *bufio.Reader, n int) {
	var t = New()
	var decodable = true
	for i := 0; i < n; i++ {
		s, _ := r.ReadString('\n')
		if decodable {
			decodable = t.InsertNoPrefix(s)
		}
	}
	if decodable {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var r = bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)
		process(r, n)
	}
}
