package iban

import "testing"

func TestCalculate(t *testing.T) {
	expectIBAN := func(a string, b string, c string) {
		d, err := Calculate(a, b)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if d != c {
			t.Errorf("`" + a + " " + b + "` expected " + c + " , got " + d)
		}
	}

	expectIBAN("123457", "0710", "CZ4907100000000000123457")
}
