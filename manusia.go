package obj

import "fmt"

// Manusia
// seluruh type yang dimulai dengan huruf besar diexport supaya dapat
// diimport dan digunakan pada package lain
type Manusia struct {
	Nama string
	Umur int
}

// method SayHi dari struct manusia
func (m Manusia) SayHi(o Manusia) string {
	return fmt.Sprintf("%s: Hello %s", m.Nama, o.Nama)
}

func (m Manusia) Print() {
	fmt.Printf("Hello, saya %s berumur %d tahun\n", m.Nama, m.Umur)
}

func (m *Manusia) TambahUmur() int {
	m.Umur++
	return m.Umur
}
