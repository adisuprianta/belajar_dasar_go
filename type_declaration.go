package main

import "fmt"

func main() {
	type NoKtp string

	var ktp NoKtp = "1234567890123456"
	fmt.Println(ktp)
	fmt.Println(NoKtp("213141421412"))
	
}