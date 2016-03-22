package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	m := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		go func(i int){
			m.Lock()
			fmt.Print(i," start ")
			time.Sleep(time.Second)
			fmt.Println(i,"end")
			m.Unlock()
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}
