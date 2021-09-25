package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {

	fmt.Println(md5V("jaxson.zhou@lunarax.io"))

}


func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
