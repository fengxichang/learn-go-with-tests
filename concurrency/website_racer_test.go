package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test faster url", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(10 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl, 10*time.Second)

		if want != got {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("return an error if server does not response within 10s", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Second)
		fastServer := makeDelayServer(10 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		_, err := Racer(slowUrl, fastUrl, 10*time.Second)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
