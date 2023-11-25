package main;

import "fmt";

func main() {
	//contoh breack untuk menghentikan perulangan
	for i := 0; i < 10; i++ {
		if i == 6 {
			break;
		}
		fmt.Println("Perulangan ke: ", i);
	}


	
}