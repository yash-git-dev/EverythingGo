package Controllers

import (
	"fmt"
	"sync"
)

type in struct {
	value    int32
	oddChan  chan int32
	evenChan chan int32
}

var serverChan chan in

func Server() {
	//var wg sync.WaitGroup
	//wg.Add(len(serverChan))
	for i := range serverChan {
		if i.value%2 == 0 {
			i.evenChan <- i.value
		} else {
			i.oddChan <- i.value
		}
		//	wg.Done()
	}
	//wg.Wait()
	close(serverChan)
}

func Main() {

	arr := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	serverChan = make(chan in, len(arr))
	oddChan := make(chan int32)
	evenChan := make(chan int32)
	for idx := 0; idx < len(arr); idx++ {
		i := idx
		serverChan <- in{arr[i], oddChan, evenChan}
	}
	odds, evens := []int32{}, []int32{}
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))
	go func() {
		for newOdd := range oddChan {
			odds = append(odds, newOdd)
			wg.Done()
		}
	}()
	go func() {
		for newEven := range evenChan {
			evens = append(evens, newEven)
			wg.Done()
		}
	}()
	go Server()
	wg.Wait()

	for _, resultItem := range odds {
		fmt.Println(resultItem)
	}

	for _, resultItem := range evens {
		fmt.Println(resultItem)
	}
}
