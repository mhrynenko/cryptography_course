package sha256

import (
	"math/bits"
)

func SmallSigma1(x uint32) uint32 {
	return bits.RotateLeft32(x, -17) ^ bits.RotateLeft32(x, -19) ^ (x >> 10)
}

func SmallSigma0(x uint32) uint32 {
	return bits.RotateLeft32(x, -7) ^ bits.RotateLeft32(x, -18) ^ (x >> 3)
}

func BigSigma0(x uint32) uint32 {
	return bits.RotateLeft32(x, -2) ^ bits.RotateLeft32(x, -13) ^ bits.RotateLeft32(x, -22)
}

func BigSigma1(x uint32) uint32 {
	return bits.RotateLeft32(x, -6) ^ bits.RotateLeft32(x, -11) ^ bits.RotateLeft32(x, -25)
}

func Ch(x, y, z uint32) uint32 {
	return (x & y) ^ (^x & z)
}

func Maj(x, y, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}
