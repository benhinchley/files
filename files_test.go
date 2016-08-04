package files

import "testing"

type test struct {
	input    interface{}
	expected interface{}
}

func TestExists(t *testing.T) {
	var tests = []test{
		{"/this/is/not/a/path", false},
		{"files.go", true},
	}

	for _, test := range tests {
		if v := Exists(test.input.(string)); v != test.expected.(bool) {
			t.Error(
				"For", test.input.(string),
				"expected", test.expected.(bool),
				"got", v,
			)
		}
	}
}

func TestStripRoot(t *testing.T) {
	var tests = []test{
		{[]string{"this/is/a/path", "this/is/a/path/to/a/file"}, "to/a/file"},
	}

	for _, test := range tests {
		i := test.input.([]string)
		e := test.expected.(string)
		if v := StripRoot(i[0], i[1]); v != e {
			t.Error(
				"For", i,
				"expected", e,
				"got", v,
			)
		}
	}
}
