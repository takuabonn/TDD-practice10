package main

import "testing"
import "fmt"
func TestNumberOfShouts(t *testing.T) {
    
	t.Run("A = 1 && B = 4 && K = 2 -> X = 2", func(t *testing.T) {
        got := numberOfShouts(1,4,2)
		want := 2

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
    })

	t.Run("A = 7 && B = 7 && K = 10", func(t *testing.T) {
        got := numberOfShouts(7,7,10)
		want := 0

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
    })

	t.Run("A = 31 && B = 415926 && K = 5", func(t *testing.T) {
        got := numberOfShouts(31,415926,5)
		want := 6
		fmt.Println(got)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
    })
}