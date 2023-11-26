package main;

import "fmt";

func factorialLoop(value int) int {
	resulst:=1;
	for i:=value; i>0; i--{
        resulst *=i;
    }
	return resulst;
}
func factorialRecursive(value int) int {
	if value == 1 {
        return 1;
    }
	return value * factorialRecursive(value-1);
	
}

func main() {
	//kalo pakek loop 
	fmt.Println(factorialLoop(10)); 
	//kalo pakek function recursive 
	fmt.Println(factorialRecursive(10));
}