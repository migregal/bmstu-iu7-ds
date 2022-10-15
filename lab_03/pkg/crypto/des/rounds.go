package des

import (
	"fmt"
	"strconv"
	"strings"
)

func rounds(
	binaryIP []string,
	keys []string,
	decrypt bool,
) (l16 []string, r16 []string, err error) {
	leftBlock, rightBlock := binaryIP[:32], binaryIP[32:]

	for i := range keys {
		var keyInt uint64
		if !decrypt {
			keyInt, err = strconv.ParseUint(keys[i], 2, 0)
		} else {
			keyInt, err = strconv.ParseUint(keys[len(keys)-i-1], 2, 0)
		}

		f, err := feistel(rightBlock, keyInt)

		fUint, err := strconv.ParseUint(strings.Join(f, ""), 2, 0)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse feistel result: %w", err)
		}

		leftBlockInt, err := strconv.ParseUint(strings.Join(leftBlock, ""), 2, 0)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse left block: %w", err)
		}

		fUint = fUint ^ leftBlockInt
		tmp := strings.Split(fmt.Sprintf("%.32b", fUint), "")

		leftBlock = rightBlock
		rightBlock = tmp
	}

	return leftBlock, rightBlock, nil
}
