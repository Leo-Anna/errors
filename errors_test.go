package errors

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("test")
	fmt.Println(errorStack(err))
}

func TestStack(t *testing.T) {
	fmt.Println(errorStack(test2()))
}

func test1() error {
	return New("err01")
}

func test2() error {
	err := test1()
	return Annotate(err, "err02")
}
