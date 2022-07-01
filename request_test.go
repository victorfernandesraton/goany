package goany_test

import (
	"fmt"
	"testing"

	"github.com/victorfernandesraton/goany"
)

func TestAsyncWorker(t *testing.T) {
	t.Run("simple sucess only test", func(t *testing.T) {

		result := make(chan goany.Response, 1)
		errCh := make(chan error, 1)

		go goany.RequestOne(result, errCh)

		data := <-result

		if !data.Value {
			t.Fail()
		}
	})

	t.Run("simple error only test", func(t *testing.T) {

		result := make(chan goany.Response, 1)

		errCh := make(chan error, 1)

		go goany.RequestTwo(result, errCh)

		data := <-result

		if data.Value {
			t.Fail()
		}
	})
	t.Run("simple concurrent test", func(t *testing.T) {

		result := make(chan goany.Response, 1)

		errCh := make(chan error, 2)

		go goany.RequestOne(result, errCh)
		go goany.RequestTwo(result, errCh)

		data := <-result

		if data.Value {
			t.Fail()
		}
	})
	t.Run("simple concurrent error cases", func(t *testing.T) {

		result := make(chan goany.Response, 1)
		errCh := make(chan error, 2)

		go goany.RequestTwo(result, errCh)
		go goany.RequestTwo(result, errCh)

		data := <-result


		err := <- errCh
		fmt.Printf("%v\v", err)
		err = <- errCh
		fmt.Printf("%v\v", err)

		if data.Value {
			t.Fail()
		}
	})
}
