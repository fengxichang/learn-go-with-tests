package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

func Countdown(writer io.Writer, sleeper Sleeper) {

	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}

	sleeper.Sleep()
	fmt.Fprint(writer, "Go!")
}
