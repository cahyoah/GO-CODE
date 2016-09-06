package main 


import "fmt"




func plus (a int, b int) int{
	

	return a + b // akan mengembalikan hasil penjumlahan a+b

}


func plusPlus (a , b, c int) int{

	
	
	return a + b + c // akan mengembalikan hasil penjumlahan a+b+c

}


func main (){
	
	res := plus (1,2) // akan melakukan penjumlahan 1+2 dengan fungsi plus
	fmt.Println("1+2 =", res) // akan menampilkan hasil penjumlahan 1+2

	res = plusPlus (1, 2, 3) //akan melakukan penjumlahan dengan 1+2+3 dengan fungsi plus plus
	fmt.Println("1+2+3 = ",res) // akan menampilkan hasil penjumlahan 1+2+3
	
}

