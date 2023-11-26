package main

import (
	"fmt";
	"strings"
)


type Blacklist  func(string) bool;

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
        fmt.Printf("You are blocked %s\n", name);
    } else {
        fmt.Printf("Welcome %s\n", name);
    }
}

func main() {

	// ini kita buat func dimasukan kedalam nilai sebuah function anonymous
	blacklist := func(name string) bool {
		return strings.EqualFold(name,"anjing");
	}
	registerUser("Jhon", blacklist);
	registerUser("Anjing", blacklist);

	//bisa kayak gini function anonumousnya
	registerUser("Jhon", func(name string) bool{
		return strings.EqualFold(name,"anjing");
	});
	registerUser("Anjing", func(name string) bool{
		return strings.EqualFold(name,"anjing");
	});

}