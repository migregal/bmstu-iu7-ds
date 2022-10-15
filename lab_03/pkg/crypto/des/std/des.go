package stddes

import (
	"crypto/cipher"
	"crypto/des"

	"github.com/migregal/bmstu-iu7-ds/lab-03/pkg/crypto/pkcs5"
)

func CompleteKey(key string) string {
	for len(key)%8 > 0 {
		key += "."
	}

	return key
}

func Cipher(key, iv, plainText []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := pkcs5.PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)

	return cryted, nil
}

func Decipher(key, iv, cipherText []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = pkcs5.PKCS5UnPadding(origData)

	return origData, nil
}
