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
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/core"
)

// InitGenesis performs genesis initialization for the evm module. It returns
// no validator updates.
func (p *plugin) InitGenesis(ctx sdk.Context, ethGen *core.Genesis) {
	p.Prepare(ctx)
	p.SetChainConfig(ethGen.Config)
}

// ExportGenesis returns the exported genesis state as raw bytes for the evm
// module.
func (p *plugin) ExportGenesis(ctx sdk.Context, ethGen *core.Genesis) {
	p.Prepare(ctx)
	ethGen.Config = p.ChainConfig()
}
