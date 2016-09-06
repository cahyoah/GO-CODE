package  main


import "fmt"

func main() {
	var a [5] int //membuat variabel baru array dengan panjang 5
	fmt.Println ("emp: ", a) //menampilkan tabel array dengan isi 0


	a[4] = 100 //akan mengisi elemen ke 4 pada tabel dengan 0
	fmt.Println("set :", a) 
	fmt.Println("get :", a[4])

	fmt.Println("length:", len(a) )//mencetak panjang array

	b:= [5] int {1,2,3,4,5} 
	fmt.Println("dcl:",b)

	var twoD [2][3]int
	for i:=0;i<2;i++{

		for j := 0; j<3;j++{

			twoD [i][j] = i+j //akan mengisi tabel 2 dimensi  dengan perulangan
		}
	}
	fmt.Println("2d : ", twoD) // memnampilkan tabel 2 dimensi
}