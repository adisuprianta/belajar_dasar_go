package main

import "fmt";

func logging(){
	fmt.Println("Selesai memanggil function");
}
func runAplication(){

	//defer membuat function logging tetap bejalan walau ada error
	// contoh bisa di lihat di panic.go 
	defer logging();
	fmt.Println("Run Apliactivation");
}


func main() {
	runAplication();
}