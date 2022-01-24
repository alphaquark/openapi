package etherscan

type Network string

const (
	Mainnet Network = "api" // Ethereum Mainnet for production

	BinanceSmartChain Network = "bsc" // Binance Smart Chain
)

// SubDomain returns the subdomain of  etherscan API
// via n provided.
func (n Network) SubDomain() (sub string) {
	return string(n)
}
