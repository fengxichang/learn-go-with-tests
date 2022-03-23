package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls += 1
}

type CountdownOperationSpy struct {
	Calls []string
}

func (c *CountdownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, "sleep")
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, "write")
	return
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
