/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"strings"
)

// ContractABI gets contract abi for verified contract source codes
func (c *Client) ContractABI(address string) (abi string, err error) {
	param := M{
		"address": address,
	}

	err = c.Call("contract", "getabi", param, &abi)
	return
}

// ContractSource gets contract source code for verified contract source codes
func (c *Client) ContractSource(address string) (source []ContractSource, err error) {
	param := M{
		"address": address,
	}

	err = c.Call("contract", "getsourcecode", param, &source)
	return
}

// ContractCreation gets contract source code for verified contract source codes
func (c *Client) ContractCreation(addresses []string) (creations []ContractCreation, err error) {
	var sb strings.Builder
	for _, address := range addresses {
		sb.WriteString(address)
		sb.WriteString(",")
	}
	addrStr := sb.String()
	param := M{
		"contractaddresses": addrStr[:len(addrStr)-1],
	}

	err = c.Call("contract", "getcontractcreation", param, &creations)
	return
}
