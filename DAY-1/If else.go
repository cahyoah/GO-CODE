package main 


import "fmt"

func main (){
	
	if 7%2 == 0{ // akan dieksekusi jika angka input ganjil
		fmt.Println ("7 is even") 
	}else {
		fmt.Println ("7 is odd")
	}
	
	if 8%4 == 0{ // akan dieksekusi jika angak yang dinisialisasi habis dibagi dengan 4
		fmt.Println ("8 is divisible by 4")
	}
	
	if  num := 9; num < 0 { //nilai num diinisialisasi dengan 9 
								// akan dieksekusi jika num < 0
		fmt.Println (num, "is negative")
	}else if num < 10 { // akan dieksekusi jika num kurang dari 10
		fmt.Println (num, "has 1 digit")
	}else { //akan dieksekusi jika tidak masuk dalam kedua kondisi
		fmt.Println (num, "has multiple digit")
	}

}

