package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := sync.RWMutex{}
	wg := sync.WaitGroup{}
	v := true
	//ch := make(chan struct{}, 1)

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(in int) {
			//time.Sleep(time.Duration(rand.Int63n(900)) * time.Millisecond)
			defer wg.Done()
			m.RLock()
			if in%2 == 0 {
				v = false
				fmt.Print("Default ")
			}
			fmt.Println("Goroutine num ", in, " val is ", v)
			m.RUnlock()
		}(i)
	}

	time.Sleep(time.Millisecond * 1)
	fmt.Println("Sent signal!")
	//ch <- struct{}{}

	wg.Wait()
}
