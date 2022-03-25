package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment 3 times", func(t *testing.T) {
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("it run safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		count := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i:=0; i<wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				count.Inc()
				w.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounter(t, &count, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int)  {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}