package main

import "fmt"

// anay atau interface kosong artinya bisa tipe data apapun
func Ups() any {
	// return 1
	// return true
	return "any"
}

// bisa pakek anny atau interface {} seperti ini untuk membuat tipe data apapun bisa masuk
// di sini kalo kita inputkan sembarang tipe data apapun tetap bisa masuk karena any atau apapun
// mirip di java dimana kalo kita kasi object berarti apapun bisa di retrun
func TestUps(a interface{}) interface{} {
	return a
}
func main() {

	fmt.Println(Ups())

	//bisa dilihat kita kasi tipe data apapun tetap hasilnya keluar karena
	// retrun dan parameter dari func ini adalah any
	fmt.Println(TestUps("foo"))
	fmt.Println(TestUps(1))
}
