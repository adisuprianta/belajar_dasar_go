package main

import "fmt"

func random()any{
	return "ok"
}

func main() {
	
	var result any=random()
	resultString := result.(string)
	fmt.Println(resultString)

	// jangan melakukan assertions ke tipe data berbeda karena akan muncul panic
	//saat kode dibawah di eksekosi maka akan terjadi panic karena tipe datanya berbeda
	// yang awalnya string di ubah jadi int
	// resultInt := result.(int)
    // fmt.Println(resultInt)

	// untuk mengatasi itu kita bisa gunakan swicth seperti dibawah ini
	//otomatis akan berubah tipe ranDatanya sesuai casenya 
	// ini membuat program kita tidak terjadi panic
	switch ranData := result.(type) {
		case string:
            fmt.Println("string ",ranData)
        case int:
            fmt.Println("int ",ranData)
        default:
            fmt.Println("unknown type")
	}
}