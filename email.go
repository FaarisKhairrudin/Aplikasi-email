/*
	Judul Tubes : membuat aplikasi email

	Anggota Kelompok:
	1. Faaris Khairrudin
	2. Khairan Dwieza Ansari
*/

package main

import (
	"fmt"
	"os"      //! import untuk melakukan clear screen
	"os/exec" //!            ~~~~~~~~
	"runtime" //!            ~~~~~~~~
	"time"    //! import untuk melakukan jeda waktu
)

//? mendekrasikan batas kapasitas suatu aray yaitu 256
const NMAX int = 256

//! variabel untuk menyimpan string yang berisikan kode warna terminal
var (
    black       = "\033[30m"
    red         = "\033[31m"
    green       = "\033[32m"
    yellow      = "\033[33m"
    blue        = "\033[34m"
    magenta     = "\033[35m"
    cyan        = "\033[36m"
    white       = "\033[37m"
    lightGray   = "\033[37m"
    lightGreen  = "\033[92m"
    lightBlue   = "\033[94m"
    lightCyan   = "\033[96m"
    lightMagenta= "\033[95m"
    boldRed     = "\033[1;31m"
    boldGreen   = "\033[1;32m"
    boldYellow  = "\033[1;33m"
    boldBlue    = "\033[1;34m"
    boldMagenta = "\033[1;35m"
    boldCyan    = "\033[1;36m"
    boldWhite   = "\033[1;37m"
    reset       = "\033[0m"
)

// * Struct untuk menyimpan informasi pengguna
type akun struct {			
	id int
	username string
	email    string
	password string
	admin    bool
	approved bool
	ada      bool
	processed bool
	inbox 	tabPesan
}


// * Struct untuk menyimpan informasi email
type pesan struct {			
	id		  int
	sender    string
	recipient string
	subject   string
	body      string
	ada       bool
	balasan   pesanBalas
}

type pesanBalas struct {
	sender    string
	recipient string
	subject   string
	body      string
	ada       bool
}

type tabAkun [NMAX]akun
type tabPesan [NMAX]pesan

//? fungsi sebagai inisiasi awal untuk fungsi lainya
func main() {	
	welcome()
	mainMenu()
	goodbye()
}

//? fungsi untuk mendeklarasikan akun bawaan (sampel)
func deklarasiAkun(user *tabAkun, jumAkun *int){
	user[0] = akun{
		id 			: 0,
		ada			: true,
		username	: "admin",
		email		: "admin@email.com",
		password	: "admin",
		admin		: true,
		approved	: true,
		processed	: true,
	}

	user[1] = akun{
		id			: 1,
		ada			: true,
		username	: "faaris",
		email		: "faaris@email.com",
		password	: "faaris",
		admin		: false,
		approved	: true,
		processed  	: true,
	}
	user[1].inbox[0] = pesan{
		id			: 0,
		sender		: "admin@email.com",
		recipient	: "faaris@email.com",
		subject		: "testing pertama",
		body		: "halo ini adalah tes email",
		ada			: true,
	}
	user[1].inbox[1] = pesan{
		id			: 1,
		sender		: "fatih@email.com",
		recipient	: "faaris@email.com",
		subject		: "testing kedua",
		body		: "halo ini adalah tes email",
		ada			: true,
	}
	user[1].inbox[2] = pesan{
		id			: 2,
		sender		: "abil@email.com",
		recipient	: "faaris@email.com",
		subject		: "testing ketiga",
		body		: "halo ini adalah tes email lagi",
		ada			: true,
	}
	user[1].inbox[3] = pesan{
		id			: 3,
		sender		: "fadhil@email.com",
		recipient	: "faaris@email.com",
		subject		: "testing keempat",
		body		: "halo ini adalah tes email lagi dan lagi",
		ada			: true,
	}
	user[1].inbox[4] = pesan{
		id			: 4,
		sender		: "reza@email.com",
		recipient	: "faaris@email.com",
		subject		: "testing kelima",
		body		: "halo ini adalah tes email lagi dan lagi dan lagi",
		ada			: true,
	}

	user[2] = akun{
		id			: 2,
		ada			: true,
		username	: "fauzan",
		email		: "fauzan@email.com",
		password	: "fauzan",
		admin		: false,
		approved	: false,
		processed  	: false,
	}	
	user[3] = akun{
		id			: 3,
		ada			: true,
		username	: "farel",
		email		: "farel@email.com",
		password	: "farel",
		admin		: false,
		approved	: false,
		processed  	: true,
	}

	user[4] = akun{
		id			: 4,
		ada			: true,
		username	: "fatih",
		email		: "fatih@email.com",
		password	: "fatih",
		admin		: false,
		approved	: true,
		processed  	: true,
	}
	user[4].inbox[0] = pesan{
		id			: 0,
		sender		: "admin@email.com",
		recipient	: "fatih@email.com",
		subject		: "selamat datang",
		body		: "Selamat datang di sistem kami, Fatih.",
		ada			: true,
	}

	user[5] = akun{
		id			: 5,
		ada			: true,
		username	: "abil",
		email		: "abil@email.com",
		password	: "abil",
		admin		: false,
		approved	: false,
		processed  	: false,
	}

	user[6] = akun{
		id			: 6,
		ada			: true,
		username	: "fadhil",
		email		: "fadhil@email.com",
		password	: "fadhil",
		admin		: false,
		approved	: true,
		processed  	: true,
	}
	user[6].inbox[0] = pesan{
		id			: 0,
		sender		: "admin@email.com",
		recipient	: "fadhil@email.com",
		subject		: "selamat datang",
		body		: "Selamat datang di sistem kami, Fadhil.",
		ada			: true,
	}

	user[7] = akun{
		id			: 7,
		ada			: true,
		username	: "reza",
		email		: "reza@email.com",
		password	: "reza",
		admin		: false,
		approved	: true,
		processed  	: true,
	}
	user[7].inbox[0] = pesan{
		id			: 0,
		sender		: "admin@email.com",
		recipient	: "reza@email.com",
		subject		: "selamat datang",
		body		: "Selamat datang di sistem kami, reza.",
		ada			: true,
	}

	user[8] = akun{
		id			: 8,
		ada			: true,
		username	: "fahmi",
		email		: "fahmi@email.com",
		password	: "fahmi",
		admin		: false,
		approved	: false,
		processed  	: true,
	}

	user[9] = akun{
		id			: 9,
		ada			: true,
		username	: "firman",
		email		: "firman@email.com",
		password	: "firman",
		admin		: false,
		approved	: false,
		processed  	: false,
	}

	*jumAkun = 10
}

//? fungsi utama dalam program ini
func mainMenu() {
	var user tabAkun
	var index int
	var pilihan,pilihanMenu string
	var login bool
	var jumAkun int = 0

	// pilihanMenu = -1
	deklarasiAkun(&user, &jumAkun)
	menuLogin()
	fmt.Scan(&pilihan)
	
	for pilihan != "3" {
		switch pilihan {
		case "1":
			signIn(&user, &login, &index, jumAkun)
			if login {
				cls()
				if user[index].admin {
					menuAdmin()
				} else {
					menu()
				}
				fmt.Scan(&pilihanMenu)
				for pilihanMenu != "0" {
					switch pilihanMenu {
					case "1":
						kirimEmail(&user, index, jumAkun)
					case "2":
						lihatInbox(&user, index, jumAkun)
					case "3":
						if user[index].admin == true {
							persetujuanAkun(&user, jumAkun)
						} else {
							cls()
							fmt.Println("Anda bukan admin")
							jedaWaktu(1000)
						}
					default:
						cls()
						fmt.Println("Pilihan tidak valid.")
						jedaWaktu(1000)
					}
					// pilihanMenu = -1
					if user[index].admin {
						menuAdmin()
					} else {
						menu()
					}
					fmt.Scan(&pilihanMenu)
				}
				login = false
			}
		case "2":
			signUp(&user, &jumAkun)
		default:
			cls()
			fmt.Println("Pilihan tidak valid.")
			jedaWaktu(1000)
		}
		// pilihan = -1
		menuLogin()
		fmt.Scan(&pilihan)
	}
}

//? fungsi untuk menampilkan tampilan menu pada awal login
func menuLogin() {
	cls()
	headerLogin()
	fmt.Println("1. Sign in")
	fmt.Println("2. Sign up")
	fmt.Println("3. keluar")
	fmt.Println()
	fmt.Print("Pilihan Anda: ")
}

//? fungsi untuk menampilkan tampilan menu untuk user setelah login berhasil
func menu() {
	cls()
	headerMainMenu()
	fmt.Println(lightGreen+ "      M A I N   M E N U"+reset)
	fmt.Println()
	fmt.Println("1. Kirim Email")
	fmt.Println("2. Lihat Inbox")
	fmt.Println("0. Keluar")
	fmt.Println()
	fmt.Print("Pilihan Anda: ")
}

//? fungsi untuk menampilkan tampilan menu untuk admin setelah login berhasil
func menuAdmin() {
	cls()
	headerMainMenu()
	// fmt.Println("=== Menu Utama ===")
	fmt.Println(lightGreen+ "      M A I N   M E N U"+reset)
	fmt.Println()
	fmt.Println("1. Kirim Email")
	fmt.Println("2. Lihat Inbox")
	fmt.Println("3. Verifikasi Akun (Admin)")
	fmt.Println("0. Keluar")
	fmt.Println()
	fmt.Print("Pilihan Anda: ")
}

//? fungsi untuk melakukan pengecekan input user terhadap akun yang sudah terdaftar
func signIn(a *tabAkun, login *bool, index *int, jumAkun int) {
	var i int
	var password, email string
	var stop bool

	cls()
	headerSignin()
	fmt.Print("Email		: ")
	fmt.Scan(&email)
	fmt.Print("Password	: ")
	fmt.Scan(&password)

	for i = 0; i < jumAkun && !stop; i++ {
		if a[i].email == email && a[i].password == password {
			if a[i].approved == true && a[i].processed == true {
				cls()
				fmt.Println(lightGreen+"\n\n\n        Login Berhasil."+reset)
				jedaWaktu(1000)
				*login = true
				*index = i
				stop = true
			} else if a[i].approved == false && a[i].processed == false {
				cls()
				fmt.Println(boldYellow+"\n\n\n akun anda belum diverifikasi oleh admin"+reset)
				jedaWaktu(2000)
				stop = true
			} else if a[i].approved == false && a[i].processed == true{
				cls()
				fmt.Println(boldRed+"\n\n\n   registrasi anda ditolak oleh admin"+ reset)
				jedaWaktu(2000)
				stop = true
			}
		} else if i == jumAkun-1 {
			cls()
			fmt.Println(red+"\n\n\n email atau password yang anda masukan salah"+reset)		
			jedaWaktu(2000)
		}
	}
}

//? fungsi untuk membuat akun 
func signUp(a *tabAkun, jumAkun *int) {
	var i int
	var username, password, email string
	var stop bool
	var pilihan string

	cls()
	headerSignUp()
	fmt.Println(lightCyan+"   lengkapi data data berikut"+reset)
	fmt.Println()
	fmt.Println(boldYellow+"* "+reset+"buat username tanpa menggunakan spasi")
	fmt.Print("Username	: ")
	fmt.Scan(&username)
	
	fmt.Println()
	fmt.Println(boldYellow+"* "+reset+"buat email dan gunakan akhiran"+boldYellow+ " @email.com"+reset)
	fmt.Print("Email		: ")
	fmt.Scan(&email)
	
	fmt.Println()
	fmt.Println(boldYellow+"* "+reset+"buatlah password untuk keamanan akun anda"+reset)
	fmt.Print("Password	: ")
	fmt.Scan(&password)
	
	stop = false

	for i = 0; i < NMAX && !stop; i++ {
		if a[i].ada == false && a[i].email != email {
			fmt.Println("\nApakah data yang dimasukan sudah benar?")
			fmt.Print("pilihan anda (y/n): ")
			fmt.Scan(&pilihan)

			if pilihan == "y" || pilihan == "Y"{
				a[i].id = i
				a[i].username = username
				a[i].email = email
				a[i].password = password
	
				a[i].ada = true
				a[i].processed = false
	
				*jumAkun++
	
				cls()
				fmt.Print(lightGreen+"\n\n\n  Registrasi Berhasil "+reset)
				jedaWaktu(1500)
				fmt.Println()
			} else if pilihan == "n" || pilihan == "N" {
				cls()
				fmt.Print(red+"\n\n\n  Registrasi Dibatalkan "+reset)
				jedaWaktu(1500)
				fmt.Println()
			} else {
				cls()
				fmt.Print(red+"\n\n\n  pilihan tidak valid "+reset)
				jedaWaktu(1500)
				fmt.Println()

			}
			stop = true
		} else if a[i].email == email {
			stop = true
			cls()
			fmt.Print(boldRed+"\n\n\n  Alamat Email telah digunakan "+reset)
			jedaWaktu(1500)
			fmt.Println()
		}
	}
}

//? fungsi untuk memproses verifikasi akun (khusus admin)
func persetujuanAkun(a *tabAkun,jumAkun int) {
	var tempAkun 	tabAkun
	var idPilihan 	int	 = -4
	var pilihanPersetujuan 	int = -4
	var i int

	tempAkun = *a
	cls()
	headerMainMenu()
	fmt.Println(lightGreen + "V E R I F I K A S I   A K U N​​​​​" + reset)
	fmt.Println()
	fmt.Println("daftar username yang terdaftar:\n")
	for i = 0; i < jumAkun; i++ {
		if a[i].approved == false && a[i].processed == false{
			fmt.Println(lightBlue+"- ID		:", a[i].id, reset)
			fmt.Println("  Username	:", a[i].username)
			fmt.Println("  Email		:", a[i].email)
		}
	}
	fmt.Println("\nketik '-1' untuk ascending berdasarkan id")
	fmt.Println("ketik '-2' untuk descending berdasarkan id")
	fmt.Println("ketik '-3' untuk mengedit status verifikasi")
	fmt.Println("ketik '0' untuk kembali")
	fmt.Print("Masukkan id pengguna yang akan disetujui / ditolak: ")
	fmt.Scan(&idPilihan)
	for idPilihan != 0{	
		if idPilihan > 0 && idPilihan < NMAX && a[idPilihan].approved == false && a[idPilihan].processed == false && a[idPilihan].ada == true {
			cls()
			headerMainMenu()
			fmt.Println(lightGreen + "V E R I F I K A S I   A K U N​​​" + reset)
			fmt.Println("\nberikut adalah detail dari akun yang ingin anda proses:\n")
			fmt.Println(lightCyan + " ID		:", a[idPilihan].id)
			fmt.Println(" Username	:", a[idPilihan].username)
			fmt.Println(" Email		:", a[idPilihan].email, reset)
			fmt.Println("\nketik '1' untuk "+ lightGreen +"menyetujui"+ reset + " verifikasi akun")
			fmt.Println("ketik '2' untuk"+ red +" menolak "+ reset +" verifikasi akun")
			fmt.Println("ketik '0' untuk kembali")
			fmt.Print("Pilihan anda (0/1/2): ")
			fmt.Scan(&pilihanPersetujuan)
			
			for pilihanPersetujuan != 0 {
				switch pilihanPersetujuan {
				case 1:
					a[idPilihan].approved = true
					a[idPilihan].processed = true
					cls()
					fmt.Println(lightBlue + a[idPilihan].username + reset, "berhasil disetujui.")
					jedaWaktu(1000)
					pilihanPersetujuan = 0
				case 2:
					cls()
					fmt.Println(lightBlue + a[idPilihan].username + reset, "berhasil ditolak.")
					jedaWaktu(1000)
					a[idPilihan].processed = true
					
					pilihanPersetujuan = 0
				default:
					pilihanPersetujuan = -4
					cls()
					fmt.Println("Input tidak valid")
					jedaWaktu(1000)
					cls()
					headerMainMenu()
					fmt.Println(lightGreen + "V E R I F I K A S I   A K U N​​​" + reset)
					fmt.Println("\nberikut adalah detail dari akun yang ingin anda proses:\n")
					fmt.Println(lightCyan + " ID		:", a[idPilihan].id)
					fmt.Println(" Username	:", a[idPilihan].username)
					fmt.Println(" Email		:", a[idPilihan].email, reset)
					fmt.Println("\nketik '1' untuk "+ lightGreen +"menyetujui"+ reset + " verifikasi akun")
					fmt.Println("ketik '2' untuk"+ red +" menolak "+ reset +" verifikasi akun")
					fmt.Println("ketik '0' untuk kembali")
					fmt.Print("Pilihan anda (0/1/2): ")
					fmt.Scan(&pilihanPersetujuan)
				}

			}
			pilihanPersetujuan = -4
			tempAkun = *a
		} else if idPilihan == -1{
			sortAscendingId(&tempAkun, jumAkun)
		} else if idPilihan == -2{
			sortDescendingId(&tempAkun, jumAkun)
		} else if idPilihan == -3{
			editProsesAkun(&*a, jumAkun)
		} else {
			cls()
			fmt.Println("ID tidak terdapat pada list, silahkan masukan ID yang ada pada list")
			jedaWaktu(2000)
		}

		idPilihan = -4 // untuk mereset id ketika dimasukan selain int
		cls()
		headerMainMenu()
		fmt.Println(lightGreen + "V E R I F I K A S I   A K U N​​​​​" + reset)
		fmt.Println()
		fmt.Println("daftar username yang terdaftar:\n")
		for i = 0; i < jumAkun; i++ {
			if tempAkun[i].approved == false && tempAkun[i].processed == false {
				fmt.Println(lightBlue+"- ID		:", tempAkun[i].id, reset)
				fmt.Println("  Username	:", tempAkun[i].username)
				fmt.Println("  Email		:", tempAkun[i].email)
			}
		}
		fmt.Println("\nketik '-1' untuk ascending berdasarkan id")
		fmt.Println("ketik '-2' untuk descending berdasarkan id")
		fmt.Println("ketik '-3' untuk mengedit status verifikasi")
		fmt.Println("ketik '0' untuk kembali")
		fmt.Print("Masukkan id pengguna yang akan disetujui / ditolak: ")
		fmt.Scan(&idPilihan)		

	}
}

//? fungsi untuk mengubah status persetujuan akun (diterima/ditolak)
func editProsesAkun(a *tabAkun, jumAkun int) {
	var i int
	var pilihan int = -1
	var pilihan2 string
	cls()
	headerMainMenu()
	fmt.Println(lightGreen + "V E R I F I K A S I   A K U N​​​​​" + reset)
	fmt.Println()
	fmt.Println("daftar akun yang sudah diproses:\n")

	for i = 0; i<jumAkun; i++ {
		if a[i].processed == true {
			fmt.Println(lightBlue+"- ID		:", a[i].id, reset)
			fmt.Println("  Username	:", a[i].username)
			fmt.Println("  Email		:", a[i].email)
			if a[i].approved{
				fmt.Println("  Status	:", boldGreen+"disetujui"+reset)
			} else {
				fmt.Println("  Status	:", boldRed+"ditolak"+reset)	
			}
		}
	}
	
	fmt.Println("\nketik '0' untuk kembali")
	fmt.Print("Pilih id yang akan diedit statusnya: ")
	fmt.Scan(&pilihan)
	for pilihan != 0 {
		if a[pilihan].processed == true {
			cls()
			headerMainMenu()
			fmt.Println(lightGreen + "V E R I F I K A S I   A K U N​​​​​" + reset)
			fmt.Println()
			fmt.Println(lightBlue+"- ID		:", a[pilihan].id, reset)
			fmt.Println("  Username	:", a[pilihan].username)
			fmt.Println("  Email		:", a[pilihan].email)
			if a[pilihan].approved{
				fmt.Println("  Status	:", boldGreen+"disetujui"+reset)
			} else {
				fmt.Println("  Status	:", boldRed+"ditolak"+reset)	
			}

			fmt.Println()
			fmt.Println("\nketik '1' untuk "+ lightGreen +"menyetujui"+ reset + " verifikasi akun")
			fmt.Println("ketik '2' untuk"+ red +" menolak "+ reset +" verifikasi akun")
			fmt.Println("ketik '0' untuk kembali")
			fmt.Print("pilihan anda: ")
			fmt.Scan(&pilihan2)
			
			for pilihan2 != "0" {
				switch pilihan2 {
				case "1":
					a[pilihan].approved = true
					cls()
					fmt.Println("status",lightBlue + a[pilihan].username + reset, "berhasil disetujui.")
					jedaWaktu(2000)
					pilihan2 = "0"
				case "2":
					cls()
					fmt.Println("status",lightBlue + a[pilihan].username + reset, "berhasil ditolak.")
					jedaWaktu(2000)
					a[pilihan].approved = false
					
					pilihan2 = "0"
				default:
					cls()
					fmt.Println("Input tidak valid")
					jedaWaktu(1000)
					cls()
					headerMainMenu()
					fmt.Println(lightGreen + "V E R I F I K A S I   A K U N​​​​​" + reset)
					fmt.Println()
					fmt.Println(lightBlue+"- ID		:", a[pilihan].id, reset)
					fmt.Println("  Username	:", a[pilihan].username)
					fmt.Println("  Email		:", a[pilihan].email)
					if a[pilihan].approved{
						fmt.Println("  Status:	", boldGreen+"disetujui"+reset)
					} else {
						fmt.Println("  Status	:", boldRed+"ditolak"+reset)	
					}

					fmt.Println()
					fmt.Println("\nketik '1' untuk "+ lightGreen +"menyetujui"+ reset + " verifikasi akun")
					fmt.Println("ketik '2' untuk"+ red +" menolak "+ reset +" verifikasi akun")
					fmt.Println("ketik '0' untuk kembali")
					fmt.Print("pilihan anda: ")
					fmt.Scan(&pilihan2)
				}

			}
		} else {
			cls()
			fmt.Println("Input tidak valid")
			jedaWaktu(1000)
		}
		pilihan = -1
		cls()
		headerMainMenu()
		fmt.Println(lightGreen + "V E R I F I K A S I   A K U N​​​​​" + reset)
		fmt.Println()
		fmt.Println("daftar akun yang sudah diproses:")
		fmt.Println()
		for i = 0; i<jumAkun; i++ {
			if a[i].processed == true {
				fmt.Println(lightBlue+"- ID		:", a[i].id, reset)
				fmt.Println("  Username	:", a[i].username)
				fmt.Println("  Email		:", a[i].email)
				if a[i].approved{
					fmt.Println("  Status	:", boldGreen+"disetujui"+reset)
				} else {
					fmt.Println("  Status	:", boldRed+"ditolak"+reset)	
				}
			}
		}
		fmt.Println()
		fmt.Println("\nketik '0' untuk kembali")
		fmt.Print("Pilih id yang akan diedit statusnya: ")
		fmt.Scan(&pilihan)
		
	}
}

//? fungsi untuk mengirimkan email antara satu akun kepada akun lainya
func kirimEmail(a *tabAkun, index, jumAkun int) {
	var recipient, subject, body string
	var i, j int
	var stop bool

	cls()
	headerMainMenu()
	fmt.Println(lightGreen + "   K I R I M   E M A I L​​​​​" + reset)
	fmt.Println()
	fmt.Print("Pengirim    : ")
	fmt.Println(a[index].email)
	fmt.Print("Penerima    : ")
	fmt.Scan(&recipient) 		// tidak boleh ada spasi setelah mengetik penerima nya
	fmt.Print("Subjek      : ")
	inputSubject(&subject)
	fmt.Print("Isi Pesan   : ")
	inputPesan(&body)

	stop = false 
	i = 0
	
	//? binary search untuk mencari pada email yang berisi recipient
	if binarySearch(&*a, jumAkun, recipient,&i) {
		j = 0
		for !stop && j < NMAX {
			if !a[i].inbox[j].ada {
				a[i].inbox[j] = pesan{
					id :	   j,
					sender:    a[index].email,
					recipient: recipient,
					subject:   subject,
					body:      body,
					ada:       true,
				}
				cls()
				fmt.Println("Email berhasil terkirim.")
				jedaWaktu(1000)
				stop = true
			}
			j++
		}
		if !stop {
			cls()
			fmt.Println("Inbox penerima penuh.")
			stop = true
			jedaWaktu(1000)
		}
	} else {
		cls()
		fmt.Println("Email penerima tidak ditemukan.")
		jedaWaktu(1000)
	}
}

//? fungsi untuk melihat inbox / pesan yang masuk atas nama email suatu akun
func lihatInbox(a *tabAkun, index, jumAkun int) {
	var pilihanInbox int = -3
	var pilihanInbox2 int = -3
	var indexInbox int
	var pengirim string
	var stop bool
	var tempAkun tabAkun
	
	tempAkun = *a
	cls()
	menuInbox(&*a, index)
	fmt.Scan(&pilihanInbox)
	for pilihanInbox != 0{
		cetakPesan(&*a, index, pilihanInbox,&stop,&indexInbox, &pengirim)
		
		if pilihanInbox == -2 {
			sortAscendingSender(&tempAkun, index, jumAkun)
		} else if pilihanInbox == -3{
			sortDescendingSender(&tempAkun, index, jumAkun)
		}else if pilihanInbox == -1{
			cetakSemuaPesan(&*a, index, pilihanInbox,&indexInbox)
		} else {
			if !stop {
				cls()
				fmt.Println("Nomor inbox tidak valid.")
				jedaWaktu(1000)
			} else if stop{
				fmt.Println("\nKetik '1' untuk "+ lightGreen +"membalas"+ reset +" pesan")
				fmt.Println("Ketik '2' untuk "+ red +"menghapus"+ reset +" pesan")
				fmt.Println("Ketik '0' untuk kembali")
				fmt.Print("Pilihan anda (0/1/2): ")
				fmt.Scan(&pilihanInbox2)
				switch pilihanInbox2{
				case 1: 
					balasPesan(&*a, index, jumAkun, pengirim,indexInbox)
				case 2: 
					hapusPesan(&*a, index, indexInbox)
					tempAkun = *a
					cls()
					fmt.Println("Pesan berhasil dihapus")
					jedaWaktu(1000)
				case 0:

				default: 
					cls()
					fmt.Println("Nomor tidak Valid")
					jedaWaktu(1000)
					cls()
				}				
			}
		}

		stop = false
	
		pilihanInbox = -4 // untuk mereset pilihan inbox ketika di inputnya bukan int
		pilihanInbox2 = -4 // untuk mereset pilihan inbox ketika di inputnya bukan int
		cls()
		menuInbox(&tempAkun, index)
		fmt.Scan(&pilihanInbox)
	}
}

//? fungsi untuk menampilkan menu pada bagian lihat inbox
func menuInbox(a *tabAkun, index int) {
	var i int

	headerMainMenu()
	fmt.Println(lightGreen + "         I N B O X​​​​​" + reset)
	fmt.Println()
	fmt.Println("berikut adalah daftar pesan yang masuk:\n")
	for i = 0; i < NMAX; i++ {
		if a[index].inbox[i].ada{
			fmt.Println(lightBlue + "- ID 		:",a[index].inbox[i].id+1, reset)
			fmt.Println(lightCyan+"  Pengirim	:", a[index].inbox[i].sender+reset)
			fmt.Println("  Subjek	:", a[index].inbox[i].subject)
		}
	}
	
	fmt.Println("\nKetik '-1' untuk menampilkan isi semua pesan yang masuk")
	fmt.Println("Ketik '-2' untuk "+boldGreen+"ascending"+reset+" berdasarkan pengirim")
	fmt.Println("Ketik '-3' untuk "+boldGreen+"descending"+reset+" berdasarkan pengirim")
	fmt.Println("Ketik '0' untuk kembali")
	fmt.Print("Pilih id pesan yang ingin dilihat: ")
}

//? fungsi untuk menampilkan suatu pesan yang masuk pada suatu akun
func cetakPesan(a *tabAkun, index,pilihanInbox int, stop *bool, indexInbox *int,pengirim *string) {
	var i int
	for i = 0; i < NMAX && !*stop ; i++ { 
		if a[index].inbox[i].id == pilihanInbox - 1 &&  a[index].inbox[i].ada{
			cls()
			headerMainMenu()
			fmt.Println(lightGreen + "         I N B O X​​​​​" + reset)
			fmt.Println()
			fmt.Println("+ + + + + + + + + + + + + + + + +\n")
			if a[index].inbox[i].balasan.ada {
				fmt.Println(boldYellow + a[index].inbox[i].sender + reset,"membalas email anda:")
				fmt.Print(cyan +" Pengirim	: ")
				fmt.Println(a[index].inbox[i].balasan.sender)
				fmt.Print(" Penerima	: ")
				fmt.Println(a[index].inbox[i].balasan.recipient)
				fmt.Print(lightCyan+ " Subjek		: ")
				fmt.Println(a[index].inbox[i].balasan.subject)
				fmt.Print(" Isi Pesan	: ")
				fmt.Println(a[index].inbox[i].balasan.body, reset)
				fmt.Println()
				fmt.Println("dengan balasan:")
				fmt.Print(cyan +">> Pengirim	: ")
				fmt.Println(a[index].inbox[i].sender)
				fmt.Print("   Penerima	: ")
				fmt.Println(a[index].email)
				fmt.Print(lightCyan+ "   Subjek	: ")
				fmt.Println(a[index].inbox[i].subject)
				fmt.Print("   Isi Pesan	: ")
				fmt.Println(a[index].inbox[i].body , reset)
			} else {
				fmt.Print(cyan +" Pengirim	: ")
				fmt.Println(a[index].inbox[i].sender)
				fmt.Print(" Penerima	: ")
				fmt.Println(a[index].email)
				fmt.Print(lightCyan+ " Subjek		: ")
				fmt.Println(a[index].inbox[i].subject)
				fmt.Print(" Isi Pesan	: ")
				fmt.Println(a[index].inbox[i].body , reset)
			}
			*pengirim = a[index].inbox[i].sender // untuk dijadikan pengirim pada saat balas pesan
			*stop = true			
			*indexInbox = i	
		}
	}
}

//? fungsi untuk menampilkan semua pesan yang masuk pada suatu akun
func cetakSemuaPesan(a *tabAkun, index,pilihanInbox int, indexInbox *int) {
	var i int
	var pilihan string

	for pilihan != "0"{
		cls()
		headerMainMenu()
		fmt.Println(lightGreen + "         I N B O X​​​​​" + reset)
		fmt.Println()
		fmt.Println("+ + + + + + + + + + + + + + + + +\n")
		for i = 0; i < NMAX; i++ { 
			if a[index].inbox[i].ada == true {
				fmt.Print(cyan +" Pengirim	: ")
				fmt.Println(a[index].inbox[i].sender)
				fmt.Print(" Penerima	: ")
				fmt.Println(a[index].email)
				fmt.Print(lightCyan+ " Subjek		: ")
				fmt.Println(a[index].inbox[i].subject)
				fmt.Print(" Isi Pesan	: ")
				fmt.Println(a[index].inbox[i].body , reset)
				fmt.Println()		
				*indexInbox = i	
			}
		}
		fmt.Println("masukan '0' untuk kembali")
		fmt.Scan(&pilihan)

	}
}

//? fungs untuk membalas pesan yang telah dikirim oleh suatu akun
func balasPesan(a *tabAkun, index, jumAkun int, pengirim string, indexInbox int){
	var subject, body string
	var i, j int
	var stop bool

	cls()
	headerMainMenu()

	fmt.Println(lightGreen + "   B A L A S   E M A I L​​​​​" + reset)
	fmt.Println()
	fmt.Print("Pengirim    : ")
	fmt.Println(a[index].email)
	fmt.Print("Penerima    : ")
	fmt.Println(pengirim)
	fmt.Print("Subjek      : ")
	inputSubject(&subject)
	fmt.Print("Isi Pesan   : ")
	inputPesan(&body)

	stop = false
	i = 0
	for !stop && i < jumAkun {
		if a[i].email == pengirim {
			j = 0
			for !stop && j < NMAX {
				if !a[i].inbox[j].ada {
					a[i].inbox[j] = pesan{
						id :	   j,
						sender:    a[index].email,
						recipient: pengirim,
						subject:   subject,
						body:      body,
						ada:       true,
					}
					a[i].inbox[j].balasan = pesanBalas{
						sender 		: a[index].inbox[indexInbox].sender,
						recipient 	: a[index].inbox[indexInbox].recipient,
						subject 	: a[index].inbox[indexInbox].subject,
						body 		: a[index].inbox[indexInbox].body,
						ada 		: true,
					} 
					cls()
					fmt.Println("Email berhasil terkirim.")
					jedaWaktu(1000)
					stop = true
				}
				j++
			}
			if !stop {
				cls()
				jedaWaktu(1000)
				fmt.Println("Inbox penerima penuh.")
				stop = true
			}
		}
		i++
	}

	if !stop {
		fmt.Println("Email penerima tidak ditemukan.")
	}
}

//? fungsi sorting ascending menggunakan selection sort berdasarkan nama pengirim
func sortAscendingSender(a *tabAkun, index, jumAkun int) {
	var i,j int
	var idx_min int
	var t pesan
	
	i = 1
	for i <= jumAkun - 1 {
 		idx_min = i - 1
 		j = i
 		for j < jumAkun {
			 if a[index].inbox[idx_min].sender >= a[index].inbox[j].sender {
 				idx_min = j
 			}
 			j = j + 1
 		}
 		t = a[index].inbox[idx_min] 
 		a[index].inbox[idx_min] = a[index].inbox[i-1]
 		a[index].inbox[i-1] = t
 		i = i + 1
	}
}

//? fungsi sorting ascending menggunakan selection sort berdasarkan email
func sortAscendingEmail(a *tabAkun, jumAkun int) {
	var i,j int
	var idx_min int
	var t akun
	
	i = 1
	for i <= jumAkun - 1 {
 		idx_min = i - 1
 		j = i
 		for j < jumAkun {
			 if a[idx_min].email >= a[j].email {
 				idx_min = j
 			}
 			j = j + 1
 		}
 		t = a[idx_min]
 		a[idx_min] = a[i-1]
 		a[i-1] = t
 		i = i + 1
	}
}

//? fungsi sorting ascending menggunakan selection sort berdasarkan id
func sortAscendingId(a*tabAkun, jumAkun int) {
	var i,j int
	var idx_min int
	var t akun
	
	i = 1
	for i <= jumAkun - 1 {
 		idx_min = i - 1
 		j = i
 		for j < jumAkun {
			 if a[idx_min].id >= a[j].id {
 				idx_min = j
 			}
 			j = j + 1
 		}
 		t = a[idx_min]
 		a[idx_min] = a[i-1]
 		a[i-1] = t
 		i = i + 1
	}
}

//? fungsi sorting descending menggunakan insertion sort berdasarkan nama pengirim
func sortDescendingSender(a *tabAkun, index, jumAkun int) {
	var i,j int
	var temp pesan
	
	i = 1
	for i <= jumAkun-1 {
		j = i
		temp = a[index].inbox[j]
		for j > 0 && temp.sender > a[index].inbox[j-1].sender {
			a[index].inbox[j] = a[index].inbox[j-1]
			j = j - 1
		}
		a[index].inbox[j] = temp
		i = i + 1
	}
}

//? fungsi sorting descending menggunakan insertion sort berdasarkan id
func sortDescendingId(a *tabAkun, jumAkun int) {
	var i,j int
	var temp akun
	
	i = 1
	for i <= jumAkun-1 {
		j = i
		temp = a[j]
		for j > 0 && temp.id > a[j-1].id {
			a[j] = a[j-1]
			j = j - 1
		}
		a[j] = temp
		i = i + 1
	}
}

//? fungsi untuk mencari nilai yang email string pada variabel penerima di dalam array tabAkun
func binarySearch(a *tabAkun, jumAkun int, penerima string, i *int) bool {
	var tempAkun tabAkun
	var kr, kn, med int
	var found bool

	tempAkun = *a
	sortAscendingEmail(&tempAkun, jumAkun)
	kr = 0
	kn = jumAkun-1
	found = false
	for kr <= kn && !found {
 		med = (kr+kn) / 2
 		if tempAkun[med].email < penerima {
 			kr = med + 1
 		} else if tempAkun[med].email > penerima {
 			kn = med - 1
 		} else {
 			found = true
			*i = tempAkun[med].id
 		}
	}
	return found
}

//? fungsi untuk menghapus inbox / pesan yang ada pada suatu akun
func hapusPesan(a *tabAkun, akunIndex, pesanIndex int) {
	a[akunIndex].inbox[pesanIndex] = pesan{
		id 	: 0,
		sender : "",
		recipient : "",
		subject : "",
		body : "",
		ada : false,
	}
}


//? fungsi untuk mengscan 1 kalimat pada subject
func inputSubject(subject *string){ 
	var karakter byte
	fmt.Scanf("\n%c", &karakter) // \n untuk pertanda bahwa karakter akan dibaca setelah ada enter (pada inputan penerima)

	for karakter != '\n' {
		*subject += string(karakter)
		fmt.Scanf("%c", &karakter)
	}
} 

//? fungsi untuk mengscan 1 kalimat pada pesan body 
func inputPesan(body *string){ 
	var karakter byte
	fmt.Scanf("%c", &karakter) // tidak perlu \n karena pada inputSubject sudah diakhiri \n dan langsung membaca karakter
	
	for karakter != '\n' {
		*body += string(karakter)
		fmt.Scanf("%c", &karakter)
	}
}


// ? fungsi untuk memberikan jeda waktu selama x milidetik
func jedaWaktu(ms int) {	
	time.Sleep(time.Duration(ms) * time.Millisecond)
}


//! ===============================	FUNGSI DILUAR ALGORITMA ===========================================================

// ? fungsi untuk membersihkan terminal
func cls() {  
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//? fungsi tampilan awal program
func welcome() {
	// Animasi sambutan
	cls()
	fmt.Println("╬════════════════════════╬")
	jedaWaktu(70)
	fmt.Println("║" + red +"███" + reset +"▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒" + red +"███" + reset +"║")
	jedaWaktu(70)
	fmt.Println("║░░" + red +"███" + reset +"▒▒▒▒▒▒▒▒▒▒▒▒▒▒" + red +"███" + reset +"░░║")
	jedaWaktu(70)
	fmt.Println("║░░░░" + red +"███" + reset +"▒▒▒▒▒▒▒▒▒▒" + red +"███" + reset +"░░░░║")
	jedaWaktu(70)
	fmt.Println("║░░░░░░" + red +"███" + reset +"▒▒▒▒▒▒" + red +"███" + reset +"░░░░░░║")
	jedaWaktu(70)
	fmt.Println("║░░░░░░░" + red +"███" + reset +"▒▒▒" + red +"███" + reset +"░░░░░░░░║")
	jedaWaktu(70)
	fmt.Println("║░░░░░░░░░" + red +"█████" + reset +"░░░░░░░░░░║")
	jedaWaktu(70)
	fmt.Println("║░░░░░░░░░░░░░░░░░░░░░░░░║")
	jedaWaktu(70)
	fmt.Println("║░░░░░░░░░░░░░░░░░░░░░░░░║")
	jedaWaktu(70)
	fmt.Println("║░░░░░░░░░░░░░░░░░░░░░░░░║")
	jedaWaktu(70)
	fmt.Println("╬════════════════════════╬"+reset)
	jedaWaktu(70)
	fmt.Println(boldYellow+ "               ╔╗          ")
	jedaWaktu(70)
	fmt.Println("       ╔═╦══╦═╗╠╬╗")
	jedaWaktu(70)
	fmt.Println("       ║╩╣║║║╬╚╣║╚╗")
	jedaWaktu(70)
	fmt.Println("       ╚═╩╩╩╩══╩╩═╝" + reset)
	fmt.Println()
	fmt.Println("	 Loading ")
	for j := 0; j < 27; j++ {
		if j < 6{
			fmt.Print("█")
			jedaWaktu(60) // Jeda waktu 1 detik
		} else if j >= 6 && j < 11{
			fmt.Print("█")
			jedaWaktu(200) // Jeda waktu 1 detik
		}else {
			fmt.Print("█")
			jedaWaktu(20) // Jeda waktu 1 detik
		}
	}
}

//? fungsi tampilan penutup program
func goodbye() {
	cls()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(boldYellow+ "               ╔╗          ")
	jedaWaktu(70)
	fmt.Println("       ╔═╦══╦═╗╠╬╗")
	jedaWaktu(70)
	fmt.Println("       ║╩╣║║║╬╚╣║╚╗")
	jedaWaktu(70)
	fmt.Println("       ╚═╩╩╩╩══╩╩═╝" + reset)
	fmt.Println()
	fmt.Println("         closing")
	for j := 0; j < 27; j++ {
		if j < 6{
			fmt.Print("█")
			jedaWaktu(60) // Jeda waktu 1 detik
		} else if j >= 6 && j < 11{
			fmt.Print("█")
			jedaWaktu(200) // Jeda waktu 1 detik
		}else {
			fmt.Print("█")
			jedaWaktu(20) // Jeda waktu 1 detik
		}
	}
	cls()
	thanks()
	jedaWaktu(2000)
	cls()
}

//? ====== fungsi header ===========

func headerLogin() {
fmt.Println(boldYellow +
`
     ╔╗      ╔══╗
     ║║      ╚╣╠╝
     ║║╔══╦══╗║║╔═╗
     ║║║╔╗║╔╗║║║║╔╗╗
     ║╚╣╚╝║╚╝╠╣╠╣║║║
     ╚═╩══╩═╗╠══╩╝╚╝
          ╔═╝║
          ╚══╝
` + reset)

	fmt.Println()
}

func headerSignin() {
fmt.Println(boldYellow +
` ____  __  ___  __ _    __  __ _ 
/ ___)(  )/ __)(  ( \  (  )(  ( \
\___ \ )(( (_ \/    /   )( /    /
(____/(__)\___/\_)__)  (__)\_)__)
` + reset)
}

func headerMainMenu() {
fmt.Println()
fmt.Println(boldYellow + 
`   		╔╗   
	╔═╦══╦═╗╠╬╗
	║╩╣║║║╬╚╣║╚╗
	╚═╩╩╩╩══╩╩═╝
` + reset)
}

func headerSignUp(){
	fmt.Println( boldYellow +
` ____  __  ___  __ _    _  _  ____ 
/ ___)(  )/ __)(  ( \  / )( \(  _ \
\___ \ )(( (_ \/    /  ) \/ ( ) __/
(____/(__)\___/\_)__)  \____/(__)  ` + reset)
fmt.Println()
}

func thanks() {
	fmt.Println(boldYellow +
`


      Good Bye !!



`+ reset)
}