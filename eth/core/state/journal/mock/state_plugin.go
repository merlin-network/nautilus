// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Blackchain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package mock

import (
	"math/big"

	"pkg.berachain.dev/polaris/eth/common"
)

//go:generate moq -out ./state_plugin.mock.go -skip-ensure -pkg mock ../ suicideStatePlugin

var (
	a1 = common.HexToAddress("0x1")
	a3 = common.HexToAddress("0x3")
	a4 = common.HexToAddress("0x4")
)

func NewSuicidesStatePluginMock() *suicideStatePluginMock {
	return &suicideStatePluginMock{
		GetCodeHashFunc: func(address common.Address) common.Hash {
			if address == a1 || address == a3 || address == a4 {
				return common.Hash{0x1}
			}
			return common.Hash{}
		},
		GetBalanceFunc: func(address common.Address) *big.Int {
			return new(big.Int)
		},
		SubBalanceFunc: func(address common.Address, amount *big.Int) {
			// no-op
		},
	}
}
