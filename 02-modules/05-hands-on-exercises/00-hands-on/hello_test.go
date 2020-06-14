package hello

import "testing"

func TestHello(t *testing.T) {
	want := "I can eat glass and it doesn't hurt me."
	if got := Glass(); got != want {
		t.Errorf("Glass() = %q, want %q", got, want)
	}
}
