package des

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/migregal/bmstu-iu7-ds/lab-03/pkg/crypto/pkcs5"
)

func CompleteKey(key string) string {
	return string(pkcs5.PKCS5Padding([]byte(key), 8))
}

func StringToBinary(s string) (res string) {
	for _, c := range []byte(s) {
		res = fmt.Sprintf("%s%.8b", res, c)
	}

	return res
}

func StringToBinSlice(s string) []string {
	return strings.Split(StringToBinary(s), "")
}

func ToString(s string) (res string) {
	arr := make([]string, 8)
	j := 0
	for i, c := range []byte(s) {
		if i%8 == 0 && i != 0 {
			j++
		}
		arr[j] += string(c)
	}

	out := make([]byte, len(arr))
	for i, a := range arr {
		tmp, err := strconv.ParseUint(a, 2, 0)
		if err != nil {
			log.Fatal(err)
		}
		out[i] = byte(tmp)
	}

	return string(out)
}
