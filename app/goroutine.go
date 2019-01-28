package main

import (
	"fmt"
	"time"
	"sync"
)

type hasil struct {
	nama string
	value int
}

// fun adalah func(int) int -> function yang terima 1 parameter int dan return int
func cetak(nama string, in chan int, out chan hasil, fun func(int) int) {
	for item := range in {
		// kirim hasil ke channel out
		out <- hasil{nama: nama, value: fun(item)}
		time.Sleep(time.Second)
	}
}

// in chan int -> channel yang menerima dan mengirimkan int
func manager(a []int, in chan int) {
	for _, i := range a {
		in <- i
	}
}

type worker struct {
		nama string
		fn func(int) int
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// unbuffered channel
	in := make(chan int)
	out := make(chan hasil)
	workers := []worker{
		// fn diisikan anonymous function (lambda)
		{nama: "A", fn: func(n int) int { return n * 2 }},
		{nama: "B", fn: func(n int) int { return n + 3 }},
		{nama: "C", fn: func(n int) int { return n - 1 }},
	}

	// untuk menunggu seluruh worker selesai bekerja
	wg := sync.WaitGroup{}
	wg.Add(len(workers))

	// menunggu print out
	wg1 := sync.WaitGroup{}
	wg1.Add(1)

	for _, w := range workers {
		go func(w worker) {
			// worker selesai bekerja
			defer wg.Done()
			cetak(w.nama, in, out, w.fn)
		}(w)
	}

	go func() {
		manager(a, in)
		// selesai, tutup channel
		close(in)
	}()

	// terima hasil dari channel out
	go func() {
		defer wg1.Done()
		total := 0
		for h := range out {
			fmt.Printf("%s: %d\n", h.nama, h.value)
			total += h.value
		}
		fmt.Println("Total: ", total)
	}()
	wg.Wait()
	close(out)
	wg1.Wait()
}
