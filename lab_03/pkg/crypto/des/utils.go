package des

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/migregal/bmstu-iu7-ds/lab-03/pkg/crypto/pkcs5"
)

func CompleteKey(key string) string {
	return string(pkcs5.PKCS5Padding([]byte(key), 8))
}

func stringToBinary(s string) (res string) {
	for _, c := range []byte(s) {
		res = fmt.Sprintf("%s%.8b", res, c)
	}

	return res
}

func stringToBinSlice(s string) []string {
	return strings.Split(stringToBinary(s), "")
}

func soString(s string) (string, error) {
	arr := make([]string, blockSize)

	j := -1
	for i, c := range []byte(s) {
		if i%blockSize == 0 {
			j++
		}

		arr[j] += string(c)
	}

	out := make([]byte, len(arr))
	for i, a := range arr {
		tmp, err := strconv.ParseUint(a, 2, 0)
		if err != nil {
			return "", err
		}

		out[i] = byte(tmp)
	}

	return string(out), nil
}
