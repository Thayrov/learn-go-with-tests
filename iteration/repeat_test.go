package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("function must repeat a character", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertFunc(t, expected, repeated)
	})
	t.Run("function must repeat a character as many times are passed in second param", func(t *testing.T) {
		repeated := Repeat("k", 7)
		expected := "kkkkkkk"
		assertFunc(t, expected, repeated)
	})
}

func ExampleRepeat() {
	repeated := Repeat("y", 8)
	fmt.Println(repeated)
	// Output: yyyyyyyy
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assertFunc(t testing.TB, expected, repeated string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
