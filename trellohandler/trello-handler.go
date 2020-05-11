package trellohandler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

var (
	keyFilename        = "key.dat"
	tokenFilename      = "token.dat"
	passphraseFilename = "pass.dat"
)

func PersistCredentials(key string, token string, passphrase string) {
	encryptedKey := encrypt(key, "password")
	encryptedToken := encrypt(token, "password")

	fileKey, _ := os.Create(keyFilename)
	defer fileKey.Close()
	fileKey.Write(encryptedKey)

	fileToken, _ := os.Create(tokenFilename)
	defer fileToken.Close()
	fileToken.Write(encryptedToken)
}

func GetCredentials() (bool, string, string) {

	if !fileExists(keyFilename) || !fileExists(tokenFilename) {
		return false, "", ""
	}

	dataKey, err := ioutil.ReadFile(keyFilename)
	if err != nil {

	}
	key := decrypt(dataKey, "password")

	dataToken, _ := ioutil.ReadFile(tokenFilename)
	token := decrypt(dataToken, "password")

	return true, string(key), string(token)
}

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}

	return true
}

func GetBoards() {

}

func encrypt(data string, passphrase string) []byte {
	block, err := aes.NewCipher([]byte(createHash(passphrase)))
	if err != nil {
		panic(err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
