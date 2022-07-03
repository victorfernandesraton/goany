package goany_test

import (
	"fmt"
	"testing"

	"github.com/victorfernandesraton/goany"
)

func TestSimpleSemaphoreConcurrent(t *testing.T) {
	result := make(chan *goany.ResultStruct)
	go goany.Semaphore("simple result but take a time", result)
	go goany.Semaphore("error", result)

	var responses []goany.ResultStruct

	defer func() {
		close(result)
	}()

	for {
		resposne := <-result
		responses = append(responses, *resposne)
		if len(responses) == 2 {
			break
		}
	}

	for _, v := range responses {
		fmt.Println(v.Value)
		fmt.Println(v.Error)

	}

	if len(responses) != 2 {
		t.Fail()
	}

}
