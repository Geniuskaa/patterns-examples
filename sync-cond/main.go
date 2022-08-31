package main

import (
	"fmt"
	"sync"
)

var (
	c  = sync.NewCond(&sync.Mutex{})
	wg = sync.WaitGroup{}

	free = false
)

func main() {
	// Signal example ///////////////////////////////////////////////
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	c.L.Lock()
	//
	//	for !free {
	//		c.Wait()
	//	}
	//
	//	fmt.Println("work")
	//	c.L.Unlock()
	//}()
	//
	//free = false
	//<-time.After(1 * time.Second)
	//free = true
	//c.Signal()
	//
	//wg.Wait()

	// Broadcast example ///////////////////////////////////////////////
	sumChn := make(chan int64, 2)

	size := 10
	arr := make([]int64, size)
	prevInd := 0
	wg.Add(2)
	for i := 2; i > 0; i-- {
		rightWall := size / i

		go func(ind int, cond *sync.Cond, prev int, right int, w *sync.WaitGroup) {
			cond.L.Lock() //Lock можно не вешать, работает и без него
			defer func() {
				cond.L.Unlock()
				w.Done()
			}()

			cond.Wait()
			fmt.Println("some")

			sum := int64(0)
			a := arr[prev:right]
			for _, val := range a {
				sum += val
			}
			sumChn <- sum

		}(i, c, prevInd, rightWall, &wg)

		prevInd = rightWall
	}

	for i, _ := range arr {
		arr[i] = int64((i + 1) * 2)
	}

	c.Broadcast()
	wg.Wait()
	s := int64(0)

	l := len(sumChn)
	for i := 0; i < l; i++ {
		s += <-sumChn
	}
	fmt.Println(s)

	//gettingReadyForMissionWithCond()
}

//func gettingReadyForMissionWithCond() {
//	cond := sync.NewCond(&sync.Mutex{})
//	go gettingReady(cond)
//	workIntervals := 0
//	cond.L.Lock()
//	for !free {
//		workIntervals++
//		cond.Wait()
//	}
//	cond.L.Unlock()
//	fmt.Println("We done job! For such times: ", workIntervals)
//}
//
//func gettingReady(cond *sync.Cond) {
//	time.Sleep(time.Duration(10) * time.Second)
//	free = true
//	cond.Signal()
//}
