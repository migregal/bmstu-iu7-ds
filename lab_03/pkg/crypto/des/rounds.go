package des

import (
	"fmt"

	"github.com/migregal/bmstu-iu7-ds/lab-03/pkg/types/bitvec"
)

func rounds(
	binaryIP []bitvec.BitVec,
	keys []bitvec.BitVec,
	decrypt bool,
) (l16 []bitvec.BitVec, r16 []bitvec.BitVec, err error) {
	leftBlock, rightBlock := binaryIP[:32], binaryIP[32:]

	for i := range keys {
		var keyInt uint64
		if !decrypt {
			keyInt, err = bitvec.ToUint(keys[i], 2, 0)
		} else {
			keyInt, err = bitvec.ToUint(keys[len(keys)-i-1], 2, 0)
		}

		f, err := feistel(rightBlock, keyInt)

		fUint, err := bitvec.ToUint(bitvec.Join(f, ""), 2, 0)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse feistel result: %w", err)
		}

		leftBlockInt, err := bitvec.ToUint(bitvec.Join(leftBlock, ""), 2, 0)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse left block: %w", err)
		}

		fUint = fUint ^ leftBlockInt
		tmp := bitvec.Split(bitvec.BitVec(fmt.Sprintf("%.32b", fUint)), "")

		leftBlock = rightBlock
		rightBlock = tmp
	}

	return leftBlock, rightBlock, nil
}
