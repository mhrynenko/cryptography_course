package main

import (
	"crypto/rand"
	"fmt"
	"github.com/pkg/errors"
	"math/big"
	"strings"
	"time"
)

var bitSizes = []int64{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}

func getSpaceSize(size int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(2), big.NewInt(size), nil)
}

func generateRandomKey(size int64) (*big.Int, error) {
	key, err := rand.Int(rand.Reader, getSpaceSize(size))
	if err != nil {
		return nil, err
	}

	return key, nil
}

func simpleBruteForce(number *big.Int, size int64) time.Duration {
	start := time.Now()

	for i := big.NewInt(0); i.Cmp(getSpaceSize(size)) == -1; i.Add(i, big.NewInt(1)) {
		if number.Cmp(i) == 0 {
			break
		}
	}

	return time.Since(start)
}

func printDelimiter(delimiter string, amount int) {
	fmt.Println(strings.Repeat(delimiter, amount))
}

func main() {
	for _, size := range bitSizes {
		printDelimiter("=", 200)

		fmt.Printf("BIT SIZE %d\n", size)

		printDelimiter("-", 200)

		fmt.Printf("Space Size: %s\n", getSpaceSize(size).String())

		key, err := generateRandomKey(size)
		if err != nil {
			panic(errors.Wrap(err, "failed to generate random keys"))
		}

		printDelimiter("-", 200)

		fmt.Printf("Random key: %s\n", key.String())

		printDelimiter("-", 200)

		timeToFind := simpleBruteForce(key, size)
		fmt.Printf("Time in ms to find this key: %d\nRounded full time to find: %s\n", timeToFind.Milliseconds(), timeToFind.Round(time.Millisecond).String())
	}
}
