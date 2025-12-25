package main

import (
	"ecommerce/cmd"
)

// var cnt int64

func main() {
	cmd.Serve()
	// var wg sync.WaitGroup
	// var mu sync.Mutex

	// for i := 1; i <= 100000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		mu.Lock()
	// 		defer wg.Done()
	// 		a := cnt
	// 		a = a + 1
	// 		cnt = a
	// 		mu.Unlock()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(cnt)
}
