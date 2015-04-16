package strings_test

import (
	"testing"

	"github.com/spinlock/go-libs/assert"
	. "github.com/spinlock/snippets/template/strings"
)

type IndexTest struct {
	s   string
	sep string
	out int
}

var indexTests = []IndexTest{
	{"", "", 0},
	{"", "a", -1},
	{"", "foo", -1},
	{"fo", "foo", -1},
	{"foo", "foo", 0},
	{"oofofoofooo", "f", 2},
	{"oofofoofooo", "foo", 4},
	{"barfoobarfoo", "foo", 3},
	{"foo", "", 0},
	{"foo", "o", 1},
	{"abcABCabc", "A", 3},
	{"", "a", -1},
	{"x", "a", -1},
	{"x", "x", 0},
	{"abc", "a", 0},
	{"abc", "b", 1},
	{"abc", "c", 2},
	{"abc", "x", -1},
}

func TestStrStr(t *testing.T) {
	for _, t := range indexTests {
		n := StrStr(t.s, t.sep)
		assert.Must(n == t.out)
	}
}

func TestKmpStr(t *testing.T) {
	for _, t := range indexTests {
		n := KmpStr(t.s, t.sep)
		assert.Must(n == t.out)
	}
}

type atoiTest struct {
	in  string
	out int
	ok  bool
}

var atoiTests = []atoiTest{
	{"", 0, false},
	{"0", 0, true},
	{"1", 1, true},
	{"12345", 12345, true},
	{"012345", 12345, true},
	{"12345x", 0, false},
	{"98765432100", 98765432100, true},
	{"18446744073709551616", 0, false},
	{"18446744073709551620", 0, false},
	{"", 0, false},
	{"0", 0, true},
	{"-0", 0, true},
	{"1", 1, true},
	{"-1", -1, true},
	{"12345", 12345, true},
	{"-12345", -12345, true},
	{"012345", 12345, true},
	{"-012345", -12345, true},
	{"98765432100", 98765432100, true},
	{"-98765432100", -98765432100, true},
	{"9223372036854775807", 1<<63 - 1, true},
	{"-9223372036854775807", -(1<<63 - 1), true},
	{"9223372036854775808", 0, false},
	{"-9223372036854775808", -1 << 63, true},
	{"9223372036854775809", 0, false},
	{"-9223372036854775809", 0, false},
}

func TestAtoi(t *testing.T) {
	for _, t := range atoiTests {
		out, ok := Atoi(t.in)
		if t.ok {
			assert.Must(ok && t.out == out)
		} else {
			assert.Must(!ok)
		}
	}
}
