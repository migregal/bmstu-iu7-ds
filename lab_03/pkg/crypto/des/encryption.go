package des

import (
	"github.com/migregal/bmstu-iu7-ds/lab-03/pkg/crypto/pkcs5"
	"github.com/migregal/bmstu-iu7-ds/lab-03/pkg/types/bitvec"
)

const blockSize = 8

func Cipher(key string, data []byte) ([]byte, error) {
	keys := GenerateKeys(key)

	data = pkcs5.PKCS5Padding(data, blockSize)

	var (
		res    bitvec.BitVec
		chunks = getChunks(data, blockSize)
	)

	for _, chunk := range chunks {
		binaryIP := ip(stringToBinSlice(chunk))
		l16, r16, err := rounds(binaryIP, keys, false)
		if err != nil {
			return nil, err
		}

		lr16 := append(r16, l16...)
		res += bitvec.Join(ipl1(lr16), "")
	}
	binRes := []byte(res)

	output := []byte{}
	for i := 0; i < len(binRes); i += blockSize {
		b, _ := bitvec.ToUint(bitvec.BitVec(binRes[i:i+blockSize]), 2, 64)
		output = append(output, byte(b))
	}

	return output, nil
}

func Decipher(key string, data []byte) ([]byte, error) {
	keys := GenerateKeys(key)

	input := []byte{}
	for _, d := range data {
		for i := blockSize - 1; i >= 0; i-- {
			input = append(input, byte('0'+((d>>i)&1)))
		}
	}

	var (
		res    string
		chunks = getChunks(input, 64)
	)

	for _, chunk := range chunks {
		binaryIP := ip(bitvec.Split(chunk, ""))
		l16, r16, err := rounds(binaryIP, keys, true)
		if err != nil {
			return nil, err
		}

		lr16 := append(r16, l16...)

		data, err := toString(bitvec.Join(ipl1(lr16), ""))
		if err != nil {
			return nil, err
		}

		res += data
	}

	return pkcs5.PKCS5UnPadding([]byte(res)), nil
}

func ip(s []bitvec.BitVec) []bitvec.BitVec {
	return []bitvec.BitVec{s[57], s[49], s[41], s[33], s[25], s[17], s[9], s[1],
		s[59], s[51], s[43], s[35], s[27], s[19], s[11], s[3],
		s[61], s[53], s[45], s[37], s[29], s[21], s[13], s[5],
		s[63], s[55], s[47], s[39], s[31], s[23], s[15], s[7],
		s[56], s[48], s[40], s[32], s[24], s[16], s[8], s[0],
		s[58], s[50], s[42], s[34], s[26], s[18], s[10], s[2],
		s[60], s[52], s[44], s[36], s[28], s[20], s[12], s[4],
		s[62], s[54], s[46], s[38], s[30], s[22], s[14], s[6]}
}

func ipl1(s []bitvec.BitVec) []bitvec.BitVec {
	return []bitvec.BitVec{s[39], s[7], s[47], s[15], s[55], s[23], s[63], s[31],
		s[38], s[6], s[46], s[14], s[54], s[22], s[62], s[30],
		s[37], s[5], s[45], s[13], s[53], s[21], s[61], s[29],
		s[36], s[4], s[44], s[12], s[52], s[20], s[60], s[28],
		s[35], s[3], s[43], s[11], s[51], s[19], s[59], s[27],
		s[34], s[2], s[42], s[10], s[50], s[18], s[58], s[26],
		s[33], s[1], s[41], s[9], s[49], s[17], s[57], s[25],
		s[32], s[0], s[40], s[8], s[48], s[16], s[56], s[24]}
}

func getChunks(s []byte, chunkSize int) []bitvec.BitVec {
	chunk := make([]byte, chunkSize)
	if chunkSize >= len(s) {
		return []bitvec.BitVec{bitvec.BitVec(s)}
	}

	var chunks []bitvec.BitVec
	len := 0
	for _, r := range s {
		chunk[len], len = r, len+1
		if len == chunkSize {
			chunks = append(chunks, bitvec.BitVec(chunk))
			len = 0
		}
	}

	return chunks
}
