package Controllers

import "fmt"

func Intro() {
	fmt.Println("Enter a whole number")
	Prompt()
}

func Prompt() {
	fmt.Print("->")

}

func IsPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime", n)
	}

	if n < 0 {
		return false, "Negative no cant be prime"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number", n)
}
