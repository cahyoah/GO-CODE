package main 


import "fmt"

func main (){
	
  i := 1 // nilai i diinisialisasi dengan 1
  for i <= 3 {
		fmt.Println (i);  // akan melakukan perulangan sesuai dengan i 
		i = i +1 
	
  }
  
  
  for j := 7; j<=9; j++ { // nilai i diinisialisasi dengan 7
		fmt.Println (j); // akan melakukan perlangan sesuai degan i 
		
	
  }
  
  for {
  
		fmt.Println("loop")
		break
  }

}

