package rsa

import (
	"fmt"
	"time"
)

func generateN(N uint64) uint64 {
	return uint64(time.Now().UnixNano()/1e6%(1000+int64(N)) + 13)
}

func generate2PrimeNumbers(N uint64) (p, q uint64, err error) {
	n := generateN(N)

	ps := sieve(n)

	l := len(ps)
	if l == 0 {
		return 0, 0, fmt.Errorf("l is 0, n=%d", n)
	}

	psm := make(map[uint64]struct{}, l)
	for _, v := range ps {
		psm[v] = struct{}{}
	}

	for k := range psm {
		if k < 5 {
			continue
		}
		p = k
	}
	for k := range psm {
		if k < 5 || k == p {
			continue
		}
		q = k
	}

	return p, q, nil
}
