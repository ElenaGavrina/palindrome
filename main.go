package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(palindrom(strings.ToLower("Aba")))
}

func palindrom (some string) bool {
	for i:=0; i < len(some)/2; i++ {
		if (some[i] != some[len(some) -1 - i] ) {
			return false
		}
	}	
	return true
}
