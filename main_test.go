package main

import (
	"fmt"
	"testing"
)

func TestAddTwo(t *testing.T) {

	t.Log("you will see this log message")

	go func() {
		_ = addTwo(3, 2)
		t.Log("this will never happen")
	}()

	go func() {
		_ = addTwo(1, 2)
		t.Fatal("even worse, this will never happen")
	}()

	errChan := make(chan error)

	go func(e chan error) {
		defer close(e)
		err := fmt.Errorf("this will succeed in failure")
		e <- err
	}(errChan)

	err := <-errChan
	if err != nil {
		t.Fatalf("error came down the channel: %v", err)
	}
}
