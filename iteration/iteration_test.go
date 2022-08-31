package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 10)
	fmt.Println(result)
	// Output: aaaaaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestContains(t *testing.T) {
	str := "Brian is fun"

	got := strings.Contains(str, "fun")
	want := true

	t.Logf("got: %t, want: %t", got, want)
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
