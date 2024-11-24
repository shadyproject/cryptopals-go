package cryptopals

import "testing"

func BenchmarkFindEncrypted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findEncrypted()
	}
}
