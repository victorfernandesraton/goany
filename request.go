package goany

import (
	"fmt"
	"time"
)

type Response struct {
	Value bool  `json:"value"`
	Error error `json:"error"`
}

func RequestOne(ch chan<- Response) {
	fmt.Println("teste 1")
	response := &Response{
		Value: true,
		Error: nil,
	}
	ch <- *response
	defer close(ch)
}

func RequestTwo(ch chan<- Response) {
	time.Sleep(10 * time.Second)
	fmt.Println("teste 2")
	response := &Response{
		Value: true,
		Error: nil,
	}
	ch <- *response
	defer close(ch)
}
