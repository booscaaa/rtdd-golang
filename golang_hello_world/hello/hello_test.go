package hello_test

import (
	"testing"

	"github.com/booscaaa/golang_hello_world/hello"
)

func TestPrintHello(t *testing.T) {
	got := hello.PrintHello("Kevin")
	want := "Hello, Kevin"

	if got != want {
		t.Errorf("Erro no teste")
	}
}
