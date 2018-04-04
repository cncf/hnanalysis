package hnanalysis

import (
	"testing"

	lib "hnanalysis"
	testlib "hnanalysis/test"
)

func TestDummy(t *testing.T) {

	// Test cases
	var testCases = []struct {
		name     string
		expected string
	}{
		{
			"Dummy test",
			"bool:true",
		},
	}

	// Execute test cases
	for index, test := range testCases {
		res := lib.Dummy()
		got := testlib.Desc(res)
		if got != test.expected {
			t.Errorf(
				"Test case number %d \"%s\"\nExpected:\n%+v\nGot:\n%+v\n",
				index+1, test.name, test.expected, got,
			)
		}
	}
}
