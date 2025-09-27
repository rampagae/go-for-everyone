package mathx

import "testing"

func TestSum(t *testing.T) {
    got := Sum(2, 3)
    want := 5
    if got != want {
        t.Errorf("Sum(2,3) = %d; want %d", got, want)
    }
}


