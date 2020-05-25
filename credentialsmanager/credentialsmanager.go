package credentialsmanager

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	color "trello/commandColors"
)

var (
	keyFilename             = "key.dat"
	tokenFilename           = "token.dat"
	passphraseFilename      = "pass.dat"
	authenticatedPassphrase string
)

func PersistCredentials(key string, token string, passphrase string) {
	encryptedKey := encrypt(key, passphrase)
	encryptedToken := encrypt(token, passphrase)

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

	var passphrase string
	var key string
	var token string

	if !fileExists(passphraseFilename) {

		if authenticatedPassphrase != "" {
			_ = readCredentials(&key, &token, authenticatedPassphrase)
			return true, key, token
		}

		var passwordSuccess bool = false
		fmt.Println(color.Yellow("@ Enter passphrase."))

		for !passwordSuccess {
			_, err := fmt.Scan(&passphrase)
			if err != nil {
				panic(err.Error())
			}

			success := readCredentials(&key, &token, passphrase)
			if !success {
				fmt.Println(color.Red("@ Wrong passphrase, try again."))
				continue
			}

			passwordSuccess = success
			authenticatedPassphrase = passphrase
		}
	} else {
		success := getPassphrase(&passphrase)
		if !success {
			fmt.Println(color.RedBold("@ Failed to get stored passphrase from file pass.dat"))
			os.Exit(1)
		}

		readCredentials(&key, &token, passphrase)
	}

	return true, key, token
}

func readCredentials(key *string, token *string, passphrase string) bool {

	dataKey, err := ioutil.ReadFile(keyFilename)
	if err != nil {
		panic(err.Error())
	}
	success, decryptedKey := decrypt(dataKey, passphrase)
	if !success {
		return false
	}

	dataToken, _ := ioutil.ReadFile(tokenFilename)
	success, decryptedToken := decrypt(dataToken, passphrase)

	*key = string(decryptedKey)
	*token = string(decryptedToken)

	return true
}

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}

	return true
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

func decrypt(data []byte, passphrase string) (bool, []byte) {
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
		return false, []byte{}
	}
	return true, plaintext
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func PersistPassphrase(passphrase string) {
	filePassphrase, _ := os.Create(passphraseFilename)
	defer filePassphrase.Close()
	filePassphrase.Write([]byte(b64.StdEncoding.EncodeToString([]byte(passphrase))))
}

func getPassphrase(passphrase *string) bool {
	if _, err := os.Stat(passphraseFilename); os.IsNotExist(err) {
		return false
	}

	dataPassphrase, err := ioutil.ReadFile(passphraseFilename)
	if err != nil {
		panic(err.Error())
	}

	decoded, _ := b64.StdEncoding.DecodeString(string(dataPassphrase))
	*passphrase = string(decoded)
	return true
}
