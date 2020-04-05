package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const N = 100

//================Code untuk clear screen===================//
//=========================================================//
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it

	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

//=================End of clear screen====================//

type member struct {
	nama string
	no   string
}

type arrMember struct {
	M [N]member
	G [N]member
	i int
	m int
}

type group struct {
	anggota arrMember
	nama    string
}

type arrGroup struct {
	G [N]group
	i int
}

var (
	dataM arrMember
	dataG arrGroup
)

func main() {
	dataM.M[0].nama = "aqmal"
	dataM.M[0].no = "081234567849"
	dataM.M[1].nama = "ahmad"
	dataM.M[1].no = "081119475843"
	dataM.M[2].nama = "dodi"
	dataM.M[2].no = "082138495047"
	dataM.M[3].nama = "budi"
	dataM.M[3].no = "082347594738"
	dataM.M[4].nama = "nazuah"
	dataM.M[4].no = "081228746392"
	dataM.i = 5

	dataG.G[0].nama = "kelas"
	dataG.G[0].anggota.G[0] = dataM.M[0]
	dataG.G[0].anggota.m = 1

	dataG.G[1].nama = "trio"
	dataG.G[1].anggota.G[0] = dataM.M[0]
	dataG.G[1].anggota.G[1] = dataM.M[2]
	dataG.G[1].anggota.G[2] = dataM.M[4]
	dataG.G[1].anggota.m = 3

	dataG.i = 2

	menu()
}

func menu() {
	CallClear()
	var input int
	for input != 3 {
		fmt.Print(`
==============================================================
		APLIKASI SOCIAL MEDIA A
==============================================================
Program menu :
1. MENU GROUP
2. MENU USER
3. Exit
===============================================================
`)
		fmt.Println("Silahkan Masukan Angka : ")
		fmt.Scan(&input)
		if input == 1 {
			menuGroup()
		} else if input == 2 {
			menuUser()
		} else if input == 3 {
			menuExit()
		}
	}
}

func menuGroup() {
	var input int
	for input != 3 {
		fmt.Print(`
==============================================================
MENU GROUP
==============================================================
Program menu :
1. LIHAT GROUP
2. UNDANG ANGGOTA
3. BUAT GROUP
4. Back
===============================================================
`)
		fmt.Println("Silahkan Masukan Angka : ")
		fmt.Scan(&input)
		if input == 1 {
			lihatGroup()
		} else if input == 2 {
			undangAnggota()
		} else if input == 3 {
			buatGroup()
		} else if input == 4 {
			menu()
		}
	}
}

func lihatGroup() {
	CallClear()
	fmt.Println("===============================================================")
	for i := 0; i <= dataG.i; i++ {
		if dataG.G[i].nama != "" {
			fmt.Println("Nama group : ", dataG.G[i].nama)
		}
		for j := 0; j <= (dataG.G[i].anggota.m); j++ {
			if dataG.G[i].anggota.G[j].nama != "" {
				fmt.Println("Nama anggota : ", dataG.G[i].anggota.G[j].nama, " - ", dataG.G[i].anggota.G[j].no)
			}
		}
		fmt.Println("===============================================================")
	}
}

func undangAnggota() {
	CallClear()
	fmt.Println("===============================================================")
	fmt.Println("Menu Mengundang Anggota Baru")
	fmt.Println("===============================================================")
	var nama string
	var anggota string
	index := 0
	found := false
	fmt.Println("Masukkan nama group : ")
	fmt.Scan(&nama)
	for i := 0; i <= (dataG.i); i++ {
		if nama == dataG.G[i].nama {
			found = true
			index = i
			break
		}
	}
	if found {
		found = false
		indexAnggota := 0
		fmt.Println("===============================================================")
		fmt.Println("Group ", "'", dataG.G[index].nama, "'", " ditemukan")
		fmt.Println("Masukkan nama anggota yang ingin diundang : ")
		fmt.Scan(&anggota)
		for j := 0; j <= (dataM.i); j++ {
			if anggota == dataM.M[j].nama {
				found = true
				indexAnggota = j
				break
			}
		}
		if found {
			dataG.G[index].anggota.m++
			dataG.G[index].anggota.G[dataG.G[index].anggota.m] = dataM.M[indexAnggota]
			fmt.Println(dataM.M[indexAnggota].nama, " berhasil diundang kedalam group")
		} else {
			fmt.Println("Nama tidak ditemukan. Kembali ke menu utama untuk menambah teman.")
			menuGroup()
		}

	} else {
		var pilih string
		fmt.Println("Nama group tidak ditemukan. Buat group baru ? (y/n)")
		fmt.Scan(&pilih)

		if pilih == "y" {
			buatGroup()
		} else {
			menuGroup()
		}
	}
}

func buatGroup() {
	CallClear()
	fmt.Println("===============================================================")
	fmt.Println("Menu Membuat Group Baru")
	fmt.Println("===============================================================")
	var nama string
	index := 0
	found := false
	fmt.Println("Masukkan nama group : ")
	fmt.Scan(&nama)
	for i := 0; i <= (dataG.i); i++ {
		if nama == dataG.G[i].nama {
			found = true
			index = i
			break
		}
	}
	if found {
		fmt.Println("Group ", "'", dataG.G[index].nama, "'", " sudah ada")
	} else {
		dataG.i++
		dataG.G[index].anggota.m = 1
		dataG.G[dataG.i].nama = nama
		dataG.G[dataG.i].anggota.G[0] = dataM.M[0]
		fmt.Println("Group ", "'", dataG.G[dataG.i].nama, "'", " berhasil dibuat")
	}

}

func menuUser() {
	var input int
	for input != 3 {
		fmt.Print(`
==============================================================
MENU USER
==============================================================
Program menu :
1. PROFILE
2. LIHAT TEMAN
3. TAMBAH TEMAN
4. HAPUS TEMAN
5. Back
===============================================================
`)
		fmt.Println("Silahkan Masukan Angka : ")
		fmt.Scan(&input)
		if input == 1 {
			profile()
		} else if input == 2 {
			lihatTeman()
		} else if input == 3 {
			tambahTeman()

		} else if input == 4 {
			hapusTeman()
		} else if input == 5 {
			menu()
		}
	}
}

func hapusTeman() {
	CallClear()
	fmt.Print(`
==============================================================
Menu Hapus Teman
==============================================================
`)
	var nama string
	index := 0
	found := false
	fmt.Println("Masukkan nama yang ingin dihapus : ")
	fmt.Scan(&nama)
	if nama == "aqmal" {
		fmt.Println("Nama tidak bisa dihapus")
	} else {
		for i := 0; i <= (dataM.i); i++ {
			if nama == dataM.M[i].nama {
				found = true
				index = i
				break
			}
		}
		if found {
			for i := 0; i < dataM.i; i++ {
				dataM.M[index] = dataM.M[index+1]
				index++
			}
			dataM.i--
			fmt.Println("Nama berhasil dihapus!")
		} else {
			fmt.Println("Nama tidak ditemukan.")
		}
	}

}

func profile() {
	CallClear()
	fmt.Print(`
==============================================================
Menu Profile
==============================================================
`)
	fmt.Println("Nama : ", dataM.M[0].nama)
	fmt.Println("No HP : ", dataM.M[0].no)
	menuUser()
}

func lihatTeman() {
	CallClear()
	fmt.Print(`
==============================================================
List Teman
==============================================================
`)
	//selection sort asc
	var min int
	for pass := 0; pass <= dataM.i; pass++ {
		min = pass
		for j := pass + 1; j <= dataM.i-1; j++ {
			if dataM.M[min].nama > dataM.M[j].nama {
				min = j
			}
		}
		temp := dataM.M[min]
		dataM.M[min] = dataM.M[pass]
		dataM.M[pass] = temp
	}

	for i := 0; i < dataM.i+1; i++ {
		if dataM.M[i].nama != "aqmal" || dataM.M[i].nama != "" {
			fmt.Println("Nama : ", dataM.M[i].nama, " - ", dataM.M[i].no)
		}
	}
	menuUser()
}

func tambahTeman() {
	CallClear()
	fmt.Print(`
==============================================================
Menu Tambah Teman
==============================================================
`)
	var nama string
	var no string
	index := 0
	found := false
	fmt.Println("Masukkan nomor teman : ")
	fmt.Scan(&no)
	for i := 0; i <= (dataM.i); i++ {
		if no == dataM.M[i].no {
			found = true
			index = i
			break
		}
	}
	if found {
		println("'", dataM.M[index].no, "'", " sudah ada didalam list teman sebagai ", "'", dataM.M[index].nama, "'")
		menuUser()
	} else {
		dataM.i++
		fmt.Println("Simpan nomor sebagai : ")
		fmt.Scan(&nama)
		dataM.M[dataM.i].nama = nama
		dataM.M[dataM.i].no = no
		fmt.Println("==============================================================")
		fmt.Println("Teman berhasil ditambahkan!")

	}
}

func menuExit() {
	dot := ""
	for i := 0; i < 3; i++ {
		dot = dot + "."
		print(dot)
		time.Sleep(1 * time.Second)
	}
	CallClear()
	fmt.Print("Terimakasih telah menggunakan aplikasi ini :)")
	os.Exit(3)
}
