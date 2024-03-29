package Controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func TryAnonymous() {

	var random = func() {
		i := rand.Intn(100)
		if i%2 == 0 {
			fmt.Println("golang")
		}
	}

	var err1 = func(err error) {
		fmt.Println(err)
	}
	rand.Seed(time.Now().UnixNano())
	err1(errors.New("Java"))
	random()
}
