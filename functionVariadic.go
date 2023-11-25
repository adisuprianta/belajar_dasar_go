package main;

import "fmt";

// variadic function artinya parameters kayak array jadi bisa kita kasi berapapun jumlah
//parameters dari functionnya
func sumAll(numbers ... int) (total int){
	for _, n:=range numbers{
		total += n;
	}
	return total;
}

func main() {
	number :=sumAll(1,2,3,4,5);
	fmt.Println(number); 

	// dengan function sama tapi parameter bisa banyak karena variadic function
	//bagusnya variadic function dari pada parameter array atau slice 
	// saat mengisi parameternya kita tida usah seperti ini 
	//sumAll([]int{1,2,3,4,5}) atau sumAll([...]int{1,2,3,4})
	// tapi cukup sumAll(1,2,3,4) hal ini membuat lebih fleksibel
	number =  sumAll(1,2,3,4,5,6,7,8,9,10,11,12,13);
	fmt.Println(number);

	//ini slice
	numbers:=[]int {10,10,10,10};

	// ini kalo kita punya slice tapi kita ingin memasukan ke dalam 
	// function kita cukup nama variable slice trs titik 3 kali
	// seperti ini sumAll(numbers ...);
	number = sumAll(numbers ...);
	fmt.Println("ini hasil dari slice ",number);
}