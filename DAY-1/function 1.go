

package main


import "fmt"

func vals () (int, int){

	return 3,7

}

func main() {
	a,b := vals()
	fmt.Println(a) // akan menampilkan angka 3
	fmt.Println(b)

	_, c := vals ()  //_, tidak akan menganggap output yang pertama
	fmt.Println(c)// akan menampilkan angka 7 saja
}