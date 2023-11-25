package main;

import "fmt";

func getGoodBye(name string) string {
	return "Good Bye " + name;
}

func main() {
	//first-class citizens, itu berarti fungsi dapat digunakan
	// dan diperlakukan dengan cara yang sama seperti tipe data lainnya, 
	//seperti integer, string, atau array.
	// seperti dibawah ini variable goodBye memiliki tipe data dari funtion getGoodBye
	// kalo goodBye := getGoodBye("kevin") ini artinya kita menggambil nilai retrun dari 
	// funtion getGoodBye
	goodBye:= getGoodBye;
	fmt.Println(goodBye("Kevin"));

	
}