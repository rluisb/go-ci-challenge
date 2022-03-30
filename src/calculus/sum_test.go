package calculus

import "testing"

func TestSum(t *testing.T) {
	got := Sum(5, 5)
	want := uint64(10)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}