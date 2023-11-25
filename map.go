package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{};
	// person["name"] = "John";
	// person["age"] = "30";

	person := map[string]string{
		"name": "John",
		"age":  "30",
	}
	fmt.Println(person["name"]);
	fmt.Println(person["age"]);
	fmt.Println(person);

	//maka akan menampilkan default value dari jenis datanya
	// hal ini terjadi karena key "salah" tidak ada
	// fmt.Println(person["salah"]);

	//functions yang ada di map ada len(), map[key], map[key] =value, 
	//make (map[tipeKey]tipevalue)
	//delete(map, key)

	buku:= make(map[string]string);
    buku["judul"] = "buku";
	buku["penulis"] = "John";
	buku["tahun"] = "2020";
	buku["wrong"] = "ups";

	fmt.Println(buku);
	delete(buku,"wrong");
	fmt.Println(buku);

}