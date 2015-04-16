package strings_test

import (
	"math"
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
	out int32
}

var atoi32Tests = []atoiTest{
	{"", 0},
	{"0", 0},
	{"1", 1},
	{"12345", 12345},
	{"012345", 12345},
	{"12345x", 12345},
	{"18446744073709551616", math.MaxInt32},
	{"-0", 0},
	{"-1", -1},
	{"12345", 12345},
	{"-12345", -12345},
	{"012345", 12345},
	{"-012345", -12345},
	{"-9223372036854775807", math.MinInt32},
}

func TestAtoi32(t *testing.T) {
	for _, t := range atoi32Tests {
		out := Atoi32(t.in)
		assert.Must(t.out == out)
	}
}
