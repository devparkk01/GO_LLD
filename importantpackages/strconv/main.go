package main 

// strings to number and number to strings conversion 
import (
	"fmt"
	"strconv"
)

func main() {
	num := 42
	s := strconv.Itoa(num)       
	fmt.Println(s) // "42"
	num, err := strconv.Atoi("123") 
	if err != nil {
		panic(err)
	}
	fmt.Println(num)  // 123

	num, err = strconv.Atoi("s123")
	if err != nil {
		panic(err)
	}
	fmt.Println(num) // never gets printed because of error in previous line 


}