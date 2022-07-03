package goany

import (
	"errors"
	"time"
)

type ResultStruct struct {
	Value string `json:"value"`
	Error error  `json:"error"`
}

func Semaphore(text string, res chan<- *ResultStruct) {
	timeout := len(text)
	time.Sleep(time.Duration(timeout) * time.Second)
	if text == "error" {
		result := &ResultStruct{
			Value: "",
			Error: errors.New("some error"),
		}
		res <- result

	} else {
		result := &ResultStruct{
			Value: text,
			Error: nil,
		}
		res <- result
	}
}
