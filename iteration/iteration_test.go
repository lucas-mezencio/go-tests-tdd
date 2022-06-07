package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		if got != want {
			t.Errorf("got %q - want %q", got, want)
		}
	}

	t.Run("concatenacao de 5 caracteres", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

}
