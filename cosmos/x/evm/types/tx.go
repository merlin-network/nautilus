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

package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"

	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

// WrappedEthereumTransaction defines a Cosmos SDK message for Ethereum transactions.
var _ sdk.Msg = (*WrappedEthereumTransaction)(nil)

// NewFromTransaction sets the transaction data from an `coretypes.Transaction`.
func NewFromTransaction(tx *coretypes.Transaction) *WrappedEthereumTransaction {
	bz, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return &WrappedEthereumTransaction{
		Data: bz,
	}
}

// GetSigners returns the address(es) that must sign over the transaction.
func (etr *WrappedEthereumTransaction) GetSigners() []sdk.AccAddress {
	sender, err := etr.GetSender()
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{sdk.AccAddress(sender.Bytes())}
}

// AsTransaction extracts the transaction as an `coretypes.Transaction`.
func (etr *WrappedEthereumTransaction) AsTransaction() *coretypes.Transaction {
	tx := new(coretypes.Transaction)
	if err := tx.UnmarshalBinary(etr.Data); err != nil {
		return nil
	}
	return tx
}

// GetSignBytes returns the bytes to sign over for the transaction.
func (etr *WrappedEthereumTransaction) GetSignBytes() ([]byte, error) {
	tx := etr.AsTransaction()
	return coretypes.LatestSignerForChainID(tx.ChainId()).
		Hash(tx).Bytes(), nil
}

// GetSender extracts the sender address from the signature values using the latest signer for the given chainID.
func (etr *WrappedEthereumTransaction) GetSender() (common.Address, error) {
	tx := etr.AsTransaction()
	signer := coretypes.LatestSignerForChainID(tx.ChainId())
	return signer.Sender(tx)
}

// GetSender extracts the sender address from the signature values using the latest signer for the given chainID.
func (etr *WrappedEthereumTransaction) GetPubKey() ([]byte, error) {
	tx := etr.AsTransaction()
	signer := coretypes.LatestSignerForChainID(tx.ChainId())
	return signer.PubKey(tx)
}

// GetSender extracts the sender address from the signature values using the latest signer for the given chainID.
func (etr *WrappedEthereumTransaction) GetSignature() ([]byte, error) {
	tx := etr.AsTransaction()
	signer := coretypes.LatestSignerForChainID(tx.ChainId())
	return signer.Signature(tx)
}

// GetGas returns the gas limit of the transaction.
func (etr *WrappedEthereumTransaction) GetGas() uint64 {
	var tx *coretypes.Transaction
	if tx = etr.AsTransaction(); tx == nil {
		return 0
	}
	return tx.Gas()
}

// GetGasPrice returns the gas price of the transaction.
func (etr *WrappedEthereumTransaction) ValidateBasic() error {
	// Ensure the transaction is signed properly
	tx := etr.AsTransaction()
	if tx == nil {
		return errors.New("transaction data is invalid")
	}

	// Ensure the transaction does not have a negative value.
	if tx.Value().Sign() < 0 {
		return txpool.ErrNegativeValue
	}

	// Sanity check for extremely large numbers.
	if tx.GasFeeCap().BitLen() > 256 { //nolint:gomnd // 256 bits.
		return core.ErrFeeCapVeryHigh
	}

	// Sanity check for extremely large numbers.
	if tx.GasTipCap().BitLen() > 256 { //nolint:gomnd // 256 bits.
		return core.ErrTipVeryHigh
	}

	// Ensure gasFeeCap is greater than or equal to gasTipCap.
	if tx.GasFeeCapIntCmp(tx.GasTipCap()) < 0 {
		return core.ErrTipAboveFeeCap
	}

	return nil
}

// GetAsEthTx is a helper function to get an EthTx from a sdk.Tx.
func GetAsEthTx(tx sdk.Tx) *coretypes.Transaction {
	if len(tx.GetMsgs()) == 0 {
		return nil
	}
	etr, ok := utils.GetAs[*WrappedEthereumTransaction](tx.GetMsgs()[0])
	if !ok {
		return nil
	}
	return etr.AsTransaction()
}
