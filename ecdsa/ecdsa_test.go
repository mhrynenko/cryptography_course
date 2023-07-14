package ecdsa

import (
	"crypto/elliptic"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var inputs = []string{
	"Hello World!",
	"hello world",
	"sha256",
	"Some example of long message",
	"Some example of loooooooooooooooooooonger message",
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	"Agreed joy vanity regret met may ladies oppose who. Mile fail as left as hard eyes. Meet made call in mean four year it to. Prospect so branched wondered sensible of up. For gay consisted resolving pronounce sportsman saw discovery not. Northward or household as conveying we earnestly believing. No in up contrasted discretion inhabiting excellence. Entreaties we collecting unpleasant at everything conviction.\n\nThat know ask case sex ham dear her spot. Weddings followed the all marianne nor whatever settling. Perhaps six prudent several her had offence. Did had way law dinner square tastes. Recommend concealed yet her procuring see consulted depending. Adieus hunted end plenty are his she afraid. Resources agreement contained propriety applauded neglected use yet.\n\nRespect forming clothes do in he. Course so piqued no an by appear. Themselves reasonable pianoforte so motionless he as difficulty be. Abode way begin ham there power whole. Do unpleasing indulgence impossible to conviction. Suppose neither evident welcome it at do civilly uncivil. Sing tall much you get nor.\n\nBarton waited twenty always repair in within we do. An delighted offending curiosity my is dashwoods at. Boy prosperous increasing surrounded companions her nor advantages sufficient put. John on time down give meet help as of. Him waiting and correct believe now cottage she another. Vexed six shy yet along learn maids her tiled. Through studied shyness evening bed him winding present. Become excuse hardly on my thirty it wanted.\n\nStyle too own civil out along. Perfectly offending attempted add arranging age gentleman concluded. Get who uncommonly our expression ten increasing considered occasional travelling. Ever read tell year give may men call its. Piqued son turned fat income played end wicket. To do noisy downs round an happy books.\n\nOught these are balls place mrs their times add she. Taken no great widow spoke of it small. Genius use except son esteem merely her limits. Sons park by do make on. It do oh cottage offered cottage in written. Especially of dissimilar up attachment themselves by interested boisterous. Linen mrs seems men table. Jennings dashwood to quitting marriage bachelor in. On as conviction in of appearance apartments boisterous.\n\nFriendship contrasted solicitude insipidity in introduced literature it. He seemed denote except as oppose do spring my. Between any may mention evening age shortly can ability regular. He shortly sixteen of colonel colonel evening cordial to. Although jointure an my of mistress servants am weddings. Age why the therefore education unfeeling for arranging. Above again money own scale maids ham least led. Returned settling produced strongly ecstatic use yourself way. Repulsive extremity enjoyment she perceived nor.\n\nOn on produce colonel pointed. Just four sold need over how any. In to september suspicion determine he prevailed admitting. On adapted an as affixed limited on. Giving cousin warmly things no spring mr be abroad. Relation breeding be as repeated strictly followed margaret. One gravity son brought shyness waiting regular led ham.\n\nYour it to gave life whom as. Favourable dissimilar resolution led for and had. At play much to time four many. Moonlight of situation so if necessary therefore attending abilities. Calling looking enquire up me to in removal. Park fat she nor does play deal our. Procured sex material his offering humanity laughing moderate can. Unreserved had she nay dissimilar admiration interested. Departure performed exquisite rapturous so ye me resources.\n\nConsider now provided laughter boy landlord dashwood. Often voice and the spoke. No shewing fertile village equally prepare up females as an. That do an case an what plan hour of paid. Invitation is unpleasant astonished preference attachment friendship on. Did sentiments increasing particular nay. Mr he recurred received prospect in. Wishing cheered parlors adapted am at amongst matters.",
}

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
