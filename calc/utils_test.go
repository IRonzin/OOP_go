package calc

import (
	"testing"
)

type TestSet struct {
	state, in, expected int
}

var AddSet = []TestSet{
	{state: 5, in: 5, expected: 10},
	{state: 6, in: 5, expected: 11},
	{state: -7, in: 5, expected: -2},
	{state: 9, in: 5, expected: 14},
	{state: 0, in: 5, expected: 5}}

func TestChanger_Add(t *testing.T) {

	for i, test := range AddSet {
		ch := NewChanger(test.state)

		ch.Add(test.in)

		if ch.state != test.expected {
			t.Errorf("Test case %v: result %v , expected %v", i, ch.state, test.expected)
		}
	}
}

func BenchmarkChanger_Add(b *testing.B) {

	ch := NewChanger(1)
	for i := 0; i < b.N; i++ {
		ch.Add(12)
	}
}
