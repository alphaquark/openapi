package etherscan

import "github.com/gravision/alphaquark-upbit-api/logger"

// AccountBalance gets ether balance for a single address
func (c *Client) AccountBalance(address string) (balance *BigInt, err error) {
	logger.Debug("Try to get account ethereum balance")

	param := map[string]interface{}{
		"tag":     "latest",
		"address": address,
	}
	balance = new(BigInt)
	err = c.call("account", "balance", param, balance)
	return
}

// TokenBalance get erc20-token account balance of address for contractAddress
func (c *Client) TokenBalance(contractAddress, address string) (balance *BigInt, err error) {
	logger.Debug("Try to get token balance")

	param := map[string]interface{}{
		"contractaddress": contractAddress,
		"address":         address,
		"tag":             "latest",
	}

	err = c.call("account", "tokenbalance", param, &balance)
	return
}
