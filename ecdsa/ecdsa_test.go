package ecdsa

import (
	"crypto/elliptic"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Vector struct {
	curve elliptic.Curve
	m     string
	k     *big.Int
	r     *big.Int
	s     *big.Int
}

var vectors = []Vector{
	{
		curve: elliptic.P256(),
		m:     "sample",
		k:     new(big.Int).SetBytes(common.Hex2Bytes("A6E3C57DD01ABE90086538398355DD4C3B17AA873382B0F24D6129493D8AAD60")),
		r:     new(big.Int).SetBytes(common.Hex2Bytes("EFD48B2AACB6A8FD1140DD9CD45E81D69D2C877B56AAF991C34D0EA84EAF3716")),
		s:     new(big.Int).SetBytes(common.Hex2Bytes("F7CB1C942D657C41D436C7A1B6E29F65F3E900DBB9AFF4064DC4AB2F843ACDA8")),
	},
	{
		curve: elliptic.P256(),
		m:     "test",
		k:     new(big.Int).SetBytes(common.Hex2Bytes("D16B6AE827F17175E040871A1C7EC3500192C4C92677336EC2537ACAEE0008E0")),
		r:     new(big.Int).SetBytes(common.Hex2Bytes("F1ABB023518351CD71D881567B1EA663ED3EFCF6C5132B354F28D3B0B7D38367")),
		s:     new(big.Int).SetBytes(common.Hex2Bytes("019F4113742A2B14BD25926B49C649155F267E60D3814B4C0CC84250E46F0083")),
	},
}

func TestECDSA(t *testing.T) {
	key := PrivateKey{
		D: new(big.Int).SetBytes(common.Hex2Bytes("C9AFA9D845BA75166B5C215767B1D6934E50C3DB36E89B127B8A622B120F6721")),
		PK: PublicKey{
			X: new(big.Int).SetBytes(common.Hex2Bytes("60FED4BA255A9D31C961EB74C6356D68C049B8923B61FA6CE669622E60F29FB6")),
			Y: new(big.Int).SetBytes(common.Hex2Bytes("7903FE1008B8BC99A41AE9E95628BC64F2F1B20C2D7E9F5177A3C294D4462299")),
		},
	}

	for i, vector := range vectors {
		msg := []byte(vector.m)

		sig, err := Sign(vector.curve, msg, key.D, vector.k)
		if err != nil {
			t.Errorf("ecdsa.Sign: unexcpected error `%s` for %d vector", err.Error(), i)
		}

		if sig.S.Cmp(vector.s) != 0 || sig.R.Cmp(vector.r) != 0 {
			t.Errorf("ecdsa.Sign: signature is not the same for %d vector", i)
		}

		isVerified, err := Verify(vector.curve, msg, sig.R, sig.S, key.PK)
		if err != nil {
			t.Errorf("ecdsa.Verify: unexcpected error `%s` for %d vector", err.Error(), i)
		}

		if !isVerified {
			t.Errorf("ecdsa.Verify: signature is not verified for %d vector", i)
		}
	}

	generated, err := GeneratePrivateKey(crypto.S256())
	if err != nil {
		t.Errorf("ecdsa.GeneratePrivateKey: unexcpected error `%s`", err.Error())
	}

	sig, err := Sign(crypto.S256(), []byte("Hello world!"), generated.D, nil)
	if err != nil {
		t.Errorf("ecdsa.Sign: unexcpected error `%s`", err.Error())
	}

	isVerified, err := Verify(crypto.S256(), []byte("Hello world!"), sig.R, sig.S, generated.PK)
	if err != nil {
		t.Errorf("ecdsa.Verify: unexcpected error `%s`", err.Error())
	}

	if !isVerified {
		t.Errorf("ecdsa.Verify: signature is not verified")
	}
}
