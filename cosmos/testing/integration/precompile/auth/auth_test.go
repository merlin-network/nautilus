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

package auth_test

import (
	"testing"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/auth"
	"pkg.berachain.dev/polaris/cosmos/testing/integration"
	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCosmosPrecompiles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/integration/precompile/auth")
}

var (
	tf             *integration.TestFixture
	authPrecompile *bindings.AuthModule
)

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	tf = integration.NewTestFixture(GinkgoT())
	authPrecompile, _ = bindings.NewAuthModule(
		common.HexToAddress("0xBDF49C3C3882102fc017FFb661108c63a836D065"), tf.EthClient)
	return nil
}, func(data []byte) {})

var _ = Describe("Auth", func() {
	BeforeEach(func() {
		_, err := tf.Network.WaitForHeight(1)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should call functions on the precompile directly", func() {
		acc, err := authPrecompile.GetAccountInfo0(nil, tf.Address("alice"))
		Expect(err).NotTo(HaveOccurred())
		Expect(acc.Addr).To(Equal(tf.Address("alice")))
	})
})
