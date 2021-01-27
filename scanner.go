package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func Search(timeoout time.Duration) (addresses []string) {
	var wg sync.WaitGroup

	for i := 3; i < 255; i++ {
		wg.Add(1)

		go func(a string) {
			defer wg.Done()

			if _, err := net.DialTimeout("tcp", a, timeoout); err == nil {
				addresses = append(addresses, a)
			}
		}(fmt.Sprintf("192.168.1.%d:554", i))
	}

	wg.Wait()
	return addresses
}
