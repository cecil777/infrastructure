package cryptoex

// ICrypto is 加密接口
type ICrypto interface {
	Decrypt(ciphertext []byte) (plaintext []byte, err error)
	Encrypt(plaintext []byte) (ciphertext []byte, err error)
	Validate(ciphertext []byte, plaintext []byte) bool
}
