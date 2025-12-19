package ui

import "testing"

func TestHexToANSI(t *testing.T) {
	var cases = []struct {
		c string
		e string
	}{
		{
			"#f0ec00ff",
			"2;240;236;0m",
		},
	}
	for _, c := range cases {
		if result := HexToANSI(c.c); result != c.e {
			t.Errorf("HexToANSI(%s) = %s, expect: %s", c.c, result, c.e)
		}
	}
}
