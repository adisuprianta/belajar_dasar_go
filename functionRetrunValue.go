package main;

import "fmt";

func tambah(a int, b int) int{
	return a+b;
}
func getHello(name string) string{
	return "Hello " + name;
}
func getAdrees(address string) string{
	response := "Adrees " + address;
	return response;
}
func getNumber(number[]  int) []int{
	return number;
}

func main() {
	fmt.Println(tambah(2,3));
	result := getHello("Kevin Candra");
	fmt.Println(result);
	responseAdsress := getHello("Ragunan");
	fmt.Println(responseAdsress);
	numbers := getNumber([]int { 1, 2, 3});
	fmt.Println(numbers);
	numbers = append(numbers, 4);
	fmt.Println(numbers);
}