package main

import (
	"fmt"
)

func main(){
	fmt.Println(palindrom("Aba"))
}

func palindrom (some string) bool {
	for i:=0; i < len(some)/2; i++ {
		if (some[i] != some[len(some) -1 - i] ) {
			return false
		}
	}	
	return true
}