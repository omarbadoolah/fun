package util

import "testing"

func TestAddAndPop(t *testing.T) {
	var q Queue
	expected := []interface{}{1, 7, 3, 4}
	for _, v := range expected {
		q.Add(v)
	}
	for i := range expected {
		actual := q.Pop().value
		if actual != expected[i] {
			t.Errorf("%d: Actual: %v, Expected: %v", i, actual, expected[i])
		}
	}
}
