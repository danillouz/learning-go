package loops

import (
	"fmt"
	"testing"
)

func TestRepeatFive(t *testing.T) {
	got := RepeatFive("x")
	want := "xxxxx"

	if got != want {
		t.Errorf("Got '%s', want '%s'", got, want)
	}
}

func ExampleRepeatFive() {
	repeated := RepeatFive("x")
	fmt.Println(repeated)
	// Output: xxxxx
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatFive("x")
	}
}
