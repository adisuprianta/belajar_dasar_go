package main

import "fmt"

func main() {
	number := [...]int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,13}
	slice1 := number[4:6]
	fmt.Println(slice1)
	slice2 := number[:3]
	fmt.Println(slice2)
	slice3 := number[3:]
	fmt.Println(slice3)
	slice4 := number[:]
	fmt.Println(slice4)

	//len mendapatkan slice
	fmt.Println("len")
	fmt.Println(len(slice1))
	fmt.Println("cap")
	fmt.Println(cap(slice1))

	//append slice untuk menambahkan kapasistas slice
	slice1 = append(slice1, 10)
	
	fmt.Println(slice1)

	//cap mendapatkan kasitas slice
	fmt.Println("cap", cap(slice1))

	//make membuat slice baru 
	slice5 := make([]int, 10, 20)
    fmt.Println(slice5)




	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daysSlice1 := days[5:]
	fmt.Println(daysSlice1);
	//ini akan merubah nilai dari array asalnya yaitu days karena yang dicopy bukan valuenya tapi 
	//alamat memorynya
	daysSlice1[0]= "sabtu baru";
	daysSlice1[1] = "minggu baru";
	fmt.Println(daysSlice1);
	fmt.Println(days);



	daysSlice1 = append(daysSlice1, "libur baru");
	daysSlice1[0] = "Sabtu Lama";
	
	//maka di dalam logic golang akan mungkin akan dibuatkan array baru di mana dibelakangnya 
	//ditambah value baru
	//dan daysSlice1 pun tidak ada hubungannya dengan array days 
	//	daysBaru := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu","libur baru"}
	fmt.Println(daysSlice1)
	fmt.Println(days);
	fmt.Println();
	//cara buat slice 
	newSlice :=make([]string, 2,5);
	newSlice[0] = "a";
	newSlice[1] = "b";
	
	fmt.Println(newSlice);
	fmt.Println("Cap",cap(newSlice));
	fmt.Println("len",len(newSlice));

	//newSlice[2] = "c";
	//cara diatas tidak akan berhasil karena leng slicenya cuma dua kita harus menggunakan append
	//untuk nambah lenght dari capisiti yang sudah di initialisasi yaitu 5
	
	newSlice = append(newSlice,"c");

	fmt.Println("sesudah append");
	fmt.Println(newSlice);
	fmt.Println("Cap",cap(newSlice));
	fmt.Println("len",len(newSlice));

	newSlice1 := append(newSlice,"d");

	//ini masi pakek array yang sama karena masi dalam kapasitinya dan saat kita ubah nilai index ke 0 
	//dari newSlice1 maka newSlice juga akan berubah
	newSlice1[0] = "aa";
	fmt.Println("sesudah ubah index ke 0");
	fmt.Println(newSlice1);
	fmt.Println(newSlice);
	fmt.Println("Cap",cap(newSlice));
	fmt.Println("len",len(newSlice));


	//copy slice
	fromSlice := days[:];
	toSlice := make([]string, len(fromSlice),cap(fromSlice));
	
	copy(toSlice, fromSlice);

	fmt.Println(toSlice);
	fmt.Println(fromSlice);

	iniArray := [...]int {1,2,3,4,5,6,7,8,9};
	iniSlice := []int {1,2,3,4,5,6,7};

	fmt.Println(iniArray);
	fmt.Println(iniSlice);

}	