package converter

import (
	"math/big"
	"strings"
	"testing"
)

type data struct {
	hex          string
	size         int
	littleEndian *big.Int
	bigEndian    *big.Int
}

func TestVectors(t *testing.T) {
	var expect = make([]data, 0)
	var val *big.Int

	val, _ = new(big.Int).SetString("115339776388732929035197660848497720713218148788040405586178452820382218977280", 10)
	expect = append(expect, data{"ff00000000000000000000000000000000000000000000000000000000000000", 32, big.NewInt(255), val})

	val, _ = new(big.Int).SetString("77193548260167611359494267807458109956502771454495792280332974934474558013440", 10)
	expect = append(expect, data{"aaaa000000000000000000000000000000000000000000000000000000000000", 32, big.NewInt(43690), val})

	expect = append(expect, data{"FFFFFFFF", 4, big.NewInt(4294967295), big.NewInt(4294967295)})

	val, _ = new(big.Int).SetString("979114576324830475023518166296835358668716483481922294110218890578706788723335115795775136189060210944584475044786808910613350098299181506809283832360654948074334665509728123444088990750984735919776315636114949587227798911935355699067813766573049953903257414411690972566828795693861196044813729172123152193769005290826676049325224028303369631812105737593272002471587527915367835952474124875982077070337970837392460768423348044782340688207323630599527945406427226264695390995320400314062984891593411332752703846859640346323687201762934524222363836094053204269986087043470117703336873406636573235808683444836432453459818599293667760149123595668832133083221407128310342064668595954073131257995767262426534143159642539179485013975461689493733866106312135829807129162654188209922755829012304582671671519678313609748646814745057724363462189490278183457296449014163077506949636570237334109910914728582640301294341605533983878368789071427913184794906223657920124153256147359625549743656058746335124502376663710766611046750739680547042183503568549468592703882095207981161012224965829605768300297615939788368703353944514111011011184191740295491255291545096680705534063721012625490368756140460791685877738232879406346334603566914069127957053440", 10)
	expect = append(expect, data{"F000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 512, big.NewInt(240), val})

	expect = append(expect, data{"13f4e0555a", 5, big.NewInt(387987862547), big.NewInt(85712721242)})

	for i := range expect {
		testVector(t, expect[i])
	}
}

func testVector(t *testing.T, value data) {
	bigEndian, length, err := HexToBigEndian(value.hex)
	if err != nil {
		t.Errorf("HexToBigEndian: doesn't excpected error: `%s`", err.Error())
	}
	if length != value.size {
		t.Errorf("HexToBigEndian: big endian length `%d` does not match expected `%d`", length, value.size)
	}
	if value.bigEndian.Cmp(bigEndian) != 0 {
		t.Errorf("HexToBigEndian: big endian value `%s` does not match expected `%s`", bigEndian.String(), value.bigEndian.String())
	}

	littleEndian, length, err := HexToLittleEndian(value.hex)
	if err != nil {
		t.Errorf("HexToLittleEndian: doesn't excpected error: `%s`", err.Error())
	}
	if length != value.size {
		t.Errorf("HexToLittleEndian: little endian length `%d` does not match expected `%d`", length, value.size)
	}
	if value.littleEndian.Cmp(littleEndian) != 0 {
		t.Errorf("HexTolitleEndian: little endian value `%s` does not match expected `%s`", littleEndian.String(), value.littleEndian.String())
	}

	hex := BigEndianToHex(bigEndian, value.size*2)
	if strings.ToLower(hex) != strings.ToLower(value.hex) {
		t.Errorf("BigEndianToHex: big endian hex `%s` does not match expected `%s`", hex, value.hex)
	}

	hex = LittleEndianToHex(littleEndian, value.size*2)
	if strings.ToLower(hex) != strings.ToLower(value.hex) {
		t.Errorf("LittleEndianToHex: little endian hex `%s` does not match expected `%s`", hex, value.hex)
	}
}
