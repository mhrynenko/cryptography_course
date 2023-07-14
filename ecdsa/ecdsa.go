package ecdsa

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"

	"github.com/mhrynenko/cryptography_course/sha256"
	"github.com/pkg/errors"
)

var (
	ErrZeroPublicKey         = errors.New("public key is zero")
	ErrPublicKeyIsNotOnCurve = errors.New("public key is not in elliptic curve")

	ErrNumberIsOutOfRange = errors.New("number is out of range")
)

type PublicKey struct {
	X *big.Int
	Y *big.Int
}

type PrivateKey struct {
	PK PublicKey
	D  *big.Int
}

type Signature struct {
	R *big.Int
	S *big.Int
}

// var Curve elliptic.Curve = secp256k1.S256()
//var Curve elliptic.Curve = elliptic.P256()

func GeneratePrivateKey(curve elliptic.Curve) (*PrivateKey, error) {
	d, err := rand.Int(rand.Reader, curve.Params().N)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate random big int")
	}
	d.Mod(d, curve.Params().N)

	pubX, pubY := curve.ScalarBaseMult(d.Bytes())

	if len(pubX.Bits()) == 0 && len(pubY.Bits()) == 0 {
		return nil, ErrZeroPublicKey
	}

	if !curve.IsOnCurve(pubX, pubY) {
		return nil, ErrPublicKeyIsNotOnCurve
	}

	return &PrivateKey{
		PK: PublicKey{X: pubX, Y: pubY},
		D:  d,
	}, nil
}

// Sign msg - message to sign, d - privateKey, k - more for test purposes, but also can be generated non programming way
func Sign(curve elliptic.Curve, msg []byte, d *big.Int, k *big.Int) (*Signature, error) {
	var err error = nil
	// H(m)
	msgHash := sha256.Compute(msg)
	h := new(big.Int).SetBytes(msgHash[:])
	h.Mod(h, curve.Params().N)

	r := big.NewInt(0)
	s := big.NewInt(0)

	//if r = 0 then go to start.
	//if s = 0 go to start.
	for len(r.Bits()) == 0 || len(s.Bits()) == 0 {
		//k <= random ∈ [1, n - 1].
		if k == nil {
			k, err = rand.Int(rand.Reader, curve.Params().N)
			if err != nil {
				return nil, errors.Wrap(err, "failed to generate random big int")
			}
		}

		//k x P = (x1, y1)
		x, _ := curve.ScalarBaseMult(k.Bytes())

		//r = x1 mod n.
		r = x.Mod(x, curve.Params().N)

		if len(r.Bits()) == 0 {
			continue
		}

		r = new(big.Int).Mod(x, curve.Params().N)

		//pow(k, -1)
		kInverse := new(big.Int).ModInverse(k, curve.Params().N)

		s = new(big.Int).Mul(d, r)
		s.Add(s, h)
		s.Mul(s, kInverse)
		s.Mod(s, curve.Params().N)

		if len(s.Bits()) == 0 {
			continue
		}
	}

	return &Signature{
		S: s,
		R: r,
	}, nil
}

func Verify(curve elliptic.Curve, msg []byte, r, s *big.Int, Q PublicKey) (bool, error) {
	if !curve.IsOnCurve(Q.X, Q.Y) {
		return false, ErrPublicKeyIsNotOnCurve
	}

	if !checkBigIntInRange(r, big.NewInt(0), curve.Params().N) || !checkBigIntInRange(s, big.NewInt(0), curve.Params().N) {
		return false, ErrNumberIsOutOfRange
	}

	// H(m)
	msgHash := sha256.Compute(msg)
	h := new(big.Int).SetBytes(msgHash[:])
	//h.Mod(h, curve.Params().N)

	//pow(s, -1)
	sInverse := new(big.Int).ModInverse(s, curve.Params().N)

	//u = (H(m)*s^-1)modN
	u := new(big.Int).Mul(h, sInverse)
	u.Mod(u, curve.Params().N)

	//v = (r*s^-1)modN
	v := new(big.Int).Mul(r, sInverse)
	v.Mod(v, curve.Params().N)

	//x0,y0 = u*G + v*Q
	x2, y2 := curve.ScalarBaseMult(u.Bytes())
	x1, y1 := curve.ScalarMult(Q.X, Q.Y, v.Bytes())
	x0, _ := curve.Add(x1, y1, x2, y2)

	return x0.Cmp(r) == 0, nil
}

func checkBigIntInRange(val, from, to *big.Int) bool {
	return !(val.Cmp(from) == -1 && val.Cmp(to) == 1)
}
