package printer

import (
	"testing"
)

func TestPrintHello(t *testing.T) {
	got := PrintHello("Igor")
	expected := "Hello, Igor!"

	if got != expected {
		t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
	}

}
func TestPrintHelloIvan(t *testing.T) {
	got := PrintHello("Ivan")
	expected := "Hello, Ivan!"

	if got != expected {
		t.Fatalf(`PrintHello("Ivan") = %q, want %q`, got, expected)
	}
}
