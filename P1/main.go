package main

type Segitiga struct {
	Alas   float64
	Tinggi float64
	Sisi   float64
}

func (s Segitiga) Luas() float64 {
	return 0.5 * s.Alas * s.Tinggi
}

func (s Segitiga) Keliling() float64 {
	return s.Sisi * 3
}

func main() {
	// kosong tidak apa-apa
}