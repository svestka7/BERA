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

package store

import (
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/lib/utils"
)

// DenomERC20 is the store type for ERC20 <-> x/bank module denominations.
type DenomKVStore interface {
	SetDenomForAddress(address common.Address)
	GetDenomForAddress(address common.Address) string
	HasDenomForAddress(address common.Address) bool
	SetAddressForDenom(address common.Address)
	GetAddressForDenom(denom string) common.Address
	HasAddressForDenom(denom string) bool
}

// denomStore is a store that stores information regarding ERC20 <-> x/bank module denominations.
type denomStore struct {
	addressToDenom storetypes.KVStore
	denomToAddress storetypes.KVStore
}

// NewDenomKVStore creates a new DenomKVStore.
func NewDenomKVStore(store storetypes.KVStore) DenomKVStore {
	return &denomStore{
		addressToDenom: prefix.NewStore(store, []byte{types.AddressToDenomKeyPrefix}),
		denomToAddress: prefix.NewStore(store, []byte{types.DenomToAddressKeyPrefix}),
	}
}

// ==============================================================================
// ERC20 -> Denom
// ==============================================================================

// SetDenomForAddress sets the address correlated to a specific denomination.
func (ds denomStore) SetDenomForAddress(address common.Address) {
	ds.addressToDenom.Set(address.Bytes(), utils.UnsafeStrToBytes(types.DenomForAddress(address)))
}

// GetDenomForAddress returns the denomination correlated to a specific address.
func (ds denomStore) GetDenomForAddress(address common.Address) string {
	bz := ds.addressToDenom.Get(address.Bytes())
	if bz == nil {
		return ""
	}
	return utils.UnsafeBytesToStr(bz)
}

// HasDenomForAddress returns true if the address has a denomination.
func (ds denomStore) HasDenomForAddress(address common.Address) bool {
	return ds.addressToDenom.Has(address.Bytes())
}

// ==============================================================================
// Denom -> ERC20
// ==============================================================================

// SetAddressForDenom sets the denomination correlated to a specific address.
func (ds denomStore) SetAddressForDenom(address common.Address) {
	ds.denomToAddress.Set(utils.UnsafeStrToBytes(types.DenomForAddress(address)), address.Bytes())
}

// GetAddressForDenom returns the address correlated to a specific denomination.
func (ds denomStore) GetAddressForDenom(denom string) common.Address {
	bz := ds.denomToAddress.Get(utils.UnsafeStrToBytes(denom))
	if bz == nil {
		return common.Address{}
	}
	return common.BytesToAddress(bz)
}

// HasAddressForDenom returns true if the denomination has an address.
func (ds denomStore) HasAddressForDenom(denom string) bool {
	return ds.denomToAddress.Has(utils.UnsafeStrToBytes(denom))
}
