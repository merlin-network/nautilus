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

package precompile

import (
	"context"
	"math/big"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state/events"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state/events/mock"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("plugin", func() {
	var p *plugin
	var e precompile.EVM
	var ctx sdk.Context

	BeforeEach(func() {
		ctx = testutil.NewContext()
		ctx = ctx.WithEventManager(
			events.NewManagerFrom(ctx.EventManager(), mock.NewPrecompileLogFactory()),
		)
		p = utils.MustGetAs[*plugin](NewPlugin(nil, &mockSP{ctx}))
		e = &mockEVM{nil, ctx}
	})

	It("should use correctly consume gas", func() {
		_, remainingGas, err := p.Run(e, &mockStateless{}, []byte{}, addr, new(big.Int), 30, false)
		Expect(err).ToNot(HaveOccurred())
		Expect(remainingGas).To(Equal(uint64(10)))
	})

	It("should error on insufficient gas", func() {
		_, _, err := p.Run(e, &mockStateless{}, []byte{}, addr, new(big.Int), 5, true)
		Expect(err.Error()).To(Equal("out of gas"))
	})

	It("should plug in custom gas configs", func() {
		Expect(p.KVGasConfig().DeleteCost).To(Equal(uint64(1000)))
		Expect(p.TransientKVGasConfig().DeleteCost).To(Equal(uint64(100)))

		p.SetKVGasConfig(storetypes.GasConfig{
			DeleteCost: 2,
		})
		Expect(p.KVGasConfig().DeleteCost).To(Equal(uint64(2)))
		p.SetTransientKVGasConfig(storetypes.GasConfig{
			DeleteCost: 3,
		})
		Expect(p.TransientKVGasConfig().DeleteCost).To(Equal(uint64(3)))
	})
})

// MOCKS BELOW.

type mockSP struct {
	ctx sdk.Context
}

func (msp *mockSP) SetGasConfig(kvg storetypes.GasConfig, tkvg storetypes.GasConfig) {
	msp.ctx = msp.ctx.WithKVGasConfig(kvg).WithTransientKVGasConfig(tkvg)
}

type mockEVM struct {
	precompile.EVM
	ctx sdk.Context
}

func (me *mockEVM) GetStateDB() vm.GethStateDB {
	return &mockSDB{nil, me.ctx}
}

type mockSDB struct {
	vm.PolarisStateDB
	ctx sdk.Context
}

func (ms *mockSDB) GetContext() context.Context {
	return ms.ctx
}

type mockStateless struct{}

var addr = common.BytesToAddress([]byte{1})

func (ms *mockStateless) RegistryKey() common.Address {
	return addr
}

func (ms *mockStateless) Run(
	ctx context.Context, _ precompile.EVM, _ []byte,
	_ common.Address, _ *big.Int, _ bool,
) ([]byte, error) {
	sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(10, "")
	return nil, nil
}

func (ms *mockStateless) RequiredGas(_ []byte) uint64 {
	return 10
}

func (ms *mockStateless) WithStateDB(vm.GethStateDB) vm.PrecompileContainer {
	return ms
}
