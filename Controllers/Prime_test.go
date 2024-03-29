package Controllers

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		n        int
		expected bool
		msg      string
	}{
		{"prime7", 7, true, "7 is a prime number"},
		{"prime0", 0, false, "0 is not prime"},
		{"prime8", 8, false, "8 is divisible by 2"},
		{"prime-99", -99, false, "Negative no cant be prime"},
	}

	for _, val := range primeTests {
		result, msg := IsPrime(val.n)
		if result != val.expected || msg != val.msg {
			t.Errorf("testtttt failed")
		}
	}
}

func Test_isPrompt(t *testing.T) {
	oldIn := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	Prompt()

	_ = w.Close()

	os.Stdout = oldIn

	out, _ := io.ReadAll(r)

	if string(out) != "->" {
		t.Error("failed in prompt")
	}
}

func Test_Intro(t *testing.T) {
	oldIn := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	Intro()

	_ = w.Close()

	os.Stdout = oldIn

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Error("failed in prompt")
	}
}
