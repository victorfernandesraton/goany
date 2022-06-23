package goany_test

import (
	"testing"

	"github.com/victorfernandesraton/goany"
)

func TestAsyncWorker(t *testing.T) {
	t.Run("simple concurrent test", func(t *testing.T) {

		var result chan goany.Response

		go goany.RequestOne(result)
		go goany.RequestOne(result)

		data := <-result

		if data.Error != nil {
			t.Fail()
		}
	})
}
