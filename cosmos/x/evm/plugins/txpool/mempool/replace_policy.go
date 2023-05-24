// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
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

package mempool

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
)

// EthereumTxPoolConfig is the configuration for the Ethereum transaction pool.
type EthereumTxPoolConfig[C comparable] struct {
	PriceBump uint64 // Minimum price bump percentage to replace an already existing transaction (nonce)
}

// EthereumTxReplacement is called when a new transaction is added to the mempool and a transaction
// with the same nonce already exists. It returns true if the new transaction should replace the
// existing transaction.
// Source: https://github.com/ethereum/go-ethereum/blob/9231770811cda0473a7fa4e2bccc95bf62aae634/core/txpool/list.go#L284
func (etpc EthereumTxPoolConfig[C]) EthereumTxReplacementPolicy(op, np C, oldTx, newTx sdk.Tx) bool {
	// Convert the transactions to Ethereum transactions.
	oldEthTx := evmtypes.GetAsEthTx(oldTx)
	newEthTx := evmtypes.GetAsEthTx(newTx)
	if oldEthTx == nil || newEthTx == nil ||
		oldEthTx.GasFeeCapCmp(newEthTx) >= 0 || oldEthTx.GasTipCapCmp(newEthTx) >= 0 {
		return false
	}

	// thresholdFeeCap = oldFC  * (100 + priceBump) / 100
	a := big.NewInt(100 + int64(etpc.PriceBump)) //nolint:gomnd // 100% + priceBump.
	aFeeCap := new(big.Int).Mul(a, oldEthTx.GasFeeCap())
	aTip := a.Mul(a, oldEthTx.GasTipCap())

	// thresholdTip    = oldTip * (100 + priceBump) / 100
	b := big.NewInt(100) //nolint:gomnd // 100% + priceBump.
	thresholdFeeCap := aFeeCap.Div(aFeeCap, b)
	thresholdTip := aTip.Div(aTip, b)

	// We have to ensure that both the new fee cap and tip are higher than the
	// old ones as well as checking the percentage threshold to ensure that
	// this is accurate for low (Wei-level) gas price replacements.
	if newEthTx.GasFeeCapIntCmp(thresholdFeeCap) < 0 || newEthTx.GasTipCapIntCmp(thresholdTip) < 0 {
		return false
	}

	return true
}
