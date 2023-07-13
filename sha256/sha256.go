package sha256

import (
	"encoding/binary"
	"math"
)

var K = [64]uint32{
	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

var H = [8]uint32{}

func initH() {
	H = [8]uint32{
		0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
	}
}

func paddingMessage(data []byte) []byte {
	messageLength := len(data)

	var res = append(data, 0x80)

	zerosAmount := (56 - (messageLength + 1)) % 64
	if zerosAmount < 0 {
		zerosAmount += 64
	}

	var zeroes = make([]byte, zerosAmount)
	res = append(res, zeroes...)

	var length = make([]byte, 8)
	binary.BigEndian.PutUint64(length, uint64(messageLength*8))
	res = append(res, length...)

	return res
}

func parsePaddedMessage(message []byte) [][]byte {
	var messageLength = len(message)

	if messageLength == 0 || messageLength >= math.MaxInt64 || messageLength%64 != 0 {
		panic("message length is invalid")
	}

	var blocksAmount = messageLength / 64
	var blocks = make([][]byte, blocksAmount)

	for i := 0; i < blocksAmount; i++ {
		block := make([]byte, 64)
		copy(block, message[i*64:])
		blocks[i] = block
	}

	return blocks
}

func prepareMessageSchedule(block []byte) [64]uint32 {
	var schedule [64]uint32

	//0 <= t <= 15
	for t := 0; t < 16; t++ {
		var j = t * 4
		schedule[t] = binary.BigEndian.Uint32(block[j:])
	}

	//16 <= t <= 63
	for t := 16; t < 64; t++ {
		schedule[t] = SmallSigma1(schedule[t-2]) + schedule[t-7] + SmallSigma0(schedule[t-15]) + schedule[t-16]
	}

	return schedule
}

func compression(schedule [64]uint32) {
	a := H[0]
	b := H[1]
	c := H[2]
	d := H[3]
	e := H[4]
	f := H[5]
	g := H[6]
	h := H[7]

	for t := 0; t < 64; t++ {
		var T1 = h + BigSigma1(e) + Ch(e, f, g) + K[t] + schedule[t]
		var T2 = BigSigma0(a) + Maj(a, b, c)

		h = g
		g = f
		f = e
		e = d + T1
		d = c
		c = b
		b = a
		a = T1 + T2
	}

	H[0] = a + H[0]
	H[1] = b + H[1]
	H[2] = c + H[2]
	H[3] = d + H[3]
	H[4] = e + H[4]
	H[5] = f + H[5]
	H[6] = g + H[6]
	H[7] = h + H[7]
}

func getBytesResult() [32]byte {
	var result [32]byte

	for i := 0; i < 32; i += 4 {
		var bytes [4]byte
		binary.BigEndian.PutUint32(bytes[:], H[i/4])
		copy(result[i:], bytes[:])
	}

	return result
}

func Compute(input []byte) [32]byte {
	initH()

	padded := paddingMessage(input)
	blocks := parsePaddedMessage(padded)

	for i := 0; i < len(blocks); i++ {
		schedule := prepareMessageSchedule(blocks[i])
		compression(schedule)
	}

	return getBytesResult()
}
