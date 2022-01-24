package etherscan

import "testing"

func TestEtherscan_craftURL(t *testing.T) {
	etherscan := initClient(Mainnet, "any_key")

	const expected = `https://api.etherscan.io/api?action=craftURL&apikey=any_key&four=d&four=e&four=f&module=testing&one=1&three=1&three=2&three=3&two=2`
	output := etherscan.craftURL("testing", "craftURL", map[string]interface{}{
		"one":   1,
		"two":   "2",
		"three": []int{1, 2, 3},
		"four":  []string{"d", "e", "f"},
	})

	if output != expected {
		t.Fatalf("output != expected, got %s, want %s", output, expected)
	}
}
