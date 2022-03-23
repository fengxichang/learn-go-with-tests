package main

import (
	"learn-go-with-tests/mocking"
	"os"
)

func main() {
	sleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
