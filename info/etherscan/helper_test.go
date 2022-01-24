package etherscan

import (
	"math/big"
	"testing"
)

func TestBigInt(t *testing.T) {
	const ansStr = "255"
	var ans = big.NewInt(255)

	b := new(BigInt)
	err := b.UnmarshalText([]byte(ansStr))
	if err != nil {
		t.Fatalf("BigInt.UnmarshalText" + err.Error())
	}

	if b.Int().Cmp(ans) != 0 {
		t.Fatalf("BigInt.UnmarshalText not working, got %v, want %v", b.Int(), ans)
	}

	textBytes, err := b.MarshalText()
	if err != nil {
		t.Fatalf("BigInt.MarshalText" + err.Error())
	}

	if string(textBytes) != ansStr {
		t.Fatalf("BigInt.MarshalText not working, got %s, want %s", textBytes, ansStr)
	}
}
