package main

import (
	"fmt"
	"time"
)

func main() {
	unixNow := time.Now().Unix()
	if unixNow%2 == 0 {
		fmt.Printf("%d is ... Even!!!\n", unixNow)
	} else {
		fmt.Printf("%d is ... Odd!!!\n", unixNow)
	}
}

type MyInterface interface {
}
