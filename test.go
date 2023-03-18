package main

import (
	"fmt"
	"os"
	"strings"
)

type student struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

type biodataGeneratpr interface {
	biodata() []student
}

type namaBiodata struct {
	nama []string
}

func newNamaBiodata(nama []string) biodataGeneratpr {
	return &namaBiodata{nama: nama}
}

func (n *namaBiodata) biodata() []student {
	address := []string{"Jakarta", "Bandung", "Ambon", "Padang", "Medan"}
	reason := []string{"Milik Abdul", "Milik Budi", "Milik Caca", "Milik Dede", "Milik Erlan"}
	generate := make([]student, 0)
	var s student
	for key, value := range n.nama {
		s.nama = value
		s.alamat = address[key]
		s.alasan = reason[key]
		generate = append(generate, s)
	}
	return generate
}

func main() {
	nama := []string{"Abdul", "Budi", "Caca", "Dede", "Erlan"}
	generator := newNamaBiodata(nama)
	output := generator.biodata()
	var arg string = os.Args[1]
	// argInt, _ := strconv.Atoi(arg)
	var inama int
	// for i, _ := range nama {
	// 	if argInt == i + 1 {
	// 		inama = i
	// 	}
	// }

	for i, v := range nama {
		if strings.ToLower(arg) == strings.ToLower(v) {
			inama = i
		}
	}

	for i, v := range output {
		if inama == i {
			fmt.Println("ID :", i)
			fmt.Println("Alamat :", v.alamat)
			fmt.Println("Alasan :", v.alasan)
		}
	}
}
