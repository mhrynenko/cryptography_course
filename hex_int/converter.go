package converter

import (
	"math/big"
	"strconv"

	"github.com/pkg/errors"
)

type Endianness string

const (
	BIG    Endianness = "BigEndianness"
	LITTLE Endianness = "LittleEndianness"
)

const hexChars = "0123456789abcdef"

func HexToBigEndian(key string) (*big.Int, int, error) {
	bytes, err := hexToBytes(BIG, key)
	if err != nil {
		return nil, -1, errors.Wrap(err, "failed to convert hex to bytes")
	}

	return new(big.Int).SetBytes(bytes), len(bytes), nil
}

func HexToLittleEndian(key string) (*big.Int, int, error) {
	bytes, err := hexToBytes(LITTLE, key)
	if err != nil {
		return nil, -1, errors.Wrap(err, "failed to convert hex to bytes")
	}

	return new(big.Int).SetBytes(bytes), len(bytes), nil
}

func hexToBytes(endianness Endianness, key string) ([]byte, error) {
	if key[:2] == "0x" {
		key = key[2:]
	}

	if len(key)%2 != 0 {
		return nil, errors.New("hex string is odd length")
	}

	var result = make([]byte, len(key)/2)

	for i := 0; i < len(result); i++ {
		var j = i * 2
		value, err := strconv.ParseUint(key[j:j+2], 16, 8)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert hex value to byte")
		}

		switch endianness {
		case BIG:
			result[i] = byte(value)
		case LITTLE:
			result[len(result)-i-1] = byte(value)
		}
	}

	return result, nil
}

// LittleEndianToHex `size` is expected hexadecimal string length
func LittleEndianToHex(key *big.Int, size int) string {
	bytes := key.Bytes()

	hex := make([]byte, 0)

	for i := len(bytes) - 1; i >= 0; i-- {
		b := bytes[i]
		hex = append(hex, hexChars[b&0xff>>4], hexChars[b&0x0f])
	}

	for len(hex) < size {
		hex = append(hex, '0')
	}

	return string(hex)
}

// BigEndianToHex `size` is expected hexadecimal string length
func BigEndianToHex(key *big.Int, size int) string {
	bytes := key.Bytes()

	hex := make([]byte, 0)

	for i := 0; i < len(bytes); i++ {
		b := bytes[i]
		hex = append(hex, hexChars[b&0xff>>4], hexChars[b&0x0f])
	}

	for len(hex) < size {
		hex = append([]byte{'0'}, hex...)
	}

	return string(hex)
}
