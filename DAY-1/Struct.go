package main 
import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {

	return r.width * r.height // akan mengembalikan hasil luas
}

func (r rect) perim() int {

	return 2*r.width + 2*r.height //akan mengembalikan hasil keliling

}

func main() {
	
	r := rect {width : 10, height : 5} // mengisi komponen pada rectn dengan tinggi 5 dan lebar 10

	fmt.Println("area : ", r.area()) // akan menmpilkan hasil dari luas dengan fungsi luas
	fmt.Println("perim: ", r.perim()) // akan menampilkan hasil dari keliling dengan fungsi perim

	rp := &r // mengisi variabel rp dengan nilai yang sama dengan r
	fmt.Println("area : ", rp.area()) // akan menampilkan hasil dari luas dengan fungsi luas dengan variabel rp
	fmt.Println("perim: ", rp.perim()) // akan menampilkan hasil dari keliling dengan fungsi perim dengan variabel r[]
}