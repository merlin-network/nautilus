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

package configuration

import (
	"encoding/json"

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/params"
)

// GetChainConfig is used to get the genesis info of the Ethereum chain.
func (p *plugin) ChainConfig() *params.ChainConfig {
	bz := p.paramsStore.Get([]byte{types.ChainConfigPrefix})
	if bz == nil {
		return nil
	}
	var chainConfig params.ChainConfig
	if err := json.Unmarshal(bz, &chainConfig); err != nil {
		panic(err)
	}
	return &chainConfig
}

// GetEthGenesis is used to get the genesis info of the Ethereum chain.
func (p *plugin) SetChainConfig(chainConfig *params.ChainConfig) {
	bz, err := json.Marshal(chainConfig)
	if err != nil {
		panic(err)
	}
	p.paramsStore.Set([]byte{types.ChainConfigPrefix}, bz)
}
