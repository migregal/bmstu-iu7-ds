package des

import "github.com/migregal/bmstu-iu7-ds/lab-03/pkg/types/bitvec"

var lsIndex = []int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

func GenerateKeys(str string) (keys []bitvec.BitVec) {
	leftBlock, rightBlock := pc1(stringToBinSlice(bitvec.BitVec(str)))

	for i := 0; i < 16; i++ {
		leftBlock = leftShift(leftBlock, lsIndex[i])
		rightBlock = leftShift(rightBlock, lsIndex[i])
		concatenateKey := append(leftBlock, rightBlock...)
		keys = append(keys, bitvec.Join(pc2(concatenateKey), ""))
	}

	return keys
}

func pc1(s []bitvec.BitVec) ([]bitvec.BitVec, []bitvec.BitVec) {
	return []bitvec.BitVec{
			s[56], s[48], s[40], s[32], s[24], s[16], s[8],
			s[0], s[57], s[49], s[41], s[33], s[25], s[17],
			s[9], s[1], s[58], s[50], s[42], s[34], s[26],
			s[18], s[10], s[2], s[59], s[51], s[43], s[35]},
		[]bitvec.BitVec{
			s[62], s[54], s[46], s[38], s[30], s[22], s[14],
			s[6], s[61], s[53], s[45], s[37], s[29], s[21],
			s[13], s[5], s[60], s[52], s[44], s[36], s[28],
			s[20], s[12], s[4], s[27], s[19], s[11], s[3]}
}

func pc2(s []bitvec.BitVec) []bitvec.BitVec {
	return []bitvec.BitVec{
		s[13], s[16], s[10], s[23], s[0], s[4], s[2], s[27],
		s[14], s[5], s[20], s[9], s[22], s[18], s[11], s[3],
		s[25], s[7], s[15], s[6], s[26], s[19], s[12], s[1],
		s[40], s[51], s[30], s[36], s[46], s[54], s[29], s[39],
		s[50], s[44], s[32], s[47], s[43], s[48], s[38], s[55],
		s[33], s[52], s[45], s[41], s[49], s[35], s[28], s[31]}
}

func leftShift(s []bitvec.BitVec, i int) []bitvec.BitVec {
	for j := 0; j < i%len(s); j++ {
		s = append([]bitvec.BitVec{s[len(s)-1]}, s[:len(s)-1]...)
	}

	return s
}
