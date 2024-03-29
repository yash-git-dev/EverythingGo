package Controllers

import (
	"fmt"
	"sync"
	"time"
)

func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func GoMain() {
	go Say("world")
	Say("hello")
}

func printer(n int, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		fmt.Println(n + i)
	}
	wg.Done()

}

// goroutine sequence 1-10
func printSeq(n int) {

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printer(i*10, &wg)

	}
	wg.Wait()
}
