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

package keeper

import (
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/block"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/configuration"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/gas"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/historical"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile/log"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/eth/core"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Compile-time interface assertion.
var _ core.BlackfuryHostChain = (*host)(nil)

// Host is the interface that must be implemented by the host.
// It includes core.BlackfuryHostChain and functions that are called in other packages.
type Host interface {
	core.BlackfuryHostChain
	GetAllPlugins() []plugins.Base
	Setup(
		storetypes.StoreKey,
		storetypes.StoreKey,
		state.AccountKeeper,
		func(height int64, prove bool) (sdk.Context, error),
	)
}

type host struct {
	// The various plugins that are are used to implement core.BlackfuryHostChain.
	bp  block.Plugin
	cp  configuration.Plugin
	gp  gas.Plugin
	hp  historical.Plugin
	pp  precompile.Plugin
	sp  state.Plugin
	txp txpool.Plugin

	pcs func() *ethprecompile.Injector
}

// Newhost creates new instances of the plugin host.
func NewHost(
	storeKey storetypes.StoreKey,
	sk block.StakingKeeper,
	ethTxMempool sdkmempool.Mempool,
	precompiles func() *ethprecompile.Injector,
) Host {
	// We setup the host with some Cosmos standard sauce.
	h := &host{}

	// Build the Plugins
	h.bp = block.NewPlugin(storeKey, sk)
	h.cp = configuration.NewPlugin(storeKey)
	h.gp = gas.NewPlugin()
	h.txp = txpool.NewPlugin(utils.MustGetAs[*mempool.EthTxPool](ethTxMempool))
	h.pcs = precompiles

	return h
}

// Setup sets up the precompile and state plugins with the given precompiles and keepers. It also
// sets the query context function for the block and state plugins (to support historical queries).
func (h *host) Setup(
	storeKey storetypes.StoreKey,
	_ storetypes.StoreKey,
	ak state.AccountKeeper,
	qc func(height int64, prove bool) (sdk.Context, error),
) {
	// Setup the state, precompile, historical, and txpool plugins
	h.sp = state.NewPlugin(ak, storeKey, log.NewFactory(h.pcs().GetPrecompiles()))
	h.pp = precompile.NewPlugin(h.pcs().GetPrecompiles(), h.sp)
	// TODO: re-enable historical plugin using ABCI listener.
	h.hp = historical.NewPlugin(h.cp, h.bp, nil, storeKey)
	h.txp.SetNonceRetriever(h.sp)

	// Set the query context function for the block and state plugins
	h.sp.SetQueryContextFn(qc)
	h.bp.SetQueryContextFn(qc)
}

// GetBlockPlugin returns the header plugin.
func (h *host) GetBlockPlugin() core.BlockPlugin {
	return h.bp
}

// GetConfigurationPlugin returns the configuration plugin.
func (h *host) GetConfigurationPlugin() core.ConfigurationPlugin {
	return h.cp
}

// GetGasPlugin returns the gas plugin.
func (h *host) GetGasPlugin() core.GasPlugin {
	return h.gp
}

func (h *host) GetHistoricalPlugin() core.HistoricalPlugin {
	return h.hp
}

// GetPrecompilePlugin returns the precompile plugin.
func (h *host) GetPrecompilePlugin() core.PrecompilePlugin {
	return h.pp
}

// GetStatePlugin returns the state plugin.
func (h *host) GetStatePlugin() core.StatePlugin {
	return h.sp
}

// GetTxPoolPlugin returns the txpool plugin.
func (h *host) GetTxPoolPlugin() core.TxPoolPlugin {
	return h.txp
}

// GetAllPlugins returns all the plugins.
func (h *host) GetAllPlugins() []plugins.Base {
	return []plugins.Base{h.bp, h.cp, h.gp, h.hp, h.pp, h.sp, h.txp}
}
