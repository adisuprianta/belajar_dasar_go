package main;

import "fmt";

func main() {
	//for biasa
	// couter :=1;
	// for couter<=10{
	// 	fmt.Println("Perulangan yang",couter);
    //     couter++;
	// }
	// fmt.Println("Seleai");
	// for couter = 1; couter <= 10; couter++ {
	// 	fmt.Println("Perulangan yang",couter);
	// }
	// fmt.Println("Selesai");

	//kalo for biasa
	// names :=[]string{"Jhone", "Eko","Khannedy","Kurniawan"};
	// for i:=0 ; i < len(names) ; i++ {
	// 	fmt.Println("name: ",names[i]);
	// }

	//fore range di golang
	names :=[]string{"Jhone", "Eko","Khannedy","Kurniawan"};
	// tanda _ ini fungsinya kalo kita tidak perlu index atau key dari sebuah 
	// array, slice, atau map 
	for _, name := range names {
        fmt.Println("name: ",name);
    }
	
	fmt.Println("selesai")

	//ini kalo kita ingin mencari indexnya 
	for index, name := range names{
		fmt.Println("name: ",name," index ke: ",index);
	}

	
}