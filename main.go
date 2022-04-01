package main

import (
	"learn-go-with-tests/context"
)

func main() {
	//sleeper := &mocking.DefaultSleeper{}
	//mocking.Countdown(os.Stdout, sleeper)

	// 没有缓冲channel
	//channel.NoBufferChannel()

	// 有缓冲区channel
	//channel.BufferChannel()

	// done channel
	//context.DoneChannel()

	// done channel with context
	context.DoneChannelWithContext()
}
