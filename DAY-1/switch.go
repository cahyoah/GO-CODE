package main 


import "fmt"
import "time"

func main (){
	
	i := 2
	fmt.Print("write ", i, " as ")
	switch i { //akan menampilkan kondisi sesuai dengan i yang telah diinisialisasi
		case 1 :
			fmt.Println ("one") 
		case 2 :
			fmt.Println ("two")	
		case 3 :
			fmt.Println ("three")	

	}


	switch time.Now().Weekday() { //akan menampilkan kondisi sesuai dengan hari saat ini
		case time.Saturday, time.Sunday:
			fmt.Println ("It's Weekend")
		default :
			fmt.Println("It's afternoon")
	}

	t := time.Now()
	switch {		//akan menampilkan kondisi sesuai dengan jam saat ini

		case t.Hour() < 12 :
			fmt.Println ("It's before noon") // akan ditampilkan bila waktu saat ini kurang dari jam 12
		default:
			fmt.Println ("It's after noon") // akan ditampilkan bila waktu saat ini lebih dari jam 12
	}
}

