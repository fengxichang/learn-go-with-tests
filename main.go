package main

import (
	"learn-go-with-tests/channel"
)

func main() {
	//sleeper := &mocking.DefaultSleeper{}
	//mocking.Countdown(os.Stdout, sleeper)

	// 没有缓冲channel
	channel.NoBufferChannel()

	// 有缓冲区channel
	//channel.BufferChannel()
}
