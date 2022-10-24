package pcks115

import "io"

// nonZeroRandomBytes fills the given slice with non-zero random octets.
func nonZeroRandomBytes(s []byte, random io.Reader) (err error) {
	_, err = io.ReadFull(random, s)
	if err != nil {
		return
	}

	for i := 0; i < len(s); i++ {
		for s[i] == 0 {
			_, err = io.ReadFull(random, s[i:i+1])
			if err != nil {
				return
			}
			s[i] ^= 0x42
		}
	}

	return
}
