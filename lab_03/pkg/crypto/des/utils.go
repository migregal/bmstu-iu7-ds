package des

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func CompleteKey(key string) string {
	for len(key)%8 > 0 {
		key += "."
	}

	return key
}

func StringToBinary(s string) (res string) {
	for _, c := range s {
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
	for i, c := range s {
		if i%8 == 0 && i != 0 {
			j++
		}
		arr[j] += string(c)
	}

	for _, a := range arr {
		tmp, err := strconv.ParseUint(a, 2, 0)
		if err != nil {
			log.Fatal(err)
		}
		res += string(tmp)
	}

	return res
}
