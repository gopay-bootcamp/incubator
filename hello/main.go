package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second * 2)

	for {
		fmt.Println("Hello World!")
		<-ticker.C
	}
}
