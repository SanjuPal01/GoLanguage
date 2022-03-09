package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func main() {
	str := "Sanju"
	mymd5 := md5.Sum([]byte(str))
	mySHA1 := sha1.Sum([]byte(str))
	mySHA256 := sha256.Sum256([]byte(str))
	fmt.Printf("MD5: %x\n", mymd5)
	fmt.Printf("SHA1: %x\n", mySHA1)
	fmt.Printf("SHA256: %x\n", mySHA256)
}
