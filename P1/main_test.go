package main

import "testing"

func TestLuasSegitiga(t *testing.T) {
	s := Segitiga{
		Alas:   10,
		Tinggi: 8,
		Sisi:   10,
	}

	hasil := s.Luas()
	expected := 40.0

	if hasil != expected {
		t.Errorf("Luas salah. Hasil = %.2f, seharusnya = %.2f", hasil, expected)
	}
}

func TestKelilingSegitiga(t *testing.T) {
	s := Segitiga{
		Alas:   10,
		Tinggi: 8,
		Sisi:   10,
	}

	hasil := s.Keliling()
	expected := 30.0

	if hasil != expected {
		t.Errorf("Keliling salah. Hasil = %.2f, seharusnya = %.2f", hasil, expected)
	}
}