package main

import (
	"fmt";
	"strings"
)

// ini alias dari function func(string) string
// ini tujuannya untuk mempersingkat 
// jika ada filter func(string, string,string, float) string
// itu kan panjang jika dijadikan parameter maka kita bisa buat aliasnya atau
// buat type declaration seperti ini
// type Filter func(string, string, string, float) string 
// ini jauh lebih singkat jadinya
type Fillter func(string) string;



// func sayHelleWithFillter(name string, fillter func(string) string) {
// 	fillterdName:=fillter(name);
// 	fmt.Println("hello", fillterdName);
// }


// fillter Fillter bisa dilakukan karena di atas kita sudah membuat alisnya 
// atau disebut type declaration di go
func sayHelleWithFillter(name string, fillter Fillter) {
	fillterdName:=fillter(name);
	fmt.Println("hello", fillterdName);
}
func spamFillter(name string) string{

	//ini cara memandingakn string case insensitive
	if strings.EqualFold(name, "anjing") {
		return "...";
	}else{
		return name;
	}
}

func main(){
	// function sebagai parameter dimana smapFellter itu adalah sebuah parameter
	//ini terjadi karena function di go disebut sebagau first class citizen 
	// dimana first class citizen artinya function bisa diartikan sebagai sebauah tipe data
	//berbeda di java class yang bisa sebagai tipe datanya atau sebagai first class citizennya
	sayHelleWithFillter("Anjing", spamFillter);
	sayHelleWithFillter("kevin", spamFillter);

	// bisa juga seperti ini
	fillter:=spamFillter;
	// artinya si fillter memiliki tipe data spamFillter
	// kembali lagi function di go itu disebut first class citizen
	sayHelleWithFillter("anjing", fillter);

}