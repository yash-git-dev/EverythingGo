package Controllers

import (
	"fmt"
	"sync"
)

func executeGoroutine(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("This is a goroutine")

}

func GoRoMain() {
	var wg sync.WaitGroup
	wg.Add(1)
	go executeGoroutine(&wg)
	wg.Wait()
}
