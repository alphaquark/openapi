package etherscan

import (
	"math"
	"math/big"
	"os"
	"strconv"

	"github.com/gravision/alphaquark-upbit-api/logger"
)

const (
	TOKEN_CONTRACT_ADDRESS = "0x2a9bDCFF37aB68B95A53435ADFd8892e86084F93"
	TOKEN_DECIMAL          = 18

	WALLET_1 = "0xdbdadf8dc466c981a3757df9510e3ac3e9289250"
	WALLET_2 = "0x53dd161458928f873b9dd63f088f1afb9a6fa489"
	WALLET_3 = "0x88f684a43beb8125a821cd24a9d04c712e12d974"

	TOTAL_SUPPLY = 30000000
)

// Serve is return AQT Circulating suuply amount
func Serve() (uint64, error) {
	decimal := uint64(math.Pow10(TOKEN_DECIMAL)) // 10 ** Token decimal
	wallet_balances := []uint64{0, 0, 0}
	wallet_addrs := []string{
		WALLET_1, WALLET_2, WALLET_3,
	}

	ETHERSCAN_API_KEY := os.Getenv("ETHERSCAN_API_KEY")
	etherscan := initClient(Mainnet, ETHERSCAN_API_KEY)

	for i := 0; i < len(wallet_addrs); i++ {
		bal, err := etherscan.TokenBalance(TOKEN_CONTRACT_ADDRESS, wallet_addrs[i])
		if err != nil {
			logger.ErrorField("error", err.Error(), "failed get token balance")
			return 0, nil
		}

		dv := new(big.Int).Div(bal.Int(), new(big.Int).SetUint64(decimal))
		wallet_balances[i] = dv.Uint64()
		logger.Info(strconv.Itoa(i) + " wallet token balance : " + strconv.Itoa(int(wallet_balances[i])))
	}

	circulatingSupply := uint64(TOTAL_SUPPLY)
	for i := 0; i < len(wallet_addrs); i++ {
		circulatingSupply = circulatingSupply - wallet_balances[i]
	}

	logger.Debug("Circulating supply : " + strconv.Itoa(int(circulatingSupply)))
	return circulatingSupply, nil
}
