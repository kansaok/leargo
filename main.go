package main

import (
	"fmt"
)

func main() {
	var firstName string = "Ganteng dong"
	var lastName string
	lastName = "Ukan"

	ganteng := "ukan dong"

	fmt.Println("masuk gan")
	fmt.Println("ganteng", "test")
	fmt.Printf("Hallo %s %s! \n", firstName, lastName)
	fmt.Println("Siapa yang ganteng ? ", ganteng)

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Printf("Iini baru %s %s %s! \n", first, second, third)

	fmt.Println("\n\n CONTOH if else")
	var point int8 = 6
	if point == 8 {
		fmt.Println("Lulus")
	} else if point == 10 {
		fmt.Println("tidak lulus")
	} else {
		fmt.Println("hampir lulus")
	}

	fmt.Println("\n\ncontoh if else lagi")
	if percent := point / 2; percent > 5 {
		fmt.Println("betul dong ", percent)
	}

	fmt.Println("\n\nCONTOH switch case")
	switch point {
	case 10:
		fmt.Println("masuk 10")
	case 20:
		fmt.Println("ini baru 20")
	default:
		fmt.Println("selain10 dan 20")
	}

	fmt.Println("\n\nCONTOH CASE DENGAN BANYAK KONDIS")
	switch point {
	case 10, 20, 30:
		fmt.Println("masuk nik semua hahaha")
	default:
		fmt.Println("masuk default baru nih")
	}

	fmt.Println("\n\nCONTOH SWITCH CASE DENGAN DEFAULT ADA KURUNG KURAWA \n GUNANYA UNTUK BOUNDERING")
	switch point {
	case 10:
		fmt.Println("masuk 10")
	default:
		{
			fmt.Println("ini masuk print")
			fmt.Println("ini baris keduanya")
		}
	}

	fmt.Println("\n\n switch case ternyata bisa juga dengan gaya if else")
	switch {
	case point > 10:
		fmt.Println("ini kondis 1")
	case point <= 10:
		fmt.Println("ini kondisi kedua")

	}

	switch {
	case (point > 3) && (point < 8):
		fmt.Println("masuk awal nih", point)
		fallthrough
	case point < 5:
		fmt.Println("baru masuk kedua", point)
	}

	fmt.Println("\n\n CONOTH LOOPING FOR")
	for i := 0; i < 5; i++ {
		fmt.Println("index ", i)
	}

	fmt.Println("\n\nFOR DENGAN ARGUMENT HANYA KONDISI, KALAU DI BAHASA LAIN WAHILE")
	i := 0
	for i < 3 {
		fmt.Println("contoh pengulangan dengan while for", i)
		i++
	}

	fmt.Println("\n\nFOR DENGAN ARGUMENT HANYA KONDISI, KALAU DI BAHASA LAIN DO WAHILE")
	x := 0
	for {
		fmt.Println("print dong", x)

		x++
		if x == 2 {
			break
		}
	}

	var xs = "123"
	for i, v := range xs {
		fmt.Println("index ", i, "value : ", v)
	}

	var yx = [5]int{10, 20, 30, 40, 50}
	for k, v := range yx {
		fmt.Println("keynya", k)
		fmt.Println("test", v)
	}

	var kvs = map[int]string{
		0: "ssd",
		1: "aaa",
		2: "ccc",
	}
	for key, value := range kvs {
		fmt.Println("ini key nya: ", key, " ini valuenya : ", value)
	}

	for i := 0; i < 5; i++ {
	outherLoop:
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outherLoop
			}
			fmt.Println("matrik [", i, "][", j, "]")
		}
	}

	var numbers = [...]int{1, 2, 3, 4, 5}
	fmt.Println("data array\n", numbers)

	var siswa = []map[string]string{
		{"name": "agus", "gender": "L", "phone": "0898788"},
		{"name": "suryana", "gender": "L", "phone": "8987878"},
		{"name": "wati", "gender": "P", "phone": "212121hghg21"},
	}

	for key, value := range siswa {
		fmt.Println(key, " name: ", value["name"])
		if value["gender"] == "L" {
			fmt.Println(key, " gender: Laki-laki")
		} else {
			fmt.Println(key, " gender: Perempuan")
		}
		fmt.Println(key, " phone: ", value["phone"])
	}
}
