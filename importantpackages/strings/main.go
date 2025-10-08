package main 
// Go strings are immutable byte slices (utf-8 encoded).

import (
	"fmt"
	"strings"
)


func main() {
	s := "INTERSTELLAR"
	t := "STELL"
	var index int 
    // Searching and checking 

	// check if substring exists
	fmt.Println(strings.Contains(s, t)) // true 
	// check if it has a prefix 
	fmt.Println(strings.HasPrefix(s, "INTER")) // true 
	// Check if it has a suffix 
	fmt.Println(strings.HasSuffix(s, "LLAR")) // true 
 
    // first index of substring, -1 if not found
	index = strings.Index(s, "TERS")
	fmt.Println(index)  // 2 
	index = strings.Index(s, "STAR")
	fmt.Println(index) // -1


	
	// Splitting and joining 

	s = "go,all,the,way,or,don't,even,try"
	words := strings.Split(s, ",") 
	fmt.Println(words) // [go all the way or don't even try]

	joined := strings.Join(words, ",") // go,all,the,way,or,don't,even,try
	fmt.Println(joined) 
	if (joined == s) {
		fmt.Println("s and joined are equal") // s and joined are equal
	}

	
	// Replacing and modifying 

	s = strings.ToUpper(s) 
	fmt.Println(s) // GO,ALL,THE,WAY,OR,DON'T,EVEN,TRY

	s = strings.ToLower(s)
	fmt.Println(s) // go,all,the,way,or,don't,even,try

	s = "  hello world.  "
	fmt.Println(strings.TrimSpace(s))  // hello world.

	// Counting 
	s = "bananaaaana"
	// Count counts the number of non-overlapping instances of substr in s.
	fmt.Println(strings.Count(s, "na")) // 2

}