package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type student struct {
	nama string
	alamat string
	pekerjaan string
	alasan string
}

const job = "IT Dev"

func biodata(listNama []string) []student {
	address := []string{"Jakarta", "Bandung", "Ambon", "Padang", "Medan"}
	reason := []string{"Milik Abdul", "Milik Budi", "Milik Caca", "Milik Dede", "Milik Erlan"}
	generate := make([]student, 0)
	var s student
	for key, value := range listNama {
		s.nama = value
		s.alamat = address[key]
		s.pekerjaan = job
		s.alasan = reason[key]
		generate = append(generate, s)
	}
	return generate
}

func main() {
	nama := []string{"Abdul", "Budi", "Caca", "Dede", "Erlan"}
	output := biodata(nama)
	var arg string = os.Args[1]
	var iname int
	argInt, _ := strconv.Atoi(arg)

	// using string input 
	// for i, v := range nama {
	// 	if strings.ToLower(arg) == strings.ToLower(v) { //check Abdul == Abdul
	// 		iname = i
	// 	}
	// }

	// using int input
	for i, _ := range nama {
		if argInt == i + 1 { //check 1 == 1
			iname = i
		}
	}

	for i, v := range output {
		if iname == i {
			fmt.Println("ID :", i)
			fmt.Println("Alamat :", v.alamat)
			fmt.Println("Pekerjaan :", v.pekerjaan)
			fmt.Println("Alasan :", v.alasan)
		}
	}
}