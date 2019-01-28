package obj

import "testing"

// perintah go test untuk menjalankan unit testing

func TestSayHi(t *testing.T) {
	b := Manusia{Nama: "Budi", Umur: 20}
	s := Manusia{Nama: "Susi", Umur: 20}
	expected := "Budi: Hello Susi"

	actual := b.SayHi(s)

	if expected != actual {
		t.Errorf("expected %v, got %v\n", expected, actual)
	}
}

func TestTambahUmur(t *testing.T) {
	b := Manusia{Nama: "Budi", Umur: 20}
	expected := 21
	actual := b.TambahUmur()

	if expected != actual {
		t.Errorf("expector %v, got %v\n", expected, actual)
	}
}
