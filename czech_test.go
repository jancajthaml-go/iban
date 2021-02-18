package iban

import "testing"

func BenchmarkCalculateCzech(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CalculateCzech("123457", "0710")
	}
}

func BenchmarkCalculateCzechParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			CalculateCzech("123457", "0710")
		}
	})
}

func TestCalculateCzech(t *testing.T) {
	expectIBAN := func(a string, b string, c string) {
		d, err := CalculateCzech(a, b)
		if err != nil {
			t.Errorf("%+v", err)
		}
		if d != c {
			t.Errorf("`" + a + " " + b + "` expected " + c + " , got " + d)
		}
	}

	expectIBAN("123457", "0710", "CZ4907100000000000123457")
}
