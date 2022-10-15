package des

import (
	"fmt"
	"strconv"

	"github.com/migregal/bmstu-iu7-ds/lab-03/pkg/types/bitvec"
)

func CompleteKey(key string) string {
	for len([]byte(key))%8 > 0 {
		key += "."
	}
	return key
}

func stringToBinary(s bitvec.BitVec) (res bitvec.BitVec) {
	for _, c := range []byte(s) {
		res = bitvec.BitVec(fmt.Sprintf("%s%.8b", res, c))
	}

	return res
}

func stringToBinSlice(s bitvec.BitVec) []bitvec.BitVec {
	tmp := bitvec.Split(stringToBinary(s), "")
	res := make([]bitvec.BitVec, len(tmp))

	for i := range tmp {
		res[i] = bitvec.BitVec(tmp[i])
	}

	return res
}

func toString(s bitvec.BitVec) (string, error) {
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
