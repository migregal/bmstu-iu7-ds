package bitvec

import (
	"strconv"
	"strings"
)

type BitVec string

func Join(val []BitVec, sep BitVec) BitVec {
	tmp := make([]string, len(val))

	for i := range val {
		tmp[i] = string(val[i])
	}

	return BitVec(strings.Join(tmp, string(sep)))
}

func Split(val BitVec, sep BitVec) []BitVec {
	tmp := strings.Split(string(val), string(sep))

	res := make([]BitVec, len(tmp))
	for i := range tmp {
		res[i] = BitVec(tmp[i])
	}

	return res
}

func ToUint(val BitVec, base, bitsize int) (uint64, error) {
	return strconv.ParseUint(string(val), base, bitsize)
}
