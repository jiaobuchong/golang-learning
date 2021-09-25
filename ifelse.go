package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

/*
if condition {
} else if condition {
} else {
}*/
func main() {
	//str := "sam_brans@hotmail.com"
	//md5Str := md5V(str)
	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
