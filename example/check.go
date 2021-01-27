package main

import (
	"fmt"
	"time"

	scanner "github.com/AviParampampam/go-cam-scanner"
)

func main() {
	addrss := scanner.Search(time.Millisecond * 750)
	fmt.Println(addrss)
}
