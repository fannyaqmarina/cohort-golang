package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	bisa := make([]interface{}, 3)
	coba := make([]interface{}, 3)

	for i := 0; i < 3; i++ {
		bisa[i] = fmt.Sprintf("bisa%d", i+1)
		coba[i] = fmt.Sprintf("coba%d", i+1)
	}

	for i := 0; i < 4; i++ {

		wg.Add(4)
		go func(i int) {
			mu.Lock()
			process(&wg, coba, i+1)
			process(&wg, bisa, i+1)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
}

func process(wg *sync.WaitGroup, data []interface{}, id int) {
	defer wg.Done()
	fmt.Printf("%s : %+v\n", data, id)
}
