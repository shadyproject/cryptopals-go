package cryptopals

import (
	"encoding/hex"
	"os"
	"strings"
)

func loadData() (data []string) {
	d, err := os.ReadFile("./data/4.txt")

	if err != nil {
		println("Error reading data file: %v", err)
		panic("Error reading data file")
	}

	data = strings.Split(string(d), "\n")
	return
}

func findEncrypted() (plaintext string, ciphertext string, score float32) {

	lines := loadData()

	bestScore := float32(0)
	bestPlaintext := ""
	bestCipherText := ""

	for _, line := range lines {
		workingCipherText, err := hex.DecodeString(line)
		if err != nil {
			println("Error decoding hex string: %v", err)
			panic("Error decoding hex string")
		}

		for i := 0; i < 256; i++ {
			key := byte(i)
			workingPlainText := singleByteXor(workingCipherText, key)
			currentScore := stringScore(workingPlainText)

			if currentScore > bestScore {
				bestScore = currentScore
				bestPlaintext = string(workingPlainText)
				bestCipherText = hex.EncodeToString(workingCipherText)
			}
		}

	}

	return bestPlaintext, bestCipherText, bestScore
}
