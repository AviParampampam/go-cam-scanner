package main

import (
	"fmt"
	"time"

	"github.com/AviParampampam/go-cam-scanner/pkg/scanner"
)

func main() {
	addrss := scanner.Search(time.Millisecond * 750)
	fmt.Println(addrss)
}
