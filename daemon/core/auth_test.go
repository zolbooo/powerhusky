package core

import "testing"

func TestTokenAuth(t *testing.T) {
	var nonce = []byte("12345test")
	token := GenerateToken("secret", nonce)
	if !VerifyToken("secret", token) {
		t.Error("failed to verify valid token")
	}
	if VerifyToken("other-secret", token) {
		t.Error("verify succeeded with wrong token")
	}
}

func FuzzTokenAuth(f *testing.F) {
	testcases := [][]byte{[]byte("hello"), {0xDE, 0xAD, 0xBE, 0xEF}, []byte("sesame-seed")}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, nonce []byte) {
		token := GenerateToken("secret", nonce)
		if !VerifyToken("secret", token) {
			t.Error("failed to verify valid token")
		}
		if VerifyToken("other-secret", token) {
			t.Error("verify succeeded with wrong token")
		}
	})
}
