package integers

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	assert := func(t *testing.T, got, want int) {
		if got != want {
			t.Errorf("Got '%d', expected '%d'", got, want)
		}
	}

	t.Run("Sums two integers", func(t *testing.T) {
		got := Sum(3, 4)
		want := 7

		assert(t, got, want)
	})
}

func ExampleSum() {
	sum := Sum(4, 4)
	fmt.Println(sum)
	// Output: 8
}
