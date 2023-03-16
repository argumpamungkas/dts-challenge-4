package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	nama, alamat, pekerjaan, alasan string
}

func biodataStudent(name ...string) []Student {

	address := []string{"Jakarta", "Sumedang", "Bogor", "Bandung", "Bekasi"}
	jobs := []string{"Programmer", "Mobile Dev", "Back-End Dev", "Mahasiswa", "Fullstack Dev"}
	reasons := []string{"Penasaran pada Golang", "Ingin membuat API", "Memperdalam Golang", "Explore Bahasa Pemrograman", "Kebutuhan pekerjaan"}

	var students = make([]Student, 0)
	var s Student

	for i, v := range name {
		s.nama = v
		s.alamat = address[i]
		s.pekerjaan = jobs[i]
		s.alasan = reasons[i]

		students = append(students, s)
	}
	return students
}

func main() {

	name := []string{"Adam", "Argumelar", "Ayu", "Rafli", "Yuda"}
	output := biodataStudent(name...)

	var arg = os.Args[1]
	number, err := strconv.Atoi(arg)

	// Mencari dengan No Absen
	if err != nil {
		fmt.Println("Masukan nomor absen")
	} else if number <= len(name) {
		for i, v := range output {
			if number == i+1 {
				fmt.Println("ID:", i+1)
				fmt.Println("Nama:", v.nama)
				fmt.Println("Alamat:", v.alamat)
				fmt.Println("Pekerjaan:", v.pekerjaan)
				fmt.Println("Alasan memilih kelas Golang:", v.alasan)
			}
		}
	} else {
		fmt.Println("No absen yang dicari tidak ditemukan")
	}

	// MENCARI DENGAN NAMA
	// for i, v := range output {
	// 	if strings.ToLower(arg) == strings.ToLower(v.nama) {
	// 		fmt.Println("ID:", i+1)
	// 		fmt.Println("Nama:", v.nama)
	// 		fmt.Println("Alamat:", v.alamat)
	// 		fmt.Println("Pekerjaan:", v.pekerjaan)
	// 		fmt.Println("Alasan memilih kelas Golang:", v.alasan)
	// 		break
	// 	} else if i == len(name)-1 {
	// 		fmt.Println("Data tidak ditemukan")
	// 	}
	// }
}
