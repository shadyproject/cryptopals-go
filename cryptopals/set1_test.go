package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"runtime/debug"
	"testing"
)

func TestSet1Challenge1(t *testing.T) {
	inHex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	wantBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	bytes, err := hex.DecodeString(inHex)
	failIfError(t, err)
	gotBase64 := base64.StdEncoding.EncodeToString(bytes)

	if wantBase64 != gotBase64 {
		t.Fatalf("want: %v, got %v\n%s", wantBase64, gotBase64, string(debug.Stack()))
	}
}

func TestSet1Challenge2(t *testing.T) {
	inHex := "1c0111001f010100061a024b53535009181c"
	xorKeyHex := "686974207468652062756c6c277320657965"
	wantHex := "746865206b696420646f6e277420706c6179"

	inBytes, err := hex.DecodeString(inHex)
	failIfError(t, err)

	xorBytes, err := hex.DecodeString(xorKeyHex)
	failIfError(t, err)

	gotBytes := make([]byte, len(inBytes))

	for i := range inBytes {
		gotBytes[i] = inBytes[i] ^ xorBytes[i]
	}

	gotHex := hex.EncodeToString(gotBytes)

	if wantHex != gotHex {
		t.Fatalf("want %v, got %v\n%s", wantHex, gotHex, string(debug.Stack()))
	}
}

func failIfError(t *testing.T, err error) {
	if err != nil {
		fmt.Printf("FAIL: want no error, got error: %v\n%s", err, string(debug.Stack()))
		t.Fatalf("want no error, got error: %v\n%s", err, string(debug.Stack()))
	}
}
