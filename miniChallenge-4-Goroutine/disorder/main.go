package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	bisa := make([]interface{}, 3)
	coba := make([]interface{}, 3)

	for i := 0; i < 3; i++ {
		bisa[i] = fmt.Sprintf("bisa%d", i+1)
		coba[i] = fmt.Sprintf("coba%d", i+1)
	}
	for i := 0; i < 4; i++ {
		wg.Add(2)
		go process(&wg, bisa, i+1)
		go process(&wg, coba, i+1)
	}

	wg.Wait()
}

// jujur gatau kenapa kak urutannya masih gak pas dan kena deadlock mungkin bisa boleh minta tolong koreksi yaa kak kurangnya dimana TIA kak
func process(wg *sync.WaitGroup, data []interface{}, id int) {
	defer wg.Done()

	fmt.Printf("%+v : %d\n", data, id)
}
