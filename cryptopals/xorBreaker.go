package cryptopals

import "strings"

type CharacterFrequency map[byte]float32

func breakSingleCharacterXor(cipherText []byte) (key byte, plainText string) {
	bestScore := float32(0)
	var bestPlaintext string
	var bestKey byte

	for i := 0; i < 256; i++ {
		key := byte(i)
		plainText := singleByteXor(cipherText, key)

		score := stringScore(plainText)

		// keep track of the cost, higher is better
		if score > bestScore {
			bestScore = score
			bestKey = key
			bestPlaintext = string(plainText)
		}
	}

	return bestKey, bestPlaintext
}

func singleByteXor(bytes []byte, xor byte) (xored []byte) {
	xored = make([]byte, len(bytes))

	for i := range bytes {
		xored[i] = bytes[i] ^ xor
	}

	return
}

// calculate the likelihood that a given string is english based on character frequency
// higher score is better
func stringScore(str []byte) (score float32) {
	freqs := englishFrequencies()

	for _, char := range str {
		// i feel like there's a better way to do this
		c := strings.ToLower(string(char))[0]
		score += freqs[c]
	}
	return
}

// Source:
// Lee: E. Stewart. "Essays about Computer Security" (PDF). University of Cambridge Computer Laboratory. p. 181.
func englishFrequencies() (freqs CharacterFrequency) {
	// straight ported this from my rust implementation and there's some obvious
	// missing error handling and such
	freqs = CharacterFrequency{
		byte(' '): 12.17, // Whitespace
		byte('.'): 6.57,  // Others
		byte('a'): 6.09,
		byte('b'): 1.05,
		byte('c'): 2.84,
		byte('d'): 2.92,
		byte('e'): 11.36,
		byte('f'): 1.79,
		byte('g'): 1.38,
		byte('h'): 3.41,
		byte('i'): 5.44,
		byte('j'): 0.24,
		byte('k'): 0.41,
		byte('l'): 2.92,
		byte('m'): 2.76,
		byte('n'): 5.44,
		byte('o'): 6.00,
		byte('p'): 1.95,
		byte('q'): 0.24,
		byte('r'): 4.95,
		byte('s'): 5.68,
		byte('t'): 8.03,
		byte('u'): 2.43,
		byte('v'): 0.97,
		byte('w'): 1.38,
		byte('x'): 0.24,
		byte('y'): 1.30,
		byte('z'): 0.03,
	}
	return
}
