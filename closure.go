package main;

import "fmt";

func main() {
	//closure jangan sering dipakek karena berbahaya


	counter :=0;
	



	increment := func(){
		fmt.Println("Increment");
		counter++;
	}

	increment();
	increment();
	increment();

	fmt.Println(counter);

}