package main

import "fmt"


// nill cuma bisa di tipe data interface map slice function pointer dan channel
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}

func main(){
	data := NewMap("")
	fmt.Println(data)
	 if data== nil {
		fmt.Println("nill")
	 }else{
		 fmt.Println(data)
	 }
	//  data["name"] = "adada"

	data = NewMap("adasda")
	 fmt.Println(data)
}