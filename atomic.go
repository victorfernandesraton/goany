package goany

import (
	"errors"
	"sync/atomic"
	"time"
)

func AtomicReq(text string, count *int64, res chan<- Response, err chan<- error) {
	timeout := len(text)
	time.Sleep(time.Duration(timeout) * time.Second)
	result := Response{
		Value: true,
	}
	if text == "error" {
		err <- errors.New("simple error response")
	}
	atomic.AddInt64(count, 1)
	res <- result
}
