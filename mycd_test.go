package mycd

import "testing"

func TestPrintHi(t *testing.T) {
	want := "Hi!"
	if got := PrintHi(); got != want {
		t.Errorf("PrintHi() = %q, want %q", got, want)
	}
}
