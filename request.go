package goany

import (
	"errors"
	"fmt"
	"time"
)

type Response struct {
	Value bool `json:"value"`
}

func RequestOne(ch chan<- Response, err chan<- error) {
	time.Sleep(20 * time.Second)

	fmt.Println("teste 1")
	response := &Response{
		Value: true,
	}
	err <- nil
	ch <- *response
}

func RequestTwo(ch chan<- Response, err chan<- error) {
	time.Sleep(10 * time.Second)
	fmt.Println("teste 2")
	response := &Response{
		Value: false,
	}
	err <- errors.New("error creating")
	ch <- *response

}
