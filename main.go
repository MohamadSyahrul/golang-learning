package main

import (
	"fmt"
	"os"
	"strconv"
)

var biodataTeman = map[int]map[string]string{
	1: {
		"nama":      "Mochamad",
		"alamat":    "Jl.Jombang-Mojokerto No.1",
		"pekerjaan": "Programmer",
		"alasan":    "Ingin belajar lebih banyak tentang Go",
	},
	2: {
		"nama":      "Syahrul",
		"alamat":    "Jl.Jombang-Mojokerto No.2",
		"pekerjaan": "Freelancer",
		"alasan":    "Ingin tau tentang microservice",
	},
	3: {
		"nama":      "Samsudin",
		"alamat":    "Jl.Jombang-Mojokerto No.3",
		"pekerjaan": "Karyawan",
		"alasan":    "Mencari jodoh",
	},
	4: {
		"nama":      "Hayolo",
		"alamat":    "Jl.Jombang-Mojokerto No.4",
		"pekerjaan": "Mahasiswa",
		"alasan":    "Mencari ijazah",
	},
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Nomor absen tidak diberikan")
		return
	}

	no := os.Args[1]

	nomorAbsenInt, err := strconv.Atoi(no)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	teman, nilaiBool := biodataTeman[nomorAbsenInt]
	if !nilaiBool {
		fmt.Println("Nomor absen tersebut tidak ditemukan")
		return
	}

	fmt.Println("Nama:", teman["nama"])
	fmt.Println("Alamat:", teman["alamat"])
	fmt.Println("Pekerjaan:", teman["pekerjaan"])
	fmt.Println("Alasan:", teman["alasan"])
}
