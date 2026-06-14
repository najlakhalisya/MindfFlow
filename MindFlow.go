package main
import ( 
	"fmt"
	"os"		//untuk sambungin program ke sistem
	"os/exec" 	//untuk jalanin perintah terminal langsung dari kode (buat function clearscreen)
	"runtime" 	//untuk deteksi user menggunakan mac/linux atau windows(biar clearscreennya lancar)
	"bufio" 	//buat baca text panjang pake spasi
	"strings" 	//buat bersihin karakter "enter" biar rapih
)

const NMAX = 1000
type Mood struct {
	tanggal, deskripsi string
	skalaEmosi int
}

type TugasHarian struct{
	tanggal, namaTugas string
	durasiPengerjaan, skalaPrioritas int
	status bool
}

type TabMood [NMAX]Mood
type TabTugas [NMAX]TugasHarian
type TabFound [NMAX] int

//biar parameternya ga kepanjangan
type MindFlowData struct{
	mood TabMood
	totalMood int
	tugas TabTugas
	totalTugas int
}

//procedure buat bersihkan layar
func clearScreen() {
	var cmd *exec.Cmd //var untuk menyimpan perintah terminal
	if runtime.GOOS == "windows" { //ini kalau user menggunakan windows
		cmd = exec.Command("cmd", "/c", "cls") //jika true maka akan menjalankan perintah "cls"
	} else { //jika osnya Mac atau linux
		cmd = exec.Command("clear") //maka akan menjalankan perintah "clear"
	}
	cmd.Stdout = os.Stdout //untuk menyambungkan command ke terminal program
	cmd.Run() //ini agar perintah tadi langsung di eksekusi
}

//procedure buat tampilin header alamat tab yang sedang dikunjungi
func Header(title string) { 
	clearScreen()
	fmt.Println(" ╭──────────────────────────────────────────────────────────────────────────╮")
	fmt.Printf(" │ ⋆𐙚 ̊. %-68s│\n", title) 
	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
}

//untuk menampilkan welcoming screen 
func welcomeScreen(){
	var next string
	
	clearScreen()
	fmt.Println()
	fmt.Println("        .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .        ")
	fmt.Println(" ╭──────────────────────────────────────────────────────────────────────────╮")
	fmt.Println(" │                                                                          │")
	fmt.Println(" │                  ✧･ﾟ:* Welcome to MINDFLOW SPACE *:･ﾟ✧                   │")
	fmt.Println(" │                  ~ Your Personal Mood & Task Manager ~                   │")
	fmt.Println(" │                                                                          │")
	fmt.Println(" │ • — ──────────────────────────── ✧ ✧ ✧ ───────────────────────────── — • │")
	fmt.Println(" │                                                                          │")
	fmt.Println(" │                  Feeling overwhelmed? Or ready to grind?                 │")
	fmt.Println(" │                 Take a deep breath. We got this together!                │")
	fmt.Println(" │                                                                          │")
	fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
	fmt.Println("        .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .        ")
	fmt.Println()
	fmt.Print("                  >> Hit [ENTER] to start your journey... ")

	fmt.Scanln(&next) 
}

//procedure untuk keluar aplikasi
func exitScreen(start *bool) {
	var hold string
	clearScreen()
	fmt.Println()
	fmt.Println("        ╭──────────────────────────────────────────────────────────╮")
	fmt.Println("        │                                                          │")
	fmt.Println("        │     ☾ ⋆･ﾟ:⋆*･ﾟ       SIGNING OFF...       *:･ﾟ✧*:･ﾟ✧     │")
	fmt.Println("        │                                                          │")
	fmt.Println("        │           Yahh sudah mau Bye-Bye, okie then...           │")
	fmt.Println("        │       See Ya!! Jangan lupa istirahat yang cukup ♡        │")
	fmt.Println("        │                                                          │")
	fmt.Println("        ╰──────────────────────────────────────────────────────────╯")
	fmt.Println()
	fmt.Print("               >> press [ENTER] to safely close the space... ")
	fmt.Scanln(&hold)
	clearScreen()
	*start = false
}

//procedure untuk menampilkan main menu
func mainMenu(start *bool, dataBase *MindFlowData, reader *bufio.Reader) { 
	var chooseMenu string
	var hold string
	*start = true
	
	Header("Main Menu")
	fmt.Println(" │                        ╭── ⋅ ⋅ ── ✩ ── ⋅ ⋅ ──╮                           │")
	fmt.Println(" │                             MINDFLOW SPACE                               │")
	fmt.Println(" │                        ╰── ⋅ ⋅ ── ✩ ── ⋅ ⋅ ──╯                           │")
	fmt.Println(" │                                                                          │")
	fmt.Println(" │  [1] ✿  Daily Mood         (Catat Perasaanmu Hari Ini)                   │")
	fmt.Println(" │  [2] ⋆  To-Do List         (Yuk, Produktif & Beresin Tugas!)             │")
	fmt.Println(" │  [3] ⌕  Mind Compass       (Throwback Memori & Spill Statistics)         │")
	fmt.Println(" │  [0] ☾  Signing Off        (Waktunya Istirahat)                          │")
	fmt.Println(" │                                                                          │")
	fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
	fmt.Print("  >> What's your focus today? (Pilih angka) : ")
	fmt.Scanln(&chooseMenu)
	
	if chooseMenu == "1"{
		menuMood(&dataBase.mood, &dataBase.totalMood, reader)
	} else if chooseMenu == "2"{
		menuTugas(&dataBase.tugas, &dataBase.totalTugas, reader)
	} else if chooseMenu == "3"{
		menuTrack(dataBase, reader)
	} else if chooseMenu == "0"{
		exitScreen(start)
	} else {
		fmt.Println()
		fmt.Print(" Oops kamu typo tau (,,•᷄‎ࡇ•᷅ ,,)? , Try Again okey! Tekan [ENTER]...")
		fmt.Scanln(&hold)
	}
}


//procedure untuk menampilkan Sub-menu Mood
func menuMood(mood *TabMood, totalMood *int, reader *bufio.Reader) {
	var hold string
	var balik bool 
	var pilih string

	for balik == false {
		Header("Main Menu > Daily Mood")
		fmt.Println(" │                         ╭── ⋅ ⋅ ── ✩ ── ⋅ ⋅ ──╮                          │")
		fmt.Println(" │                              DAILY MOOD HUB                              │")
		fmt.Println(" │                         ╰── ⋅ ⋅ ── ✩ ── ⋅ ⋅ ──╯                          │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" │   [1] ✿  Log My Vibe     (Cerita Hari Ini)                               │")
		fmt.Println(" │   [2] ⋆  Mood Gallery    (Lihat Riwayat Hati)                            │")
		fmt.Println(" │   [0] ⨯  Back to Home    (Kembali)                                       │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Print("   >> what's your choice? : ")
		fmt.Scanln(&pilih)

		if pilih == "1" {
			addMood(mood, totalMood, reader)
		} else if pilih == "2" {
			viewMood(mood, totalMood, reader)
		} else if pilih == "0" {
			balik = true
		} else {
			fmt.Print("\n   [!] Pilihan tidak tersedia... [ENTER]")
			fmt.Scanln(&hold)
		}
	}
}

//prosedur menambahkan data mood user
func addMood(mood *TabMood, totalMood *int, reader *bufio.Reader) {
	var skala int
	var date, story, hold, hari, bulan string
	var isValid, cekTanggal bool // isValid Untuk membantu pengecekan loop

	Header("Main Menu > Daily Mood > Add Mood")
	fmt.Println(" │  ( 1 / 3 )  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .  │")
	fmt.Println(" │  ⋆  When did this happen?                                                │")
	fmt.Println(" │      ( let's mark the calendar.. ) DD/MM/YY                              │")
	
	// logika untuk memvalidasi tanggal agar user tidak salah menginputkan tanggal
	isValid = false
	cekTanggal = false
	for !isValid {
		fmt.Print(" │  >> Date  : ")
		_, err := fmt.Scanln(&date)
		if err != nil { //jika error-nya TIDAK kosong aka terdapat error
			reader.ReadString('\n') // Bersihkan buffer jika terjadi error input
		}
    
		if len(date) == 8 && date[2] == '/' && date[5] == '/' &&
			date[0:2] >= "01" && date[0:2] <= "31" &&
			date[3:5] >= "01" && date[3:5] <= "12" &&
			date[6:8] >= "00" && date[6:8] <= "99" {
			isValid = true
			fmt.Printf("\033[1A\033[76C│\n")
			
			//cek tanggal 29, 30 dan 31  di setiap bulan
			hari = date[0:2]
			bulan = date[3:5]
			
			if bulan == "02" && hari <= "29"{
				cekTanggal = true
			} else if (bulan == "04" || bulan == "06" || bulan == "09" || bulan == "11") && hari <= "30" {
				cekTanggal = true
			} else if ( bulan == "01" || bulan == "03" || bulan == "05" || bulan == "07" || bulan == "08" || bulan == "10" || bulan == "12"){
				cekTanggal = true
			}
		}
		if cekTanggal {
			isValid = true
			fmt.Printf("\033[1A\033[76C│\n")
		} else {
			fmt.Printf("\033[1A\033[76C│\n")
			fmt.Println(" │  ( ! ) Format salah! Gunakan DD/MM/YY (Contoh: 20/05/26)                 │")
			fmt.Print(" │  >> Press [ENTER] to try again... ")
			fmt.Scanln(&hold) 
			
			fmt.Print("\033[1A\033[2K") // Naik & hapus baris "press enter"
			fmt.Print("\033[1A\033[2K") // Naik & hapus baris "format salah"
			fmt.Print("\033[1A\033[2K") // Naik & hapus baris "input user" (>> Date : ...)
		}
	}

	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
	fmt.Println(" │  ( 2 / 3 )  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .  │")
	fmt.Println(" │  ⋆  How's your heart feeling? (1-5)                                      │")
	fmt.Println(" │      ( 1: Feeling Sad  /  5: Super Happy )                               │")
	
	// logika validasi tanggal
	isValid = false
	for !isValid {
		fmt.Printf(" │  >> Scale : ")
		_, err := fmt.Scanln(&skala)
		if err != nil {
			reader.ReadString('\n') // Bersihkan buffer jika user menginput huruf
		}
    
		if skala >= 1 && skala <= 5 {
			isValid = true
			fmt.Printf("\033[1A\033[76C│\n")
		} else {
			fmt.Printf("\033[1A\033[76C│\n")
			fmt.Println(" │  ( ! ) Skala harus angka 1 sampai 5!                                     │")
			fmt.Print(" │  >> Press [ENTER] to try again... ")
        
			fmt.Scanln(&hold)

			fmt.Print("\033[1A\033[2K") // Naik & hapus baris "press enter"
			fmt.Print("\033[1A\033[2K") // Naik & hapus baris "format salah"
			fmt.Print("\033[1A\033[2K") // Naik & hapus baris "input user" (>> Scale : ...)
		}
	}
	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
	fmt.Println(" │  ( 3 / 3 )  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .  │")
	fmt.Println(" │  ⋆  Spill the tea! What's on your mind?                                  │")
	fmt.Println(" │      ( tell me a little short story.. )                                  │")
	fmt.Print(" │  >> Story : ")
	story, _ = reader.ReadString('\n')
	story = strings.TrimSpace(story)
	fmt.Printf("\033[1A\033[76C│\n")
	fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")

	if *totalMood < NMAX {
		mood[*totalMood].tanggal = date
		mood[*totalMood].skalaEmosi = skala
		mood[*totalMood].deskripsi = story
		*totalMood++

		fmt.Println("\n                        ⭒ ─ ⋅ ⋅ ─── ✩ ─── ⋅ ⋅ ─ ⭒              ")
		fmt.Println("                          noted! your vibe is safe             ")
		fmt.Println("                            inside our heart.                  ")
		fmt.Println("                        ⭒ ─ ⋅ ⋅ ─── ✩ ─── ⋅ ⋅ ─ ⭒              ")
	} else {
		fmt.Println("\n                   ( ! )  storage is full, dear..         ")
	}

	fmt.Print("\n                     >> press [ENTER] to go back home ")
	fmt.Scanln(&hold)
}

//procedure untuk tampilkan Gallery Mood
func displayGallery(mood TabMood, totalMood int) {
	var i int
	if totalMood == 0 {
		fmt.Println(" │                                                                          │")
		fmt.Println(" │        ( ! ) gallery is empty... yuk isi dulu.                           │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
	} else {
		fmt.Println(" │             . ˚ ₊  ✧  YOUR PERSONAL MOOD JOURNEY  ✧  ₊ ˚ .               │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
		fmt.Println(" │ ⋆ NO ⋆ │ ✦ DATE ✦ │ ⋆ SCALE ⋆ │           ✦ SHORT DESCRIPTION ✦          │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
		for i = 0; i < totalMood; i++ {
			fmt.Printf(" │ ˚₊ [%-2d]  %-8s  ♡ mood %-1d/5   ˖ %-38.38s │\n",
				i+1, mood[i].tanggal, mood[i].skalaEmosi, mood[i].deskripsi)
		}
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
	}
}

//procedure untuk menghapus data sesuai index yang dipilih
func deleteMood(mood *TabMood, totalMood *int) {
	var num, idx, j int
	var isValid bool
	var hold, confirm string

	if *totalMood == 0 {
		fmt.Println("\n   ⋆.ೃ࿔* journal masih kosong, belum ada yang bisa dihapus...")
		fmt.Scanln()
	} else {
		isValid = false
		for !isValid {
			fmt.Print("\n   ╰┈➤ nomor memori yang mau dihapus : ")
			_, err := fmt.Scanln(&num)
			if err != nil {
				// Bersihkan buffer langsung menggunakan os.Stdin jika user memasukkan huruf
				bufio.NewReader(os.Stdin).ReadString('\n')
			}

			if num > 0 && num <= *totalMood {
				fmt.Print("   ╰┈➤ kamu yakin mau hapus mission ini? \n        ‧₊˚ yes ❲ 1 ❳ / no ❲ 2 ❳ ˚₊‧ : ")
				fmt.Scanln(&confirm)
				
				if confirm == "1" {
					isValid = true
					idx = num - 1
					for j = idx; j < *totalMood-1; j++ {
						mood[j] = mood[j+1]
				}
					*totalMood = *totalMood - 1
					fmt.Print("\n   ˗ˏˋ poof! Journalmu udah berhasil dihapus ˎˊ˗")
					fmt.Scanln()
				} else {
					isValid = true
					fmt.Print("\n   ˗ˏˋ fiuh! mission-nya aman, gak jadi dihapus ˎˊ˗")
					fmt.Scanln()
				}
			} else {
				fmt.Println("   ₊˚. oops, memori tidak ditemukan ˚.₊")
				fmt.Print("   >> Press [ENTER] to try again... ")
				fmt.Scanln(&hold)
				
				// Membersihkan baris error agar tampilan tetap rapi saat looping
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
			}
		}
	}
}

//procedure buat nulis ulang atau perbaiki data yang sudah ada
func rewriteMood(mood *TabMood, totalMood int, reader *bufio.Reader){
	var num, idx int
	var hold, newStory, hari, bulan string
	var isValid, isFound, cekTanggal bool
	
	if totalMood == 0{
		fmt.Println("\n   ⋆.ೃ࿔* gallery is empty, ga ada yang bisa ditulis ulang...")
		fmt.Scanln()
	} else {
		isFound = false
		for !isFound {
			fmt.Print("   ╰┈➤ mana nih yang mau di-rewrite? : ")
			_, err := fmt.Scanln(&num)
			if err != nil {
				reader.ReadString('\n') // Bersihkan buffer jika terjadi error input
			}
			
			if num > 0 && num <= totalMood{
				isFound = true
				idx = num - 1
				fmt.Println("\n   ╭── ⋅ ⋅ ─── ✩ REWRITE MOMENT ✩ ─── ⋅ ⋅ ──╮")
				fmt.Println("   │    let's fix this memory, shall we?    │")
				fmt.Println("   ╰────────────────────────────────────────╯")
				isValid = false
				for !isValid {
					cekTanggal = false
					fmt.Print("   ˚₊· update tanggal-nya dong (DD/MM/YY) : ")
					_, err := fmt.Scanln(&mood[idx].tanggal)
					if err != nil {
						reader.ReadString('\n') 
					}
					
					if len(mood[idx].tanggal) == 8 && mood[idx].tanggal[2] == '/' && mood[idx].tanggal[5] == '/' &&
						mood[idx].tanggal[0:2] >= "01" && mood[idx].tanggal[0:2] <= "31" &&
						mood[idx].tanggal[3:5] >= "01" && mood[idx].tanggal[3:5] <= "12" &&
						mood[idx].tanggal[6:8] >= "00" && mood[idx].tanggal[6:8] <= "99" {
						
						hari = mood[idx].tanggal[0:2]
						bulan = mood[idx].tanggal[3:5]
						
						if bulan == "02" && hari <= "29" {
							cekTanggal = true
						} else if (bulan == "04" || bulan == "06" || bulan == "09" || bulan == "11") && hari <= "30" {
							cekTanggal = true
						} else if bulan == "01" || bulan == "03" || bulan == "05" || bulan == "07" || bulan == "08" || bulan == "10" || bulan == "12" {
							cekTanggal = true
						}
					}
					
					if cekTanggal {
						isValid = true
					} else {
						fmt.Println("   ( ! ) Format salah! Gunakan DD/MM/YY (Contoh: 20/05/26)")
						fmt.Print("   >> Press [ENTER] to try again... ")
						fmt.Scanln(&hold)
						fmt.Print("\033[1A\033[2K")
						fmt.Print("\033[1A\033[2K") 
						fmt.Print("\033[1A\033[2K") 
					}
				}
				isValid = false
				for !isValid {
					fmt.Print("   ˚₊· how's your mood now? (scale 1-5) : ")
					_, err := fmt.Scanln(&mood[idx].skalaEmosi)
					if err != nil {
						reader.ReadString('\n') 
						mood[idx].skalaEmosi = 0 // reset nilai ke 0 kalau user masukin huruf
					}

					if mood[idx].skalaEmosi >= 1 && mood[idx].skalaEmosi <= 5 {
						isValid = true
					} else {
						fmt.Println("   ( ! ) Skala harus angka 1 sampai 5!")
						fmt.Print("   >> Press [ENTER] to try again... ")
						fmt.Scanln(&hold)
						fmt.Print("\033[1A\033[2K")
						fmt.Print("\033[1A\033[2K") 
						fmt.Print("\033[1A\033[2K")
					}
				}	
				fmt.Print("   ˚₊· spill the new story : ")
				newStory, _ = reader.ReadString('\n')
				mood[idx].deskripsi = strings.TrimSpace(newStory)
				fmt.Println("\n   ˗ˏˋ ♡ phew! memory-nya udah berhasil di-update! ˎˊ˗")
				fmt.Print("   >> hit [ENTER] to save your new mood... ")
				fmt.Scanln(&hold)
			} else {
				fmt.Println("\n   ₊˚. oops, nomornya nggak ketemu. coba cek lagi deh ˚.₊")
				fmt.Print("   >> Press [ENTER] to try again... ")
				fmt.Scanln(&hold)
				
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
			}
		}
	}
}

//prosedur untuk melihat riwayat input data, menghapus data, serta memodifikasi data
func viewMood(mood *TabMood, totalMood *int, reader *bufio.Reader) {
	var pilih string
	var nextmenu bool
	var hold string
	nextmenu = true

	for nextmenu {
		Header("Main Menu > Daily Mood > Mood Gallery")
		displayGallery(*mood, *totalMood)
		
		fmt.Println("\n ˗ˏˋ  [1] new mood  ┊  [2] rewrite  ┊  [3] erase  ┊  [4] sort  ┊  [0] back  ˎˊ˗")
		fmt.Print("\n   ╰┈➤  what's next? : ")
		
		fmt.Scanln(&pilih)

		if pilih == "0" {
			nextmenu = false
		} else if pilih == "1" {
			addMood(mood, totalMood, reader)
		} else if pilih == "2" {
			rewriteMood(mood, *totalMood, reader)
		} else if pilih == "3" {
			deleteMood(mood, totalMood)
		} else if pilih == "4" { 
			menuSortMood (mood, totalMood)
		} else {
			fmt.Print("\n   ⋆⭒˚ maybe there's a tiny typo in your choice ˚⭒⋆  [ENTER]")
			fmt.Scanln(&hold)
		}
	}
}


//procedure menu tugas to do list
func menuTugas(tugas *TabTugas, totalTugas *int, reader *bufio.Reader) {
	var hold string
	var balik bool 
	var pilih string
	
	balik = false

	for balik == false {
		Header("Main Menu > To-Do List")
		fmt.Println(" │                          ╭── ⋅ ⋅ ── ✩ ── ⋅ ⋅ ──╮                         │")
		fmt.Println(" │                              MISSION CONTROL                             │")
		fmt.Println(" │                          ╰── ⋅ ⋅ ── ✩ ── ⋅ ⋅ ──╯                         │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" │   [1] ✿  Add Mission    (Catat Tugas Baru)                               │")
		fmt.Println(" │   [2] ⋆  Mission Log    (Lihat & Kelola Tugas)                           │")
		fmt.Println(" │   [0] ⨯  Back to Home   (Kembali)                                        │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Println()
		fmt.Print("   >> what's your choice? : ")
		fmt.Scanln(&pilih)

		if pilih == "1" {
			addTask(tugas, totalTugas, reader)
		} else if pilih == "2" {
			viewTask(tugas, totalTugas, reader)
		} else if pilih == "0" {
			balik = true
		} else {
			fmt.Print("\n   [!] Pilihan tidak tersedia... [ENTER]")
			fmt.Scanln(&hold)
		}
	}
}

//procedure untuk meng input data tugas
func addTask(tugas *TabTugas, totalTugas *int, reader *bufio.Reader) {
	var prioritas, durasi int
	var date, nama, hold, hari, bulan string
	var isValid, cekTanggal bool

	Header("Main Menu > To-Do List > Add Mission")
	fmt.Println(" │  ( 1 / 4 )  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .  │")
	fmt.Println(" │  ⋆  When is the deadline?                                                │")
	fmt.Println(" │      ( let's mark the calendar.. ) DD/MM/YY                              │")
	isValid = false
	for !isValid {
		cekTanggal = false
		fmt.Print(" │  >> Date  : ")
		_, err := fmt.Scanln(&date)
		if err != nil {
			reader.ReadString('\n') 
		}

		if len(date) == 8 && date[2] == '/' && date[5] == '/' &&
			date[0:2] >= "01" && date[0:2] <= "31" &&
			date[3:5] >= "01" && date[3:5] <= "12" &&
			date[6:8] >= "00" && date[6:8] <= "99" {
			
			hari = date[0:2]
			bulan = date[3:5]
			
			if bulan == "02" && hari <= "29" {
				cekTanggal = true
			} else if (bulan == "04" || bulan == "06" || bulan == "09" || bulan == "11") && hari <= "30" {
				cekTanggal = true
			} else if bulan == "01" || bulan == "03" || bulan == "05" || bulan == "07" || bulan == "08" || bulan == "10" || bulan == "12" {
				cekTanggal = true
			}
		}

		if cekTanggal {
			isValid = true
			fmt.Printf("\033[1A\033[76C│\n")
		} else {
			fmt.Printf("\033[1A\033[76C│\n")
			fmt.Println(" │  ( ! ) Format salah! Gunakan DD/MM/YY (Contoh: 20/05/26)                 │")
			fmt.Print(" │  >> Press [ENTER] to try again... ")
			fmt.Scanln(&hold)
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K") 
			fmt.Print("\033[1A\033[2K")
		}
	}


	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
	fmt.Println(" │  ( 2 / 4 )  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .  │")
	fmt.Println(" │  ⋆  What's the mission called?                                           │")
	fmt.Println(" │      ( nama tugas atau matkulnya.. )                                     │")
	fmt.Print(" │  >> Task  : ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	fmt.Printf("\033[1A\033[76C│\n")

	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
	fmt.Println(" │  ( 3 / 4 )  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .  │")
	fmt.Println(" │  ⋆  How long will it take? (in minutes)                                  │")
	fmt.Println(" │      ( estimasi waktu pengerjaan.. )                                     │")
	
	isValid = false
	for !isValid {
		fmt.Print(" │  >> Time  : ")
		_, err := fmt.Scanln(&durasi)
		// Cek apakah inputnya huruf/kosong
		if err != nil {
			reader.ReadString('\n')
			fmt.Printf("\033[1A\033[76C│\n")
			fmt.Println(" │  ( ! ) Format salah! Masukkan durasi dalam bentuk angka.                 │")
			fmt.Print(" │  >> Press [ENTER] to try again... ")
			fmt.Scanln(&hold)
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K") 
			fmt.Print("\033[1A\033[2K")
			
		// Cek kalau durasinya minus atau nol
		} else if durasi <= 0 {
			fmt.Printf("\033[1A\033[76C│\n")
			fmt.Println(" │  ( ! ) Wait, 0 or minus? Nugas pake cheat code ya? Be for real yuk~      │")
			fmt.Print(" │  >> Press [ENTER] to try again... ")
			fmt.Scanln(&hold)
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K") 
			fmt.Print("\033[1A\033[2K")
			
		// Cek kalau durasinya lebih dari 6 jam (360 menit)
		} else if durasi > 360 {
			fmt.Printf("\033[1A\033[76C│\n")
			fmt.Println(" │  ( ! ) Limit 360 mins pls! Take a break before you get overwhelmed~      │")
			fmt.Print(" │  >> Press [ENTER] to try again... ")
			fmt.Scanln(&hold)
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K") 
			fmt.Print("\033[1A\033[2K")
			
		// Kalau inputnya angka bener (1 - 360)
		} else {
			isValid = true
			fmt.Printf("\033[1A\033[76C│\n")
		}
	}

	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
	fmt.Println(" │  ( 4 / 4 )  . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .  │")
	fmt.Println(" │  ⋆  How urgent is this? (1-5)                                            │")
	fmt.Println(" │      ( 1: Santai banget  /  5: SUPER URGENT! )                           │")
	isValid = false
	for !isValid {
		fmt.Printf(" │  >> Scale : ")
		_, err := fmt.Scanln(&prioritas)
		if err != nil {
			reader.ReadString('\n')
		}

		if prioritas >= 1 && prioritas <= 5 {
			isValid = true
			fmt.Printf("\033[1A\033[76C│\n")
		} else {
			fmt.Printf("\033[1A\033[76C│\n")
			fmt.Println(" │  ( ! ) Skala harus angka 1 sampai 5!                                     │")
			fmt.Print(" │  >> Press [ENTER] to try again... ")
			fmt.Scanln(&hold)
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K") 
			fmt.Print("\033[1A\033[2K")
		}
	}
	fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")

	if *totalTugas < NMAX {
		tugas[*totalTugas].tanggal = date
		tugas[*totalTugas].namaTugas = nama
		tugas[*totalTugas].durasiPengerjaan = durasi
		tugas[*totalTugas].skalaPrioritas = prioritas
		tugas[*totalTugas].status = false 
		*totalTugas++

		fmt.Println("\n                         ⭒ ─ ⋅ ⋅ ─── ✩ ─── ⋅ ⋅ ─ ⭒             ")
		fmt.Println("                         mission locked!! ready to            ")
		fmt.Println("                          grind and be productive          ")
		fmt.Println("                         ⭒ ─ ⋅ ⋅ ─── ✩ ─── ⋅ ⋅ ─ ⭒             ")
	} else {
		fmt.Println("\n                      ( ! )  mission log is full, dear..     ")
	}

	fmt.Print("\n                         >> press [ENTER] to go back ")
	fmt.Scanln(&hold)
}

//procedure buat nampilin tugasnya
func displayTasks(tugas TabTugas, totalTugas int) {
	var i int
	var statusTugas, nama string

	if totalTugas == 0 {
		fmt.Println(" │                                                                          │")
		fmt.Println(" │        ( ! ) mission log is empty... asik bisa rebahan dulu.             │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
	} else {
		fmt.Println(" │                 . ˚ ₊  ✧  YOUR MISSION LOG  ✧  ₊ ˚ .                     │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
		fmt.Println(" │ ⋆ NO ⋆│ Deadline │ STS │           ✦ TASK ✦           │  DURASI │  PRIO  │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
		for i = 0; i < totalTugas; i++ {
			if tugas[i].status {
				statusTugas = "[✓]" 
			} else {
				statusTugas = "[ ]" 
			}
			
			nama = tugas[i].namaTugas
			if len(nama) > 27 {
				nama = nama[:24] + "..."
			}

			fmt.Printf(" │ ˚₊ [%-2d] %-8s ┊ %-3s ┊ %-28s ┊ %3d min ┊ Prio %1d │\n",
				i+1, tugas[i].tanggal, statusTugas, nama, tugas[i].durasiPengerjaan, tugas[i].skalaPrioritas)
		}
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
	}
}

//procedure untuk menghapus data
func deleteTask(tugas *TabTugas, totalTugas *int) {
	var num, idx, j int
	var isValid bool
	var hold, confirm string
	
	if *totalTugas == 0 {
		fmt.Println("\n   ⋆.ೃ࿔* mission log masih kosong, belum ada yang bisa dihapus...")
		fmt.Scanln()
	} else {
		isValid = false
		for !isValid {
			fmt.Print("\n   ╰┈➤ nomor mission yang mau dihapus : ")
			_, err := fmt.Scanln(&num)
			if err != nil {
				bufio.NewReader(os.Stdin).ReadString('\n')
			}

			if num > 0 && num <= *totalTugas {
				fmt.Print("   ╰┈➤ kamu yakin mau hapus mission ini? \n        ‧₊˚ yes ❲ 1 ❳ / no ❲ 2 ❳ ˚₊‧  : ")
				fmt.Scanln(&confirm)
				
				if confirm == "1" {
					isValid = true
					idx = num - 1
					for j = idx; j < *totalTugas - 1; j++ {
						tugas[j] = tugas[j+1]
					}
					*totalTugas = *totalTugas - 1
					fmt.Print("\n   ˗ˏˋ poof! mission-nya udah berhasil dihapus ˎˊ˗")
					fmt.Scanln()
				} else {
					isValid = true
					fmt.Print("\n   ˗ˏˋ fiuh! mission-nya aman, gak jadi dihapus ˎˊ˗")
					fmt.Scanln()
				}	
			} else {
				fmt.Println("   ₊˚. oops, mission tidak ditemukan ˚.₊")
				fmt.Print("   >> Press [ENTER] to try again... ")
				fmt.Scanln(&hold)
				
				// Membersihkan baris error agar tampilan tetap rapi saat looping
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
			}
		}
	}
}

//procedure untuk mengedit data
func rewriteTask(tugas *TabTugas, totalTugas int, reader *bufio.Reader) {
	var num, idx int
	var hold, newNama, hari, bulan string
	var isValid, isFound, cekTanggal bool
	
	if totalTugas == 0 {
		fmt.Println("\n   ⋆.ೃ࿔* log is empty, ga ada yang bisa ditulis ulang...")
		fmt.Scanln()
	} else {
		isFound = false
		for !isFound {
			fmt.Print("   ╰┈➤  mana nih mission yang mau di-rewrite? : ")
			_, err := fmt.Scanln(&num)
			if err != nil {
				reader.ReadString('\n') 
			}

			if num > 0 && num <= totalTugas {
				isFound = true
				idx = num - 1
				fmt.Println("\n   ╭── ⋅ ⋅ ─── ✩ REWRITE MISSION ✩ ─── ⋅ ⋅ ──╮")
				fmt.Println("   │      let's fix this plan, shall we?     │")
				fmt.Println("   ╰─────────────────────────────────────────╯")
				isValid = false
				for !isValid {
					cekTanggal = false
					fmt.Print("   ˚₊· update tanggal deadline (DD/MM/YY) : ")
					_, err := fmt.Scanln(&tugas[idx].tanggal)
					if err != nil {
						reader.ReadString('\n') 
					}
					
					if len(tugas[idx].tanggal) == 8 && tugas[idx].tanggal[2] == '/' && tugas[idx].tanggal[5] == '/' &&
						tugas[idx].tanggal[0:2] >= "01" && tugas[idx].tanggal[0:2] <= "31" &&
						tugas[idx].tanggal[3:5] >= "01" && tugas[idx].tanggal[3:5] <= "12" &&
						tugas[idx].tanggal[6:8] >= "00" && tugas[idx].tanggal[6:8] <= "99" {
						
						hari = tugas[idx].tanggal[0:2]
						bulan = tugas[idx].tanggal[3:5]
						
						if bulan == "02" && hari <= "29" {
							cekTanggal = true
						} else if (bulan == "04" || bulan == "06" || bulan == "09" || bulan == "11") && hari <= "30" {
							cekTanggal = true
						} else if bulan == "01" || bulan == "03" || bulan == "05" || bulan == "07" || bulan == "08" || bulan == "10" || bulan == "12" {
							cekTanggal = true
						}
					}
					
					if cekTanggal {
						isValid = true
					} else {
						fmt.Println("   ( ! ) Format salah! Gunakan DD/MM/YY (Contoh: 20/05/26)")
						fmt.Print("   >> Press [ENTER] to try again... ")
						fmt.Scanln(&hold)
						fmt.Print("\033[1A\033[2K")
						fmt.Print("\033[1A\033[2K") 
						fmt.Print("\033[1A\033[2K")
					}
				}
				fmt.Print("   ˚₊· update nama mission nya : ")
				newNama, _ = reader.ReadString('\n')
				tugas[idx].namaTugas = strings.TrimSpace(newNama)
				
				isValid = false
				for !isValid {
					fmt.Print("   ˚₊· update durasi pengerjaan (menit) : ")
					_, err := fmt.Scanln(&tugas[idx].durasiPengerjaan)
					if err != nil {
						reader.ReadString('\n')
						fmt.Println("  ( ! ) Format salah! Masukkan durasi dalam bentuk angka.                 ")
						fmt.Print("  >> Press [ENTER] to try again... ")
						fmt.Scanln(&hold)
						fmt.Print("\033[1A\033[2K")
						fmt.Print("\033[1A\033[2K") 
						fmt.Print("\033[1A\033[2K")
			
					// Cek kalau durasinya minus atau nol
					} else if tugas[idx].durasiPengerjaan <= 0 {
						fmt.Println("  ( ! ) Wait, 0 or minus? Nugas pake cheat code ya? Be for real yuk~      ")
						fmt.Print("  >> Press [ENTER] to try again... ")
						fmt.Scanln(&hold)
						fmt.Print("\033[1A\033[2K")
						fmt.Print("\033[1A\033[2K") 
						fmt.Print("\033[1A\033[2K")
			
					// Cek kalau durasinya lebih dari 6 jam (360 menit)
					} else if tugas[idx].durasiPengerjaan > 360 {
						fmt.Println("  ( ! ) Limit 360 mins pls! Take a break before you get overwhelmed~      ")
						fmt.Print("  >> Press [ENTER] to try again... ")
						fmt.Scanln(&hold)
						fmt.Print("\033[1A\033[2K")
						fmt.Print("\033[1A\033[2K") 
						fmt.Print("\033[1A\033[2K")
			
					// Kalau inputnya angka bener (1 - 360)
					} else {
						isValid = true
					}
				}

				isValid = false
				for !isValid {
					fmt.Print("   ˚₊· update skala prioritas (1-5) : ")
					_, err := fmt.Scanln(&tugas[idx].skalaPrioritas)
					if err != nil {
						reader.ReadString('\n') 
						tugas[idx].skalaPrioritas = 0 //reset nilai ke 0 kalau user masukin huruf
					}
					
					if tugas[idx].skalaPrioritas >= 1 && tugas[idx].skalaPrioritas <= 5 {
						isValid = true
					} else {
						fmt.Println("   ( ! ) Skala harus angka 1 sampai 5!")
						fmt.Print("   >> Press [ENTER] to try again... ")
						fmt.Scanln(&hold)
						fmt.Print("\033[1A\033[2K")
						fmt.Print("\033[1A\033[2K") 
						fmt.Print("\033[1A\033[2K")
					}
				}
				
				tugas[idx].status = false 

				fmt.Println("\n   ˗ˏˋ ♡ phew! mission-nya udah berhasil di-update! ˎˊ˗")
				fmt.Print("   >> hit [ENTER] to save your new plan... ")
				fmt.Scanln(&hold)
			} else {
				fmt.Println("\n   ₊˚. oops, nomornya nggak ketemu. coba cek lagi deh ˚.₊")
				fmt.Print("   >> Press [ENTER] to try again... ")
				fmt.Scanln(&hold)
				
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
			}
		}
	}
}

//procedure untuk menandai tugas yang sudah selesai
func markTaskDone(tugas *TabTugas, totalTugas int) {
	var num, idx int
	var hold string
	var isValid bool

	if totalTugas == 0 {
		fmt.Println("\n   ⋆.ೃ࿔* log is empty...")
		fmt.Print("   >> press [ENTER] to go back ")
		fmt.Scanln(&hold)
	} else {
		isValid = false
		for !isValid {
			fmt.Print("   ╰┈➤ nomor mission yang udah kelar? : ")
			_, err := fmt.Scanln(&num)
			if err != nil {
				bufio.NewReader(os.Stdin).ReadString('\n')
			}

			if num > 0 && num <= totalTugas {
				idx = num - 1
				isValid = true 

				if !tugas[idx].status {
					tugas[idx].status = true
					fmt.Println("\n   ╭── ⋅ ⋅ ─── ✩ MISSION ACCOMPLISHED ✩ ─── ⋅ ⋅ ──╮")
					fmt.Println("   │  yay! one step closer to rebahan tenang~     │")
					fmt.Println("   ╰──────────────────────────────────────────────╯")
				} else {
					fmt.Println("\n   ╭──── ⋅ ⋅ ───── ✩ ALREADY DONE ✩ ───── ⋅ ⋅ ────╮")
					fmt.Println("   │  mission ini sudah kamu selesaikan, lho!     │")
					fmt.Println("   │  pilih mission lain yang masih kosong ya.    │")
					fmt.Println("   ╰──────────────────────────────────────────────╯")
				}

				fmt.Print("   >> hit [ENTER] to continue... ")
				fmt.Scanln(&hold)

			} else {
				fmt.Println("\n   ( ! ) oops, nomornya nggak ketemu. coba cek lagi..")
				fmt.Print("            >> Press [ENTER] to try again... ")
				fmt.Scanln(&hold)
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K") 
				fmt.Print("\033[1A\033[2K")
			}
		}
	}
}

//procedure untuk menampilkan, menginput, mengedit, menghapus, dan mark tugas jika sudah selesai
func viewTask(tugas *TabTugas, totalTugas *int, reader *bufio.Reader) {
	var pilih string
	var nextmenu bool
	var hold string
	
	nextmenu = true

	for nextmenu {
		Header("Main Menu > To-Do List > Mission Log")
		displayTasks(*tugas, *totalTugas)

		fmt.Println("\n      ˗ˏˋ [1] new  ┊  [2] rewrite  ┊  [3] erase  ┊  [4] mark done ˎˊ˗\n")
		fmt.Println("                 ˗ˏˋ [5] sort      ┊      [0] back ˎˊ˗")
		fmt.Print("\n   ╰┈➤  what's next? : ")
		fmt.Scanln(&pilih)

		if pilih == "0" {
			nextmenu = false
		} else if pilih == "1" {
			addTask(tugas, totalTugas, reader)
		} else if pilih == "2" {
			rewriteTask(tugas, *totalTugas, reader)
		} else if pilih == "3" {
			deleteTask(tugas, totalTugas)
		} else if pilih == "4" {
			markTaskDone(tugas, *totalTugas)
		} else if pilih == "5" {
			menuSortTask(tugas, totalTugas)
		} else {
			fmt.Print("\n   ⋆⭒˚ maybe there's a tiny typo in your choice ˚⭒⋆  [ENTER]")
			fmt.Scanln(&hold)
		}
	}
}


//procedure untuk menampilkan menu Tracker
func menuTrack(data *MindFlowData, reader *bufio.Reader) {
	var pilih string
	var hold string
	var balik bool
	
	balik = false
	
	for balik == false {
		Header("Main Menu > Mind Compass")
		fmt.Println(" │                         ╭── ⋅ ⋅ ── ✩ ── ⋅ ⋅ ──╮                          │")
		fmt.Println(" │                             COMPASS STATION                              │")
		fmt.Println(" │                         ╰── ⋅ ⋅ ── ✩ ── ⋅ ⋅ ──╯                          │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" │   [1] ⌕  Trace Memory   (Lacak Catatan Mood & Histori Tugas)             │")
		fmt.Println(" │   [2] ⊞  Data Insight   (Pantau Rekap Mood & Produktivitas)              │")
		fmt.Println(" │   [0] ⨯  Back to Home   (Kembali ke Menu Utama)                          │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Println()
		fmt.Print("   >> what's your focus today? : ")
		fmt.Scanln(&pilih)
		
		if pilih == "1" {
			menuSearch(data, reader)
		} else if pilih == "2" {
			menuStatistik(data, reader)
		} else if pilih == "0" {
			balik = true
		} else {
			fmt.Print("\n   [!] Pilihan tidak tersedia... [ENTER]")
			fmt.Scanln(&hold)
		}
	}
}


//procedure untuk menu searching
func menuSearch(dataBase *MindFlowData, reader *bufio.Reader){
	var pilih string
	var hold string
	var balik bool
	
	balik = false

	for balik == false{
		Header("Main Menu > Mind Compass > Trace Memory")
		fmt.Println(" │                         ╭── ⋅ ⋅ ── ⌕ ── ⋅ ⋅ ──╮                          │")
		fmt.Println(" │                              MEMORY SEEKER                               │")
		fmt.Println(" │                         ╰── ⋅ ⋅ ── ⌕ ── ⋅ ⋅ ──╯                          │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" │   [1] ✧  Find Mood      (Look back at how your heart felt lately)        │")
		fmt.Println(" │   [2] ✦  Find Mission   (Review the missions you've been carrying)       │")
		fmt.Println(" │   [0] ⨯  Back           (Close the archives and take a step back)        │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Print("\n   >> which piece of memory do you want to find today? : ")
		fmt.Scanln(&pilih)
		
		if pilih == "1"{
			searchMood (dataBase, reader)
		} else if pilih == "2" {
			searchTask (dataBase, reader)
		} else if pilih == "0" {
			balik = true
		} else {
			fmt.Print("\n [!] Oops kamu typo tau (,,•᷄‎ࡇ•᷅ ,,)? , Try Again okey! Tekan [ENTER]...")
			fmt.Scanln(&hold)
		}
	}
}

//procedure searching bagian menu mood
func searchMood(data *MindFlowData, reader *bufio.Reader){
	var pilih string
	var hold string
	var balik bool
	
	balik = false
	
	for balik == false{
		Header("Main Menu > Mind Compass > Seeker > Mood")
		fmt.Println(" │                         ╭── ⋅ ⋅ ── ✧ ── ⋅ ⋅ ──╮                          │")
		fmt.Println(" │                               MOOD SEEKER                                │")
		fmt.Println(" │                         ╰── ⋅ ⋅ ── ✧ ── ⋅ ⋅ ──╯                          │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" │   [1] ⌕  Time Travel    (Cari di tanggal tertentu)                       │")
		fmt.Println(" │   [2] ⌕  Word Tracker   (Cari lewat kata kunci)                          │")
		fmt.Println(" │   [3] ⌕  Emotion Scale  (Filter skor emosimu)                            │")
		fmt.Println(" │   [0] ⨯  Back           (Kembali ke menu sebelumnya)                     │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Print("\n   >> what memory are we looking for? : ")
		fmt.Scanln(&pilih)
		
		if pilih == "1" {
			searchMoodDate (data.mood, data.totalMood)
		} else if pilih == "2" {
			searchMoodKeyword (data.mood, data.totalMood, reader)
		} else if pilih == "3" {
			searchMoodScale (data.mood, data.totalMood)
		} else if pilih == "0" {
			balik = true
		} else {
			fmt.Print("\n   [!] oops, that's a typo. opsinya tidak valid... [ENTER]")
			fmt.Scanln(&hold)
		}
	}
}

//function logika untuk mencari berdasarkan tanggal (menggunakan binary search)
func moodDateFound(mood TabMood, totalMood int, target string) TabFound {
	var result TabFound
	var left, right, mid, found, count, i int
	var convertTarget, convertMid string
	
	for i = 0; i < NMAX; i++ {
		result[i] = -1
	}
	
	//paksa format target menjadi YYMMDD
	convertTarget = target[6:8] + target[3:5] + target[0:2]
	
	left = 0
	right = totalMood - 1
	found = -1
	
	//binery search (membelah data jadi 2 bagian)
	for left <= right && found == -1 {
		mid = (left + right) / 2
		convertMid = mood[mid].tanggal[6:8] + mood[mid].tanggal[3:5] + mood[mid].tanggal[0:2]
		if convertTarget < convertMid {
			right = mid - 1
		} else if convertTarget > convertMid {
			left = mid + 1
		} else {
			found = mid
		}
	}
	//kalau salah satu ketemu, cek kanan kirinya
	if found != -1 {
		i = found 
		for i >= 0 && mood[i].tanggal == target{ //bakal mundur terus selagi tanggalnya masih sama
			i = i - 1
		}
		i++ //kembalikan nilai i ke kanan 1 langkah biar pas di posisi awal tanggal yang sama
		for i < totalMood && mood[i].tanggal == target { //bakal maju ke kanan buat ngumpulin berapa jumlah tanggal yang sama
			result[count] = i
			count++
			i++
		}
	}
	return result
}

//procedure untuk menampilkan pencarian berdasarkan tanggal 
func searchMoodDate(mood TabMood, totalMood int) {
	var target, hold, hari, bulan string
	var sortedMood TabMood
	var arrIdx TabFound
	var found, i, idx int
	var isValid, cekTanggal bool

	if totalMood == 0 {
		fmt.Println("\n   ⋆.ೃ࿔* Jurnal masih kosong, belum ada memori yang bisa dicari...")
		fmt.Scanln(&hold)
	} else {
		isValid = false
		for !isValid {
			cekTanggal = false
			fmt.Print("\n   ╰┈➤ masukkan tanggal cerita yang ingin dicari (DD/MM/YY) : ")
			_, err := fmt.Scanln(&target)
			if err != nil {
				// Bersihkan buffer jika terjadi error input (menggunakan os.Stdin langsung)
				bufio.NewReader(os.Stdin).ReadString('\n')
			}

			if len(mood[idx].tanggal) == 8 && mood[idx].tanggal[2] == '/' && mood[idx].tanggal[5] == '/' &&
				mood[idx].tanggal[0:2] >= "01" && mood[idx].tanggal[0:2] <= "31" &&
				mood[idx].tanggal[3:5] >= "01" && mood[idx].tanggal[3:5] <= "12" &&
				mood[idx].tanggal[6:8] >= "00" && mood[idx].tanggal[6:8] <= "99" {
	
				hari = mood[idx].tanggal[0:2]
				bulan = mood[idx].tanggal[3:5]
				
				if bulan == "02" && hari <= "29" {
					cekTanggal = true
				} else if (bulan == "04" || bulan == "06" || bulan == "09" || bulan == "11") && hari <= "30" {
					cekTanggal = true
				} else if bulan == "01" || bulan == "03" || bulan == "05" || bulan == "07" || bulan == "08" || bulan == "10" || bulan == "12" {
					cekTanggal = true
				}
			}

			if cekTanggal {
				isValid = true
			} else {
				fmt.Println("   ( ! ) Format salah! Gunakan DD/MM/YY (Contoh: 20/05/26)")
				fmt.Print("   >> Press [ENTER] to try again... ")
				fmt.Scanln(&hold)
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
			}
		}

		sortedMood = sortMood(mood, totalMood, 1, true)
		arrIdx = moodDateFound(sortedMood, totalMood, target)
		
		Header("Main Menu > Mind Compass > Seeker > Mood > Time Travel")
		fmt.Println(" │                 . ˚ ₊  ⌕  TIME TRAVEL RESULT  ⌕  ₊ ˚ .                   │")
		fmt.Println(" │                  —  powered by binary search engine  —                   │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
		
		i = 0
		found = 0
		for i < NMAX && arrIdx[i] != -1 {
			idx = arrIdx[i]
			fmt.Printf(" │ ˚₊ %-8s  ♡ mood %-1d/5  ˖ %-45.43s │\n",
				sortedMood[idx].tanggal, sortedMood[idx].skalaEmosi, sortedMood[idx].deskripsi)
			found++
			i++
		}
		
		if found == 0 {
			fmt.Println(" │                                                                          │")
			fmt.Println(" │          ( ! ) Tanggal tersebut tidak ditemukan di memori.               │")
			fmt.Println(" │                                                                          │")
		}
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Printf("\n                       ✧ . ˚ %d Memory Found in Time ˚ . ✧\n", found)
		fmt.Println("              ────────────────────────────────────────────────────")
		fmt.Print("                       >> press [enter] to return home << ")
		fmt.Scanln(&hold)
	}
}

//function logika untuk mencari kata kunci di dalam kalimat (menggunakan sequential search)
func keyMoodFound (mood TabMood, totalMood int, keyword string) TabFound {
	var result TabFound
	var i, count int
	var deskLow, keyLow string //low = lowercase

	count = 0
	for i = 0; i < NMAX; i++ {
		result[i] = -1
	}
	keyLow = strings.ToLower (keyword) //memaksa agar string menjadi lowercase jadi kalau di scan tetap akan terbaca walaupun diawal kata ada Uppercase
	for i = 0; i < totalMood; i++{
		deskLow = strings.ToLower(mood[i].deskripsi)
		// strings.Contains itu gunanya untuk mencari Kata kunci di dalam kalimat, kalau pake "==" dia bakal nyari exact match aka kalimat yang sama persis
		if strings.Contains(deskLow, keyLow){
			result[count] = i
			count++
		}
	}
	return result
}

//procedure untuk menampilkan pencarian berdasarkan keyword
func searchMoodKeyword (mood TabMood, totalMood int, reader *bufio.Reader){
	var keyword, hold string
	var arrIndxDesk  TabFound
	var found, i, idx int
	found = 0
	
	if totalMood == 0{
		fmt.Println("\n   ⋆.ೃ࿔* Jurnal masih kosong, belum ada memori yang bisa dicari...")
		fmt.Scanln(&hold)
	} else {
		fmt.Print("\n   ╰┈➤ masukkan kata kunci cerita yang ingin dicari : ")
		keyword, _ = reader.ReadString('\n')
		keyword = strings.TrimSpace(keyword)
		
		arrIndxDesk = keyMoodFound (mood, totalMood, keyword)
		
		Header("Main Menu > Mind Compass > Seeker > Mood > Word Tracker")
		fmt.Println(" │              . ˚ ₊  ⌕  SEEKER FOUND THESE MEMORIES  ⌕  ₊ ˚ .             │")
		fmt.Println(" │                 —  powered by sequential search engine  —                │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
		for i < NMAX && arrIndxDesk[i] != -1{
			idx = arrIndxDesk[i]
			fmt.Printf(" │ ˚₊ %-8s  ♡ mood %-1d/5  ˖ %-45.43s │\n",
				mood[idx].tanggal, mood[idx].skalaEmosi, mood[idx].deskripsi)
			
			found++
			i++
		}
		if found == 0 {
			fmt.Println(" │                                                                          │")
			fmt.Println(" │         ( ! ) tidak ditemukan memori dengan kata kunci tersebut.         │")
			fmt.Println(" │                                                                          │")
		}
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Printf("\n                  ✧ . ˚ %d Memory Captured in Keywords ˚ . ✧\n", found)
		fmt.Println("             ────────────────────────────────────────────────────")
		fmt.Print("                    >> press [enter] to return home << ")
		fmt.Scanln(&hold)
	}
}

//function logika untuk mencari skala emosi mood (menggunakan sequential search)
func scaleMoodFound (mood TabMood, totalMood, target int) TabFound{
	var result TabFound
	var i, count int
	
	count = 0
	for i = 0; i < NMAX; i++{
		result[i] = -1
	}
	for i = 0; i < totalMood; i++{
		if mood[i].skalaEmosi == target{
			result[count] = i
			count++
		}
	}
	return result
}

//procedure untuk menampilkan pencarian berdasarkan skala emosi
func searchMoodScale(mood TabMood, totalMood int) {
	var target, found, i, idx int
	var arrIdx TabFound
	var hold string
	var isValid bool

	if totalMood == 0 {
		fmt.Println("\n   ⋆.ೃ࿔* Jurnal masih kosong, belum ada memori yang bisa dicari...")
		fmt.Scanln(&hold)
	} else {
		isValid = false
		for !isValid {
			fmt.Print("\n   ╰┈➤ masukkan skala emosi yang ingin difilter (1 - 5) : ")
			_, err := fmt.Scanln(&target)
			if err != nil {
				// Bersihkan buffer jika terjadi error input (menggunakan os.Stdin langsung)
				bufio.NewReader(os.Stdin).ReadString('\n')
			}

			if target >= 1 && target <= 5 {
				isValid = true
			} else {
				fmt.Println("   ( ! ) Skala harus angka 1 sampai 5!")
				fmt.Print("   >> Press [ENTER] to try again... ")
				fmt.Scanln(&hold)
				
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
			}
		}

		arrIdx = scaleMoodFound(mood, totalMood, target)
		Header("Main Menu > Mind Compass > Seeker > Mood > Emotion Scale")
		fmt.Println(" │                . ˚ ₊  ⌕  EMOTION SCALE RESULTS  ⌕  ₊ ˚ .                 │")
		fmt.Println(" │                —  powered by sequential search engine  —                 │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
		
		found = 0
		i = 0
		for i < NMAX && arrIdx[i] != -1 {
			idx = arrIdx[i]
			fmt.Printf(" │ ˚₊ %-8s  ♡ mood %-1d/5  ˖ %-45.43s │\n",
				mood[idx].tanggal, mood[idx].skalaEmosi, mood[idx].deskripsi)
			found++
			i++
		}

		if found == 0 {
			fmt.Println(" │                                                                          │")
			fmt.Println(" │          ( ! ) Tidak ada memori dengan skala emosi tersebut.               │")
			fmt.Println(" │                                                                          │")
		}
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Printf("\n                       ✧ . ˚ %d Memory Matched by Scale ˚ . ✧\n", found)
		fmt.Println("              ────────────────────────────────────────────────────")
		fmt.Print("                        >> press [enter] to return home << ")
		fmt.Scanln(&hold)
	}
}


//procedure searching bagian menu task
func searchTask(data *MindFlowData, reader *bufio.Reader) {
	var pilih string
	var hold string
	var balik bool 
	
	balik = false
	
	for balik == false {
		Header("Main Menu > Mind Compass > Seeker > Mission")
		fmt.Println(" │                         ╭── ⋅ ⋅ ── ✦ ── ⋅ ⋅ ──╮                          │")
		fmt.Println(" │                             MISSION SEEKER                               │")
		fmt.Println(" │                         ╰── ⋅ ⋅ ── ✦ ── ⋅ ⋅ ──╯                          │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" │   [1] ⌕  Deadline Date  (Cari tanggal deadline)                          │")
		fmt.Println(" │   [2] ⌕  Course Name    (Cari spesifik nama tugas)                       │")
		fmt.Println(" │   [3] ⌕  Time Needed    (Filter durasi pengerjaan)                       │")
		fmt.Println(" │   [4] ⌕  Panic Level    (Cek skala prioritas tugas)                      │")
		fmt.Println(" │   [5] ⌕  Progress Check (Lihat status penyelesaian)                      │")
		fmt.Println(" │   [0] ⨯  Back           (Kembali ke menu sebelumnya)                     │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Print("\n   >> what mission are we tracking today? : ")
		fmt.Scanln(&pilih)
		
		if pilih == "1" {
			searchTaskDate(data.tugas, data.totalTugas)
		} else if pilih == "2" {
			searchTaskKeyword(data.tugas, data.totalTugas, reader)
		} else if pilih == "3" {
			searchTaskDuration(data.tugas, data.totalTugas)
		} else if pilih == "4" {
			searchTaskPriority(data.tugas, data.totalTugas)
		} else if pilih == "5" {
			searchTaskStatus(data.tugas, data.totalTugas)
		} else if pilih == "0" {
			balik = true
		} else {
			fmt.Print("\n   [!] oops, that's a typo. opsinya tidak valid... [ENTER]")
			fmt.Scanln(&hold)
		}
	}
}

//function logika untuk mencari berdasarkan tanggal (menggunakan binary search)
func taskDateFound(tugas TabTugas, totalTugas int, target string) TabFound {
	var result TabFound
	var left, right, mid, found, count, i int
	var convertTarget, convertMid string

	for i = 0; i < NMAX; i++ {
		result[i] = -1
	}
	left = 0
	right = totalTugas - 1
	found = -1
	count = 0

	convertTarget = target[6:8] + target[3:5] + target[0:2]

	for left <= right && found == -1 {
		mid = (left + right) / 2
		convertMid = tugas[mid].tanggal[6:8] + tugas[mid].tanggal[3:5] + tugas[mid].tanggal[0:2]

		if convertTarget < convertMid {
			right = mid - 1
		} else if convertTarget > convertMid {
			left = mid + 1
		} else {
			found = mid
		}
	}
	if found != -1 {
		i = found
		for i >= 0 && tugas[i].tanggal == target {
			i--
		}
		i++
		for i < totalTugas && tugas[i].tanggal == target {
			result[count] = i
			count++
			i++
		}
	}
	return result
}

//procedure untuk menampilkan pencarian berdasarkan tanggal
func searchTaskDate(tugas TabTugas, totalTugas int) {
	var target, hold, statusTugas, hari, bulan string
	var sortedTask TabTugas
	var arrIdx TabFound
	var found, i, idx int
	var isValid, cekTanggal bool

	if totalTugas == 0 {
		fmt.Println("\n   ⋆.ೃ࿔* Mission log masih kosong, belum ada tugas yang bisa dicari...")
		fmt.Scanln(&hold)
	} else {
		isValid = false
		for !isValid {
			cekTanggal = false
			fmt.Print("\n   ╰┈➤ Masukkan tanggal deadline dicari (DD/MM/YY) : ")
			_, err := fmt.Scanln(&target)
			if err != nil {
				// Bersihkan buffer langsung menggunakan os.Stdin tanpa perlu deklarasi variabel reader di awal
				bufio.NewReader(os.Stdin).ReadString('\n')
			}

			if len(tugas[idx].tanggal) == 8 && tugas[idx].tanggal[2] == '/' && tugas[idx].tanggal[5] == '/' &&
				tugas[idx].tanggal[0:2] >= "01" && tugas[idx].tanggal[0:2] <= "31" &&
				tugas[idx].tanggal[3:5] >= "01" && tugas[idx].tanggal[3:5] <= "12" &&
				tugas[idx].tanggal[6:8] >= "00" && tugas[idx].tanggal[6:8] <= "99" {
	
				hari = tugas[idx].tanggal[0:2]
				bulan = tugas[idx].tanggal[3:5]
				
				if bulan == "02" && hari <= "29" {
					cekTanggal = true
				} else if (bulan == "04" || bulan == "06" || bulan == "09" || bulan == "11") && hari <= "30" {
					cekTanggal = true
				} else if bulan == "01" || bulan == "03" || bulan == "05" || bulan == "07" || bulan == "08" || bulan == "10" || bulan == "12" {
					cekTanggal = true
				}
			}

			if cekTanggal {
				isValid = true
			} else {
				fmt.Println("   ( ! ) Format salah! Gunakan DD/MM/YY (Contoh: 20/05/26)")
				fmt.Print("   >> Press [ENTER] to try again... ")
				fmt.Scanln()
				
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
			}
		}

		sortedTask = sortTasks(tugas, totalTugas, 1, true)
		arrIdx = taskDateFound(sortedTask, totalTugas, target)

		Header("Main Menu > Mind Compass > Seeker > Mission > Deadline Date")
		fmt.Println(" │               . ˚ ₊  ⌕  DEADLINE TRACKER RESULTS  ⌕  ₊ ˚ .               │")
		fmt.Println(" │                  —  powered by binary search engine  —                   │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")

		i = 0
		found = 0
		for i < NMAX && arrIdx[i] != -1 {
			idx = arrIdx[i]
			if sortedTask[idx].status {
				statusTugas = "[✓]"
			} else {
				statusTugas = "[ ]"
			}
			fmt.Printf(" │ ˚₊ %-8s ┊ %-3s ┊ %-33.30s ┊ %3d min ┊ Prio %1d │\n",
				sortedTask[idx].tanggal, statusTugas, sortedTask[idx].namaTugas, sortedTask[idx].durasiPengerjaan, sortedTask[idx].skalaPrioritas)
			found++
			i++
		}

		if found == 0 {
			fmt.Println(" │                                                                          │")
			fmt.Println(" │          ( ! ) Tidak ditemukan tugas dengan deadline tersebut.               │")
			fmt.Println(" │                                                                          │")
		}
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Printf("\n                       ✧ . ˚ %d Mission Found in the Timeline ˚ . ✧\n", found)
		fmt.Println("              ────────────────────────────────────────────────────")
		fmt.Print("                       >> press [enter] to return home << ")
		fmt.Scanln(&hold)
	}
}

//function logika untuk mencari kata kunci di dalam kalimat (menggunakan sequential search)
func keyTaskFound(tugas TabTugas, totalTugas int, keyword string) TabFound {
	var result TabFound
	var i, count int
	var nameLow, keyLow string

	for i = 0; i < NMAX; i++ {
		result[i] = -1
	}
	count = 0
	keyLow = strings.ToLower(keyword)

	for i = 0; i < totalTugas; i++ {
		nameLow = strings.ToLower(tugas[i].namaTugas)
		if strings.Contains(nameLow, keyLow) {
			result[count] = i
			count++
		}
	}
	return result
}

//procedure untuk menampilkan pencarian berdasarkan keyword
func searchTaskKeyword(tugas TabTugas, totalTugas int, reader *bufio.Reader) {
	var keyword, hold, statusTugas string
	var arrIdx TabFound
	var found, i, idx int

	if totalTugas == 0 {
		fmt.Println("\n   ⋆.ೃ࿔* Mission log masih kosong, belum ada tugas yang bisa dicari...")
		fmt.Scanln(&hold)
	} else {
		fmt.Print("\n   ╰┈➤ Masukkan kata kunci nama tugas / matkul : ")
		keyword, _ = reader.ReadString('\n')
		keyword = strings.TrimSpace(keyword)

		arrIdx = keyTaskFound(tugas, totalTugas, keyword)

		Header("Main Menu > Mind Compass > Seeker > Mission > Mission Name")
		fmt.Println(" │                . ˚ ₊  ⌕  COURSE TRACKER RESULTS  ⌕  ₊ ˚ .                │")
		fmt.Println(" │                —  powered by sequential search engine  —                 │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")

		i = 0
		found = 0
		for i < NMAX && arrIdx[i] != -1 {
			idx = arrIdx[i]
			if tugas[idx].status {
				statusTugas = "[✓]"
			} else {
				statusTugas = "[ ]"
			}

			fmt.Printf(" │ ˚₊ %-8s ┊ %-3s ┊ %-33.30s ┊ %3d min ┊ Prio %1d │\n",
				tugas[idx].tanggal, statusTugas, tugas[idx].namaTugas, tugas[idx].durasiPengerjaan, tugas[idx].skalaPrioritas)
			found++
			i++
		}

		if found == 0 {
			fmt.Println(" │                                                                          │")
			fmt.Println(" │         ( ! ) Tidak ditemukan tugas dengan kata kunci tersebut.          │")
			fmt.Println(" │                                                                          │")
		}
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Printf("\n                  ✧ . ˚ %d Mission Found in the Course ˚ . ✧\n", found)
		fmt.Println("             ────────────────────────────────────────────────────")
		fmt.Print("                     >> press [enter] to return home << ")
		fmt.Scanln(&hold)
	}
}

//function logika untuk mencari berdasarkan durasi (menggunakan binary search)
func taskDurationFound(tugas TabTugas, totalTugas int, target int) TabFound {
	var result TabFound
	var left, right, mid, found, count, i int

	for i = 0; i < NMAX; i++ {
		result[i] = -1
	}
	left = 0
	right = totalTugas - 1
	found = -1
	count = 0

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if target < tugas[mid].durasiPengerjaan {
			right = mid - 1
		} else if target > tugas[mid].durasiPengerjaan {
			left = mid + 1
		} else {
			found = mid
		}
	}
	if found != -1 {
		i = found
		for i >= 0 && tugas[i].durasiPengerjaan == target {
			i--
		}
		i++
		for i < totalTugas && tugas[i].durasiPengerjaan == target {
			result[count] = i
			count++
			i++
		}
	}
	return result
}

//procedure untuk menampilkan pencarian berdasarkan Durasi
func searchTaskDuration(tugas TabTugas, totalTugas int) {
	var target, found, i, idx int
	var hold, statusTugas string
	var sortedTask TabTugas
	var arrIdx TabFound
	var isValid bool

	if totalTugas == 0 {
		fmt.Println("\n   ⋆.ೃ࿔* Mission log masih kosong, belum ada tugas yang bisa dicari...")
		fmt.Scanln(&hold)
	} else {
		isValid = false
		for !isValid {
			fmt.Print("\n   ╰┈➤ Cari durasi pengerjaan yang spesifik (dalam menit) : ")
			_, err := fmt.Scanln(&target)
			if err == nil {
				isValid = true
			} else {
				// Bersihkan buffer langsung menggunakan os.Stdin
				bufio.NewReader(os.Stdin).ReadString('\n')
				fmt.Println("   ( ! ) Format salah! Masukkan durasi dalam bentuk angka.")
				fmt.Print("   >> Press [ENTER] to try again... ")
				fmt.Scanln(&hold)
				
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
			}
		}

		sortedTask = sortTasks(tugas, totalTugas, 3, true)
		arrIdx = taskDurationFound(sortedTask, totalTugas, target)

		Header("Main Menu > Mind Compass > Seeker > Mission > Time Needed")
		fmt.Println(" │               . ˚ ₊   ⏱  DURATION TRACKER RESULTS ⏱  ₊ ˚ .               │")
		fmt.Println(" │                   —  powered by binary search engine  —                  │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")

		i = 0
		found = 0
		for i < NMAX && arrIdx[i] != -1 {
			idx = arrIdx[i]
			if sortedTask[idx].status {
				statusTugas = "[✓]"
			} else {
				statusTugas = "[ ]"
			}

			fmt.Printf(" │ ˚₊ %-8s ┊ %-3s ┊ %-33.30s ┊ %3d min ┊ Prio %1d │\n",
				sortedTask[idx].tanggal, statusTugas, sortedTask[idx].namaTugas, sortedTask[idx].durasiPengerjaan, sortedTask[idx].skalaPrioritas)
			found++
			i++
		}

		if found == 0 {
			fmt.Println(" │                                                                          │")
			fmt.Println(" │         ( ! ) Tidak ada tugas dengan durasi menit tersebut.              │")
			fmt.Println(" │                                                                          │")
		}
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Printf("\n                  ✧ . ˚ %d Mission Found in Work Time ˚ . ✧\n", found)
		fmt.Println("             ────────────────────────────────────────────────────")
		fmt.Print("                    >> press [enter] to return home << ")
		fmt.Scanln(&hold)
	}
}

//function logika untuk mencari prioritas tugas (menggunakan sequential search)
func priorityTaskFound(tugas TabTugas, totalTugas, target int) TabFound {
	var result TabFound
	var i, count int

	for i = 0; i < NMAX; i++ {
		result[i] = -1
	}
	count = 0
	for i = 0; i < totalTugas; i++ {
		if tugas[i].skalaPrioritas == target {
			result[count] = i
			count++
		}
	}
	return result
}

//procedure untuk menampilkan pencarian berdasarkan prioritas tugas
func searchTaskPriority(tugas TabTugas, totalTugas int) {
	var target, found, i, idx int
	var hold, statusTugas string
	var arrIdx TabFound
	var isValid bool

	if totalTugas == 0 {
		fmt.Println("\n   ⋆.ೃ࿔* Mission log masih kosong, belum ada tugas yang bisa dicari...")
		fmt.Scanln(&hold)
	} else {
		isValid = false
		for !isValid {
			fmt.Print("\n   ╰┈➤ Masukkan skala prioritas yang ingin difilter (1-5) : ")
			_, err := fmt.Scanln(&target)
			if err != nil {
				// Bersihkan buffer langsung menggunakan os.Stdin
				bufio.NewReader(os.Stdin).ReadString('\n')
			}

			if target >= 1 && target <= 5 {
				isValid = true
			} else {
				fmt.Println("   ( ! ) Skala harus angka 1 sampai 5!")
				fmt.Print("   >> Press [ENTER] to try again... ")
				fmt.Scanln()
				
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
			}
		}

		arrIdx = priorityTaskFound(tugas, totalTugas, target)

		Header("Main Menu > Mind Compass > Seeker > Mission > Panic Level")
		fmt.Println(" │              . ˚ ₊  ⌕  PRIORITY TRACKER RESULTS  ⌕  ₊ ˚ .                │")
		fmt.Println(" │                —  powered by sequential search engine  —                 │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")

		i = 0
		found = 0
		for i < NMAX && arrIdx[i] != -1 {
			idx = arrIdx[i]
			if tugas[idx].status {
				statusTugas = "[✓]"
			} else {
				statusTugas = "[ ]"
			}

			fmt.Printf(" │ ˚₊ %-8s ┊ %-3s ┊ %-33.30s ┊ %3d min ┊ Prio %1d │\n",
				tugas[idx].tanggal, statusTugas, tugas[idx].namaTugas, tugas[idx].durasiPengerjaan, tugas[idx].skalaPrioritas)
			found++
			i++
		}

		if found == 0 {
			fmt.Println(" │                                                                          │")
			fmt.Println(" │          ( ! ) Tidak ada tugas dengan level prioritas tersebut.             │")
			fmt.Println(" │                                                                          │")
		}
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Printf("\n                       ✧ . ˚ %d Mission Captured by Urgency ˚ . ✧\n", found)
		fmt.Println("              ────────────────────────────────────────────────────")
		fmt.Print("                       >> press [enter] to return home << ")
		fmt.Scanln(&hold)
	}
}

//function logika untuk mencari berdasarkan status (menggunakan sequential search)
func statusTaskFound(tugas TabTugas, totalTugas int, target bool) TabFound {
	var result TabFound
	var i, count int

	for i = 0; i < NMAX; i++ {
		result[i] = -1
	}
	count = 0
	for i = 0; i < totalTugas; i++ {
		if tugas[i].status == target {
			result[count] = i
			count++
		}
	}
	return result
}

//procedure untuk menampilkan pencarian berdasarkan status tugas
func searchTaskStatus(tugas TabTugas, totalTugas int) {
	var pilihMenu, found, i, idx int
	var hold, statusTugas string
	var targetStatus bool
	var arrIdx TabFound
	var isValid bool

	if totalTugas == 0 {
		fmt.Println("\n   ⋆.ೃ࿔* Mission log masih kosong, belum ada tugas yang bisa dicari...")
		fmt.Scanln(&hold)
	} else {
		isValid = false
		for !isValid {
			fmt.Println("\n   ╭┈➤ Pilih status tugas yang ingin ditampilkan:")
			fmt.Println("   │                                               ")
			fmt.Println("   │   [1] Mission Accomplished (Sudah Selesai [✓])")
			fmt.Println("   │   [2] Still On Progress    (Belum Selesai [ ])")
			fmt.Println("   │                                               ")
			fmt.Print("   ╰┈➤ Pilihan kamu (1/2) : ")
			_, err := fmt.Scanln(&pilihMenu)

			if err != nil {
				// Bersihkan buffer langsung menggunakan os.Stdin
				bufio.NewReader(os.Stdin).ReadString('\n')
			}

			if pilihMenu == 1 || pilihMenu == 2 {
				isValid = true
				if pilihMenu == 1 {
					targetStatus = true
				} else {
					targetStatus = false
				}
			} else {
				fmt.Println("   ( ! ) Pilihan tidak tersedia! Masukkan angka 1 atau 2.")
				fmt.Print("   >> Press [ENTER] to try again... ")
				fmt.Scanln()
				
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
				fmt.Print("\033[1A\033[2K")
			}
		}

		arrIdx = statusTaskFound(tugas, totalTugas, targetStatus)

		Header("Main Menu > Mind Compass > Seeker > Mission > Progress Check")
		fmt.Println(" │               . ˚ ₊  ⌕  PROGRESS TRACKER RESULTS  ⌕  ₊ ˚ .               │")
		fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")

		i = 0
		found = 0
		for i < NMAX && arrIdx[i] != -1 {
			idx = arrIdx[i]
			if tugas[idx].status {
				statusTugas = "[✓]"
			} else {
				statusTugas = "[ ]"
			}

			fmt.Printf(" │ ˚₊ %-8s ┊ %-3s ┊ %-33.30s ┊ %3d min ┊ Prio %1d │\n",
				tugas[idx].tanggal, statusTugas, tugas[idx].namaTugas, tugas[idx].durasiPengerjaan, tugas[idx].skalaPrioritas)
			found++
			i++
		}

		if found == 0 {
			fmt.Println(" │                                                                          │")
			fmt.Println(" │          ( ! ) Tidak ditemukan tugas dengan kriteria status ini.           │")
			fmt.Println(" │                                                                          │")
		}
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Printf("\n                       ✧ . ˚ %d Mission Matched by Progress ˚ . ✧\n", found)
		fmt.Println("              ────────────────────────────────────────────────────")
		fmt.Print("                        >> press [enter] to return home << ")
		fmt.Scanln(&hold)
	}
}


//function sorting menu mood menggunakan selection sort
func sortMood (mood TabMood, totalMood, sortby int, isAscend bool) TabMood{
	var result TabMood
	var pass, idx, i, nilaiI, nilaiIdx int
	var temp Mood
	var valuei, valueIdx string
	var harusTukar bool
	
	for i = 0; i < totalMood; i++ {
		result[i] = mood[i]
	}
	pass = 1
	for pass <= totalMood-1{
		idx = pass - 1
		i = pass
		
		for i < totalMood{
			harusTukar = false
			if sortby == 1{ //sorting tanggal
				valuei = result[i].tanggal[6:8] + result[i].tanggal[3:5] + result[i].tanggal[0:2]
				valueIdx = result[idx].tanggal[6:8] + result[idx].tanggal[3:5] + result[idx].tanggal[0:2]
				if isAscend {
					harusTukar = valuei < valueIdx
				} else {
					harusTukar = valuei > valueIdx
				}
			} else if sortby == 2{ //sorting skala emosi
				nilaiI = result[i].skalaEmosi
				nilaiIdx = result[idx].skalaEmosi
				if isAscend{
					harusTukar = nilaiI < nilaiIdx
				} else {
					harusTukar = nilaiI > nilaiIdx
				}
			} else if sortby == 3{ //sorting berdasarkan deskripsi
				valuei = strings.ToLower(result[i].deskripsi)
				valueIdx = strings.ToLower(result[idx].deskripsi)
				if isAscend{
					harusTukar = valuei < valueIdx
				} else {
					harusTukar = valuei > valueIdx
				}
			}
			if harusTukar {
				idx = i
			}
			i = i + 1
		}
		temp = result[pass-1]
		result[pass-1] = result[idx]
		result[idx] = temp
		pass = pass + 1
	}
	return result
}

//tampilan menu sorting mood
func menuSortMood(mood *TabMood, totalMood *int) {
	var kategori, urutan string
	var katInt int
	var isAsc, isValidKategori, isValidUrutan bool
	var hold string

	isValidKategori = false
	for !isValidKategori {
		fmt.Println("\n                    ╭──── ⋅ ⋅ ──── ✩ SORT MOOD ✩ ──── ⋅ ⋅ ────╮")
		fmt.Println("                    │         powered by selection sort       │")
		fmt.Println("                    ├─────────────────────────────────────────┤")
		fmt.Println("                    │                                         │")
		fmt.Println("                    │     [1] Date   [2] Scale   [3] Story    │")
		fmt.Println("                    │                                         │")
		fmt.Print("                    ╰┈➤ Let's organize it by (1-3) : ")
		fmt.Scanln(&kategori)

		if kategori == "1" || kategori == "2" || kategori == "3" {
			isValidKategori = true
			
			isValidUrutan = false
			for !isValidUrutan {
				fmt.Println("\n                    ╭───── ⋅ ⋅ ─── ✩ ORDER FLOW ✩ ─── ⋅ ⋅ ─────╮")
				fmt.Println("                    │                                          │")
				fmt.Println("                    │      [1] Ascending   [2] Descending      │")
				fmt.Println("                    │                                          │")
				fmt.Print("                    ╰┈➤ How should we arrange it? (1-2) : ")
				fmt.Scanln(&urutan)

				if urutan == "1" || urutan == "2" {
					isValidUrutan = true
					isAsc = (urutan == "1")
					
					// Terjemahkan string ke int untuk dikirim ke fungsi sortMood
					if kategori == "1" {
						katInt = 1
					} else if kategori == "2" {
						katInt = 2
					} else {
						katInt = 3
					}
					
					*mood = sortMood(*mood, *totalMood, katInt, isAsc)
					fmt.Println("\n   ₊˚.♡ All done! Jurnal kamu udah tersusun rapi (Selection Sort applied)\n")
				} else {
					fmt.Println("\n   [!] Hmm.. sepertinya angka yang kamu masukin ga ada di opsi deh...")
					fmt.Print("   >> Press [ENTER] to try again... ")
					fmt.Scanln()
					
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
				}
			}
		} else {
			fmt.Println("\n   [!] Hmm.. sepertinya angka yang kamu masukin ga ada di opsi deh...")
			fmt.Print("   >> Press [ENTER] to try again... ")
			fmt.Scanln()
			
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
		}
	}
	fmt.Print("   >> Hit [ENTER] to continue... ")
	fmt.Scanln(&hold)
}


//function sorting menu Task Menggunakan Insertion Sort
func sortTasks (tugas TabTugas, totalTugas, sortby int, isAscend bool)TabTugas{
	var result TabTugas
	var pass, i, nilaiTemp, nilaiMin1 int
	var temp TugasHarian
	var valueTemp, valueMin1 string
	var terpenuhi bool
	
	for i = 0; i < totalTugas; i++{
		result[i] = tugas[i]
	}
	pass = 1
	for pass <= totalTugas - 1{
		i = pass
		temp = result[pass]
		terpenuhi = true
		for i > 0 && terpenuhi == true{
			terpenuhi = false
			if sortby == 1{ //sorting berdasarkan tanggal
				valueTemp = temp.tanggal[6:8] + temp.tanggal[3:5] + temp.tanggal[0:2]
				valueMin1 = result[i-1].tanggal[6:8] + result[i-1].tanggal[3:5] + result[i-1].tanggal[0:2]
				if isAscend{
					terpenuhi = valueTemp < valueMin1
				} else {
					terpenuhi = valueTemp > valueMin1
				}
			} else if sortby == 2 { //sorting berdasarkan tugas
				valueTemp = strings.ToLower(temp.namaTugas)
				valueMin1 = strings.ToLower(result[i-1].namaTugas)
				if isAscend {
					terpenuhi = valueTemp < valueMin1
				} else {
					terpenuhi = valueTemp > valueMin1
				}
			} else if sortby == 3{ //sorting berdasarkan diurasi ppengerjaan
				nilaiTemp = temp.durasiPengerjaan
				nilaiMin1 = result[i-1].durasiPengerjaan
				if isAscend {
					terpenuhi = nilaiTemp < nilaiMin1
				} else {
					terpenuhi = nilaiTemp > nilaiMin1
				}
			} else if sortby == 4 { // sorting berdasarkan prioritas tugas
				nilaiTemp = temp.skalaPrioritas
				nilaiMin1 = result[i-1].skalaPrioritas
				if isAscend {
					terpenuhi = nilaiTemp < nilaiMin1
				} else {
					terpenuhi = nilaiTemp > nilaiMin1
				}
			}
			if terpenuhi == true {
				result[i] = result[i-1]
				i--
			}
		}
		result[i] = temp
		pass = pass + 1
	}
	return result
}

//tampilan menu sorting task
func menuSortTask(tugas *TabTugas, totalTugas *int) {
	var kategori, urutan string
	var katInt int
	var isAsc, isValidKategori, isValidUrutan bool
	var hold string

	isValidKategori = false
	for !isValidKategori {
		fmt.Println("\n                    ╭─── ⋅ ⋅ ─── ✩ SORT MISSION ✩ ─── ⋅ ⋅ ───╮")
		fmt.Println("                    │        powered by insertion sort       │")
		fmt.Println("                    ├────────────────────────────────────────┤")
		fmt.Println("                    │                                        │")
		fmt.Println("                    │   [1] Date [2] Name [3] Time [4] Prio  │")
		fmt.Println("                    │                                        │")
		fmt.Print("                    ╰┈➤ Let's organize it by (1-4) : ")
		fmt.Scanln(&kategori)

		if kategori == "1" || kategori == "2" || kategori == "3" || kategori == "4" {
			isValidKategori = true
			
			isValidUrutan = false
			for !isValidUrutan {
				fmt.Println("\n                    ╭──── ⋅ ⋅ ─── ✩ ORDER FLOW ✩ ─── ⋅ ⋅ ────╮")
				fmt.Println("                    │                                        │")
				fmt.Println("                    │     [1] Ascending   [2] Descending     │")
				fmt.Println("                    │                                        │")
				fmt.Print("                    ╰┈➤ How should we arrange it? (1-2) : ")
				fmt.Scanln(&urutan)

				if urutan == "1" || urutan == "2" {
					isValidUrutan = true
					isAsc = (urutan == "1")

					// Terjemahkan string ke int untuk dikirim ke fungsi sortTasks
					if kategori == "1" {
						katInt = 1
					} else if kategori == "2" {
						katInt = 2
					} else if kategori == "3" {
						katInt = 3
					} else {
						katInt = 4
					}

					*tugas = sortTasks(*tugas, *totalTugas, katInt, isAsc)
					fmt.Println("\n   ✦ Check! Mission log kamu udah berbaris rapi (Insertion Sort applied) \n")
				} else {
					fmt.Println("\n   [!] Hmm.. sepertinya angka yang kamu masukin ga ada di opsi deh...")
					fmt.Print("   >> Press [ENTER] to try again... ")
					fmt.Scanln()
					
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
					fmt.Print("\033[1A\033[2K")
				}
			}
		} else {
			fmt.Println("\n   [!] Hmm.. sepertinya angka yang kamu masukin ga ada di opsi deh...")
			fmt.Print("   >> Press [ENTER] to try again... ")
			fmt.Scanln()
			
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
		}
	}
	fmt.Print("   >> Hit [ENTER] to continue... ")
	fmt.Scanln(&hold)
}


//menu statistik data 
func menuStatistik(data *MindFlowData, reader *bufio.Reader){
	var pilih string
	var hold string
	var balik bool
	
	balik = false
	for balik == false {
		Header("Main Menu > Mind Compass > Insight Hub")
		fmt.Println(" │                         ╭── ⋅ ⋅ ── ⊞ ── ⋅ ⋅ ──╮                          │")
		fmt.Println(" │                               INSIGHT HUB                                │")
		fmt.Println(" │                         ╰── ⋅ ⋅ ── ⊞ ── ⋅ ⋅ ──╯                          │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" │   [1] ✿  Mood Analytics (Statistik Suasana Hati)                         │")
		fmt.Println(" │   [2] ✦  Task Analytics (Tingkat Penyelesaian Tugas)                     │")
		fmt.Println(" │   [3] ⌕  Correlation    (Hubungan Mood & Produktivitas)                  │")
		fmt.Println(" │   [0] ⨯  Back           (Kembali)                                        │")
		fmt.Println(" │                                                                          │")
		fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
		fmt.Print("\n   ╰┈➤ Choose your insight (1-3) : ")
		fmt.Scanln(&pilih)
		
		if pilih == "1" {
			statMood(data)
		} else if pilih == "2" {
			statTask(data)
		} else if pilih == "3" {
			viewStatConnect(data)
		} else if pilih == "0" {
			balik = true
		} else {
			fmt.Print("\n   [!] Oops, sepertinya opsinya ga ada... [ENTER]")
			fmt.Scanln(&hold)
		}
	}
}

//function untuk tampilan grafik bar
func grafikBar(percent int) string {
	var bar string
	var i, blocks, totalblocks int
	totalblocks = 20
	blocks = percent / 5 
	
	for i = 0; i < blocks; i++ {
		bar += "■"
	}
	for i = blocks; i < totalblocks; i++ {
        bar += "□"
    }
	return bar
}

//function untuk menampilkan statistik bedasarkan skala mood
func statMood(data *MindFlowData){
	var m1, m2, m3, m4, m5, i int
	var p1, p2, p3, p4, p5 int
	var average float64
	var hold string
	
	for i = 0; i < data.totalMood; i++ {
		if data.mood[i].skalaEmosi == 1{
			m1++
		} else if data.mood[i].skalaEmosi == 2{
			m2++
		} else if data.mood[i].skalaEmosi == 3{
			m3++
		} else if data.mood[i].skalaEmosi == 4{
			m4++
		} else if data.mood[i].skalaEmosi == 5{
			m5++
		}
	}
	
	if data.totalMood > 0{
		average = float64((1*m1) + (2*m2) + (3*m3) + (4*m4) + (5*m5)) / float64(data.totalMood)
		p1 = (m1*100) / data.totalMood
		p2 = (m2*100) / data.totalMood
		p3 = (m3*100) / data.totalMood
		p4 = (m4*100) / data.totalMood
		p5 = (m5*100) / data.totalMood
	}
	
	Header("Main Menu > Mind Compass > Insight Hub > Mood Analytics")
	fmt.Println(" │                 . ˚ ₊  ✿  YOUR MOOD ANALYTICS  ✿  ₊ ˚ .                  │")
	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
	fmt.Printf(" │   [5] Ethereal ✧        : %-3d days ┊ %-20s %3d%%           │\n", m5, grafikBar(p5), p5)
	fmt.Printf(" │   [4] Peaceful ❀        : %-3d days ┊ %-20s %3d%%           │\n", m4, grafikBar(p4), p4)
	fmt.Printf(" │   [3] Neutral ⊹         : %-3d days ┊ %-20s %3d%%           │\n", m3, grafikBar(p3), p3)
	fmt.Printf(" │   [2] Drained ♡         : %-3d days ┊ %-20s %3d%%           │\n", m2, grafikBar(p2), p2)
	fmt.Printf(" │   [1] Burnout ⋆         : %-3d days ┊ %-20s %3d%%           │\n", m1, grafikBar(p1), p1)
	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
	fmt.Printf(" │   Total logs saved      : %-3d memori                                     │\n", data.totalMood)
	fmt.Printf(" │   Average mood score    : %.2f / 5.00                                    │\n", average)
	fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
	fmt.Print("\n   ₊˚.♡ Hit [ENTER] to return... ")
	fmt.Scanln(&hold)
}

//function untuk menampilkan statistik berdasarkan progress tugas
func statTask(data *MindFlowData){
	var done, pending, i int
	var pDone, pPending int
	var hold string
	
	for i = 0; i < data.totalTugas; i++{
		if data.tugas[i].status == true {
			done++
		} else {
			pending++
		}
	}
	if data.totalTugas > 0{
		pDone = (done * 100) / data.totalTugas
		pPending = (pending * 100) / data.totalTugas
	}
	
	Header("Main Menu > Mind Compass > Insight Hub > Task Analytics")
	fmt.Println(" │                . ˚ ₊  ✦   PRODUCTIVITY FLOW   ✦  ₊ ˚ .                   │")
	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
	fmt.Printf(" │   [✓] Done Gracefully ✧    : %-3d tasks ┊ %-20s %3d%%       │\n", done, grafikBar(pDone), pDone)
	fmt.Printf(" │   [ ] Still In Motion ⊹    : %-3d tasks ┊ %-20s %3d%%       │\n", pending, grafikBar(pPending), pPending)
	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
	fmt.Printf(" │   Total tasks logged       : %-3d tasks                                   │\n", data.totalTugas)
	fmt.Printf(" │   Productivity flow rate   : %3d %%                                       │\n", pDone)
	fmt.Printf(" │   ✧ . ˚ [ %-20s ] ˚ . ✧                                   │\n",grafikBar(pDone))
	fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
	fmt.Print("\n   ₊˚.♡ Hit [ENTER] to return... ")
	fmt.Scanln(&hold)
}

//untuk mepermudah perhitungan korelasi
type hasilHitung struct{
	goodMood, netralMood, badMood int
	taskGood, taskNetral, taskBad int
	doneGood, doneNetral, doneBad int
}

//function logika untuk menghitung korelasi antara skala mood dan progress tugas
func hitungKorelasi(data *MindFlowData, startDate, endDate string, hasil *hasilHitung){
	var cvStartDate, cvEndDate, cvMood, cvTask, tglMood string
	var i, j, k, totalSkala, jumlahMood, average int
	var isDupe bool
	
	cvStartDate = startDate[6:8] + startDate[3:5] + startDate[0:2]
	cvEndDate = endDate[6:8] + endDate[3:5] + endDate[0:2]
	
	for j = 0; j < data.totalMood; j++{
		isDupe = false
		for k = 0; k < j; k++{
			if data.mood[j].tanggal == data.mood[k].tanggal { //cek apakah ada tanggal yang sama
				isDupe = true
			}
		}
		if isDupe == false {
			tglMood = data.mood[j].tanggal
			cvMood = tglMood[6:8] + tglMood[3:5] + tglMood[0:2]
			if cvMood >= cvStartDate && cvMood <= cvEndDate {
				totalSkala = data.mood[j].skalaEmosi
				jumlahMood = 1
				for k = j+1; k < data.totalMood; k++ {
					if data.mood[j].tanggal == data.mood[k].tanggal {
						totalSkala = totalSkala + data.mood[k].skalaEmosi
						jumlahMood++
					}
				}
				average = int(float64(totalSkala)/float64(jumlahMood) + 0.5)
				if average >= 4 {
					hasil.goodMood++
				} else if average == 3 {
					hasil.netralMood++
				} else {
					hasil.badMood++
				}
			}
		}
	}
	
	for i = 0; i < data.totalTugas; i++ {
		cvTask = data.tugas[i].tanggal[6:8] + data.tugas[i].tanggal[3:5] + data.tugas[i].tanggal[0:2]
		if cvTask >= cvStartDate && cvTask <= cvEndDate{
			totalSkala = 0
			jumlahMood = 0
			for j = 0; j < data.totalMood; j++{
				if data.tugas[i].tanggal == data.mood[j].tanggal{
					totalSkala = totalSkala + data.mood[j].skalaEmosi
					jumlahMood++
				}
			}
			if jumlahMood > 0{
				average = int(float64(totalSkala) / float64(jumlahMood) + 0.5)
				if average >= 4 {
					hasil.taskGood++
					if data.tugas[i].status {
						hasil.doneGood++
					}
				} else if average == 3 {
					hasil.taskNetral++
					if data.tugas[i].status {
						hasil.doneNetral++
					}
				} else {
					hasil.taskBad++
					if data.tugas[i].status {
						hasil.doneBad++
					}
				}
			}
		}
	}
}

//function untuk menampilkan statistik korelasi data mood dan tugas
func viewStatConnect(data *MindFlowData) {
	var startDate, endDate, hold, hari, bulan string
	var pGood, pNetral, pBad int
	var hasil hasilHitung
	var isValidStart, isValidEnd, cekTanggal bool

	Header("Main Menu > Mind Compass > Insight Hub > Correlation")
	fmt.Println(" │                  . ˚ ₊  ⌕  MOOD & PRODUCTIVITY  ⌕  ₊ ˚ .                 │")
	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
	fmt.Println(" │  Let's check your vibe in a specific timeline!                           │")

	isValidStart = false
	for !isValidStart {
		cekTanggal = false
		fmt.Print(" │  ╰┈➤ Start Date (DD/MM/YY) : ")
		_, err := fmt.Scanln(&startDate)
		fmt.Printf("\033[1A\033[76C│\n")
		if err != nil {
			// Bersihkan buffer langsung menggunakan os.Stdin
			bufio.NewReader(os.Stdin).ReadString('\n')
		}
		
		if len(startDate) == 8 && startDate[2] == '/' && startDate[5] == '/' &&
			startDate[0:2] >= "01" && startDate[0:2] <= "31" &&
			startDate[3:5] >= "01" && startDate[3:5] <= "12" &&
			startDate[6:8] >= "00" && startDate[6:8] <= "99" {
			
			hari = startDate[0:2]
			bulan = startDate[3:5]
			
			if bulan == "02" && hari <= "29" {
				cekTanggal = true
			} else if (bulan == "04" || bulan == "06" || bulan == "09" || bulan == "11") && hari <= "30" {
				cekTanggal = true
			} else if bulan == "01" || bulan == "03" || bulan == "05" || bulan == "07" || bulan == "08" || bulan == "10" || bulan == "12" {
				cekTanggal = true
			}
		}
		
		if cekTanggal {
			isValidStart = true
		} else {
			fmt.Println(" │  ( ! ) Format salah! Gunakan DD/MM/YY                                    |")
			fmt.Print(" │  >> Press [ENTER] to try again... ")
			fmt.Scanln()
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
		}
	}

	isValidEnd = false
	for !isValidEnd {
		fmt.Print(" │  ╰┈➤ End Date   (DD/MM/YY) : ")
		_, err := fmt.Scanln(&endDate)
		fmt.Printf("\033[1A\033[76C│\n")
		if err != nil {
			// Bersihkan buffer langsung menggunakan os.Stdin
			bufio.NewReader(os.Stdin).ReadString('\n')
		}
		
		if len(endDate) == 8 && endDate[2] == '/' && endDate[5] == '/' &&
			endDate[0:2] >= "01" && endDate[0:2] <= "31" &&
			endDate[3:5] >= "01" && endDate[3:5] <= "12" &&
			endDate[6:8] >= "00" && endDate[6:8] <= "99" {
			
			isValidEnd = true
		} else {
			fmt.Println(" │  ( ! ) Format salah! Gunakan DD/MM/YY                                    |")
			fmt.Print(" │  >> Press [ENTER] to try again... ")
			fmt.Scanln()
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
			fmt.Print("\033[1A\033[2K")
		}
	}

	hitungKorelasi(data, startDate, endDate, &hasil)

	if hasil.taskGood > 0 {
		pGood = (hasil.doneGood * 100) / hasil.taskGood
	}
	if hasil.taskNetral > 0 {
		pNetral = (hasil.doneNetral * 100) / hasil.taskNetral
	}
	if hasil.taskBad > 0 {
		pBad = (hasil.doneBad * 100) / hasil.taskBad
	}

	Header("Main Menu > Mind Compass > Insight Hub > Correlation")
	fmt.Println(" │                  . ˚ ₊  ⌕  TIME CAPSULE REPORT  ⌕  ₊ ˚ .                 │")
	fmt.Printf(" │                         [ %-8s ➔  %-8s ]                         │\n", startDate, endDate)
	fmt.Println(" ├──────────────────────────────────────────────────────────────────────────┤")
	fmt.Println(" │  [ ✿ ] Periodic Vibe Summary                                             │")
	fmt.Printf(" │        ✧ High Spirits : %-3d days                                         │\n", hasil.goodMood)
	fmt.Printf(" │        ⊹ Balanced     : %-3d days                                         │\n", hasil.netralMood)
	fmt.Printf(" │        ♡ Low Battery  : %-3d days                                         │\n", hasil.badMood)
	fmt.Println(" │                                                                          │")
	fmt.Println(" │  [ ✦ ] Productivity Impact (Tugas yang Selesai Berdasarkan Mood)         │")
	fmt.Println(" │                                                                          │")
	fmt.Println(" │    ✧ On High Spirits (Scale 4-5)                                         │")
	fmt.Printf(" │      Cleared : %-3d out of %-3d tasks ┊ %-20s %3d%%          │\n", hasil.doneGood, hasil.taskGood, grafikBar(pGood), pGood)
	fmt.Println(" │    ⊹ Balanced Energy (Scale 3)                                           │")
	fmt.Printf(" │      Cleared : %-3d out of %-3d tasks ┊ %-20s %3d%%          │\n", hasil.doneNetral, hasil.taskNetral, grafikBar(pNetral), pNetral)
	fmt.Println(" │    ♡ Low on Battery (Scale 1-2)                                          │")
	fmt.Printf(" │      Cleared : %-3d out of %-3d tasks ┊ %-20s %3d%%          │\n", hasil.doneBad, hasil.taskBad, grafikBar(pBad), pBad)
	fmt.Println(" ╰──────────────────────────────────────────────────────────────────────────╯")
	fmt.Print("\n   ₊˚.♡ Hit [ENTER] to return... ")
	fmt.Scanln(&hold)
}

func main(){
	var dataBase MindFlowData
	var reader *bufio.Reader
	var start bool 
	
	dataBase.totalMood = 0
	dataBase.totalTugas = 0
	
	generateDummyData(&dataBase)
	
	reader = bufio.NewReader(os.Stdin)
	start = true

	welcomeScreen()
	for start == true {
		mainMenu(&start, &dataBase, reader)
	}
}

func generateDummyData(data *MindFlowData) {
	//data untuk menu mood
	data.mood[0] = Mood{"12/01/26", "Kurang tidur parah, pusing banget", 1}
	data.mood[1] = Mood{"20/05/26", "Got a compliment soal tubes", 5}
	data.mood[2] = Mood{"04/04/26", "Biasa aja, flat banget hari ini", 3}
	data.mood[3] = Mood{"15/01/26", "Nothing special today, flat banget", 3}
	data.mood[4] = Mood{"12/01/26", "lagi burnout, butuh istirahat bentar", 2}
	data.mood[5] = Mood{"30/03/26", "Fokus nugas ditemenin my fav coffee", 4}
	data.mood[6] = Mood{"08/03/26", "Finally tubes kelar, lega bisa nyantai", 5}
	data.mood[7] = Mood{"04/04/26", "duh bete, dosennya killer banget deh", 1}
	data.mood[8] = Mood{"20/05/26", "Sorenya agak capek tapi aman lah", 3}
	data.mood[9] = Mood{"10/02/26", "Lagi dapet ide cemerlang buat project", 5}
	data.mood[10] = Mood{"15/01/26", "Agak mager tapi harus gerak", 3}
	data.mood[11] = Mood{"22/04/26", "Capek bgt abis kumpul organisasi", 2}
	data.mood[12] = Mood{"10/02/26", "Diskusi sama temen seru bgt", 4}
	data.mood[13] = Mood{"22/04/26", "Sedih dapet nilai kuis kecil", 1}
	data.totalMood = 14
	
	//data untuk menu tugas
	data.tugas[0] = TugasHarian{"20/05/26", "Kuis Alpro", 45, 5, true}
	data.tugas[1] = TugasHarian{"12/01/26", "Drafting esai", 45, 1, false}
	data.tugas[2] = TugasHarian{"04/04/26", "practice buat interview", 30, 3, false}
	data.tugas[3] = TugasHarian{"08/03/26", "Revisi laporan tubes", 60, 4, true}
	data.tugas[4] = TugasHarian{"12/01/26", "Tugas Kalkulus", 90, 4, false}
	data.tugas[5] = TugasHarian{"20/05/26", "mencatat materi ujian", 60, 2, false}
	data.tugas[6] = TugasHarian{"30/03/26", "Tugas basis data", 30, 2, true}
	data.tugas[7] = TugasHarian{"15/01/26", "Praktikum alpro", 90, 3, true}
	data.tugas[8] = TugasHarian{"04/04/26", "perbaiki design", 180, 5, true}
	data.tugas[9] = TugasHarian{"15/01/26", "Update dokumentasi", 30, 2, false}
	data.tugas[10] = TugasHarian{"22/04/26", "Cek referensi", 20, 1, false}
	data.tugas[11] = TugasHarian{"10/02/26", "Finalisasi project", 120, 5, true}
	data.tugas[12] = TugasHarian{"15/01/26", "Review materi", 40, 2, true}
	data.tugas[13] = TugasHarian{"10/02/26", "Submit berkas", 15, 4, true}
	data.totalTugas = 14
}