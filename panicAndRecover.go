package main

import 	"fmt"

func endApp() {
	fmt.Println("End App")

	// dengan recover ini maka application tidak langsung berhenti dan
	//akan lanjut ke kode selanjutnya
	message:=recover()
	fmt.Println("terjadi error messagenya ini",message)
}
func runApp(err bool) {

	// walau ada panic endApp tetap di eksekusi karena keyword defer
	defer endApp()
	if err {
		panic("error running")
	}
	// ini cara salah menggunakan recover() harusnya di taruh di endApp karena
	// endApp selalu di eksekusi sebab ada keyword defer
	// message:=recover()
	// fmt.Println("terjadi erro messagenya ini",message)
}

func main() {
	runApp(true)
	fmt.Println("app")
}
