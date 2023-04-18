package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(isPalindrom("Aba"))
}

func isPalindrom (some string) bool {
	l :=strings.ToLower(some)
	for i:=0; i < len(l)/2; i++ {
		if (l[i] != l[len(l) -1 - i] ) {
			return false
		}
	}	
	return true
}
