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

package state_test

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/crypto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Genesis", func() {
	var (
		ctx  sdk.Context
		sp   state.Plugin
		code []byte
	)

	BeforeEach(func() {
		var ak state.AccountKeeper
		ctx, ak, _, _ = testutil.SetupMinimalKeepers()
		sp = state.NewPlugin(ak, testutil.EvmKey, nil)

		// Create account for alice.
		sp.Reset(ctx)
		code = []byte("code")
	})

	It("should init and export genesis", func() {
		genesis := new(core.Genesis)
		genesis.Alloc = make(core.GenesisAlloc)

		genesis.Alloc[alice] = core.GenesisAccount{
			Balance: big.NewInt(5e18),
			Storage: map[common.Hash]common.Hash{
				common.BytesToHash([]byte("key")): common.BytesToHash([]byte("value")),
			},
			Code: code,
		}
		// Call Init Genesis
		sp.InitGenesis(ctx, genesis)

		// Check that the code is set.
		sp.Reset(ctx)
		Expect(sp.GetCode(alice)).To(Equal(code))
		sp.Finalize()

		// Check that the code hash is set.
		sp.Reset(ctx)
		Expect(sp.GetCodeHash(alice)).To(Equal(crypto.Keccak256Hash(code)))
		sp.Finalize()
		Expect(sp.GetBalance(alice)).To(Equal(big.NewInt(5e18)))
		Expect(sp.GetCode(alice), code)

		// Very exported genesis is equal.
		var exportedGenesis core.Genesis
		sp.ExportGenesis(ctx, &exportedGenesis)
		Expect(exportedGenesis.Alloc).To(Equal(genesis.Alloc))
	})
})
