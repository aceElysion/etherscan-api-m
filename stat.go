/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

// EtherTotalSupply gets total supply of ether
func (c *Client) EtherTotalSupply() (totalSupply *BigInt, err error) {
	err = c.Call("stats", "ethsupply", nil, &totalSupply)
	return
}

// EtherLatestPrice gets the latest ether price, in BTC and USD
func (c *Client) EtherLatestPrice() (price LatestPrice, err error) {
	err = c.Call("stats", "ethprice", nil, &price)
	return
}

// TokenTotalSupply gets total supply of token on specified contract address
func (c *Client) TokenTotalSupply(contractAddress string) (totalSupply *BigInt, err error) {
	param := M{
		"contractaddress": contractAddress,
	}

	err = c.Call("stats", "tokensupply", param, &totalSupply)
	return
}
