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
       "path"
        "path/filepath"
	color "trello/commandColors"
)

var (
	keyFilename             = "tre-key.dat"
	tokenFilename           = "tre-token.dat"
	passphraseFilename      = "tre-pass.dat"
	authenticatedPassphrase string
	key                     string
	token                   string
)

func PersistCredentials(inputKey string, inputToken string, passphrase string) {
	encryptedKey := encrypt(inputKey, passphrase)
	encryptedToken := encrypt(inputToken, passphrase)

	fileKey, _ := os.Create(buildFilePath(keyFilename))
	defer fileKey.Close()
	fileKey.Write(encryptedKey)

	fileToken, _ := os.Create(buildFilePath(tokenFilename))
	defer fileToken.Close()
	fileToken.Write(encryptedToken)

	key = inputKey
	token = inputToken
}

func GetCredentials() (bool, string, string) {

	if !fileExists(keyFilename) || !fileExists(tokenFilename) {
		return false, "", ""
	}

	var passphrase string

	if !fileExists(passphraseFilename) {
		if key != "" && token != "" {
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

		successRead := readCredentials(&key, &token, passphrase)
		if !successRead {
			fmt.Println(color.RedBold("@ Failed to read credentials with stored passphrase. Set credentials again with 'credentials' command."))
			os.Exit(1)
		}
	}

	return true, key, token
}

func readCredentials(key *string, token *string, passphrase string) bool {

	dataKey, err := ioutil.ReadFile(buildFilePath(keyFilename))
	if err != nil {
		panic(err.Error())
	}
	success, decryptedKey := decrypt(dataKey, passphrase)
	if !success {
		return false
	}

	dataToken, _ := ioutil.ReadFile(buildFilePath(tokenFilename))
	success, decryptedToken := decrypt(dataToken, passphrase)

	*key = string(decryptedKey)
	*token = string(decryptedToken)

	return true
}

func fileExists(fileName string) bool {
	if _, err := os.Stat(buildFilePath(fileName)); os.IsNotExist(err) {
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
	filePassphrase, _ := os.Create(buildFilePath(passphraseFilename))
	defer filePassphrase.Close()
	filePassphrase.Write([]byte(b64.StdEncoding.EncodeToString([]byte(passphrase))))
}

func getPassphrase(passphrase *string) bool {
	if _, err := os.Stat(buildFilePath(passphraseFilename)); os.IsNotExist(err) {
		return false
	}

	dataPassphrase, err := ioutil.ReadFile(buildFilePath(passphraseFilename))
	if err != nil {
		panic(err.Error())
	}

	decoded, _ := b64.StdEncoding.DecodeString(string(dataPassphrase))
	*passphrase = string(decoded)
	return true
}

func buildFilePath(fileName string) string {
        ex, err := os.Executable()
        if(err != nil) {
               panic(err.Error())
        }

       return path.Join(filepath.Dir(ex), fileName)
}
