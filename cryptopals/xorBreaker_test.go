package cryptopals

import (
	"encoding/hex"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	t.Run("single vowel score", func(t *testing.T) {
		expected := float32(11.36)
		actual := stringScore([]byte("e"))

		verifyExpected(t, expected, actual)
	})
	t.Run("single consonant score", func(t *testing.T) {
		expected := float32(2.92)
		actual := stringScore([]byte("d"))

		verifyExpected(t, expected, actual)
	})
	t.Run("single word score", func(t *testing.T) {
		expected := float32(22.8)
		actual := stringScore([]byte("THE"))

		verifyExpected(t, expected, actual)
	})
}

func BenchmarkCalculateScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringScore([]byte("The quick brown fox jumped over the lazy dog"))
	}
}

func verifyExpected(t testing.TB, expected, actual float32) {
	if expected != actual {
		t.Errorf("expected %f but got %f", expected, actual)
	}
}

func TestSingleByteXor(t *testing.T) {
	t.Run("one character input", func(t *testing.T) {

	})
	t.Run("two character input", func(t *testing.T) {

	})
	t.Run("10 character input", func(t *testing.T) {

	})
}

func BenchmarkSingleByteXor(b *testing.B) {
	cipherTextHex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	cipherText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		breakSingleCharacterXor(cipherText)
	}
}
