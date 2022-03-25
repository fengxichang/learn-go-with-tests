package concurrency

import (
	"errors"
	"net/http"
	"time"
)

//func Racer(a, b string) (winner string)  {
//	aDuration := measureResponseTime(a)
//
//	bDuration := measureResponseTime(b)
//
//	if aDuration < bDuration {
//		return a
//	}
//	return b
//}
//
//func measureResponseTime(url string) time.Duration {
//	start := time.Now()
//	http.Get(url)
//	duration := time.Since(start)
//
//	return duration
//}

var TimeoutError = errors.New("time out")

func Racer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", TimeoutError
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		ch <- struct{}{}
	}()

	return ch
}