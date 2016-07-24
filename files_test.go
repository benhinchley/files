package files

import "testing"

type test struct {
	input    interface{}
	expected interface{}
}

var existsTests = []test{
	{"/this/is/not/a/path", false},
	{"files.go", true},
}

func TestExists(t *testing.T) {
	for _, test := range existsTests {
		if v := Exists(test.input.(string)); v != test.expected.(bool) {
			t.Error(
				"For", test.input.(string),
				"expected", test.expected.(bool),
				"got", v,
			)
		}
	}
}

var stripRootTests = []test{
	{[]string{"this/is/a/path", "this/is/a/path/to/a/file"}, "to/a/file"},
}

func TestStripRoot(t *testing.T) {
	for _, test := range stripRootTests {
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
