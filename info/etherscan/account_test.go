package etherscan

import (
	"math/big"
	"os"
	"testing"
)

func TestEtherscan_AccountBalance(t *testing.T) {
	ETHERSCAN_API_KEY := os.Getenv("ETHERSCAN_API_KEY")
	etherscan := initClient(Mainnet, ETHERSCAN_API_KEY)

	balance, err := etherscan.AccountBalance("0x0000000000000000000000000000000000000000")
	if err != nil {
		t.Fatalf("failed AccountBalance call")
	}

	if balance.Int().Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("failed AccountBalance compare result is not 1")
	}

	balance, err = etherscan.TokenBalance(TOKEN_CONTRACT_ADDRESS, WALLET_1)
	if err != nil {
		t.Fatalf("failed TokenBalance call")
	}

	if balance.Int().Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("failed TokenBalance compare result is not 1")
	}
}
