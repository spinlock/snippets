package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	next [2]*node
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
		if !(b == '0' || b == '1') {
			break
		}
		if x.tail {
			return false
		}
		i := b - '0'
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

func main() {
	var r = bufio.NewReader(os.Stdin)
	for i := 1; ; i++ {
		t := New()
		var decodable = true
		for {
			s, _ := r.ReadString('\n')
			if s == "" {
				return
			}
			if s[0] == '9' {
				break
			}
			if decodable {
				decodable = t.InsertNoPrefix(s)
			}
		}
		if decodable {
			fmt.Printf("Set %d is immediately decodable\n", i)
		} else {
			fmt.Printf("Set %d is not immediately decodable\n", i)
		}
	}
}
