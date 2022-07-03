package goany_test

import (
	"fmt"
	"testing"

	"github.com/victorfernandesraton/goany"
)

func TestValidAtomic(t *testing.T) {
	t.Run("Should be execution one sucess", func(t *testing.T) {
		var count int64
		result := make(chan goany.Response, 1)
		errCh := make(chan error, 1)
		go goany.AtomicReq("test", &count, result, errCh)
		response := <-result
		if !response.Value {
			t.Fail()
		}
	})
	t.Run("Should be execution one error", func(t *testing.T) {
		var count int64
		result := make(chan goany.Response, 1)
		errCh := make(chan error, 1)
		go goany.AtomicReq("error", &count, result, errCh)
		response := <-result
		if !response.Value {
			t.Fail()
		}
	})

	t.Run("Should be execution one error and one sucess", func(t *testing.T) {
		var count int64
		var response goany.Response
		result := make(chan goany.Response, 1)
		errCh := make(chan error, 1)
		go goany.AtomicReq("sucess but after a ling time", &count, result, errCh)
		go goany.AtomicReq("error", &count, result, errCh)

		select {
		case <-result:
			fmt.Println(result)
			response = <-result
		default:
		}
		if response.Value {
			t.Fail()
		}
	})
}
