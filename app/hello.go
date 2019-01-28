package main

import (
	"fmt"
	"obj"
)

type kata string

type Cetak interface {
	Print()
}

func (k kata) Print() {
	fmt.Println(k)
}

// pointer
func tambah(n *int, i int) {
	// deference
	*n = *n + i
}

func increment(n int) int {
	return n + 1
}

// menerima function as parameter
func operasi(n *int, m int, fn func(*int, int)) {
	fn(n, m)
}

// define funckali as type func(int) int
type funckali func(int) int

// closure
func kali(n int) funckali {
	return func(m int) int {
		return n * m
	}
}

func main() {
	// full version, deklarasi + inisialisasi
	var nama kata = "Budi"
	fmt.Printf("Hello, %s\n", nama)

	// type inference
	budi := obj.Manusia{Nama: "Budi", Umur: 20}
	susi := obj.Manusia{Nama: "Susi", Umur: 22}
	budi.SayHi(susi)

	a := []Cetak{kata("testing"), budi}
	for _, i := range a {
		i.Print()
	}

	budi.Print()
	budi.TambahUmur()
	budi.Print()

	n := 10
	fmt.Printf("n: %d\n", n);
	// reference
	tambah(&n, 1)
	fmt.Printf("n: %d\n", n);

	c := 10 // c int
	p := &c // p *int
	fmt.Printf("%v %v %v\n", c, p, *p)
	*p = 100
	fmt.Printf("%v %v %v\n", c, p, *p)

	// add tipe nya adalah function yang terima 2 parameter *int dan int
	var add func(*int, int) = tambah
	add(&n, 1)
	fmt.Println(n)
	// anonymous function/lambda -> function literal
	add = func(n *int, m int) {
		*n = *n * m
	}
	add(&n, 2)
	fmt.Println(n)
	operasi(&n, 1, tambah)
	fmt.Println(n)
	operasi(&n, 2, add)
	fmt.Println(n)
	operasi(&n, 2, add)

	var kali10 funckali = kali(10)
	fmt.Println(kali10(3))
}
