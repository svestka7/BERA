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

package staking

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/staking"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Contract is the precompile contract for the staking module.
type Contract struct {
	ethprecompile.BaseContract

	msgServer stakingtypes.MsgServer
	querier   stakingtypes.QueryServer
}

// NewContract is the constructor of the staking contract.
func NewPrecompileContract(sk *stakingkeeper.Keeper) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.StakingModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(stakingtypes.ModuleName)),
		),
		msgServer: stakingkeeper.NewMsgServerImpl(sk),
		querier:   stakingkeeper.Querier{Keeper: sk},
	}
}

// GetDelegation implements `getDelegation(address)` method.
func (c *Contract) GetDelegation(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	del, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	val, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return c.getDelegationHelper(
		ctx, cosmlib.AddressToAccAddress(del), cosmlib.AddressToValAddress(val),
	)
}

// GetDelegation0 implements `getDelegation(string)` method.
func (c *Contract) GetDelegation0(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	bech32DelAddr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	del, err := sdk.AccAddressFromBech32(bech32DelAddr)
	if err != nil {
		return nil, err
	}
	bech32ValAddr, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	val, err := sdk.ValAddressFromBech32(bech32ValAddr)
	if err != nil {
		return nil, err
	}

	return c.getDelegationHelper(ctx, del, val)
}

// GetUnbondingDelegation implements the `getUnbondingDelegation(address)` method.
func (c *Contract) GetUnbondingDelegation(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	del, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	val, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return c.getUnbondingDelegationHelper(
		ctx, cosmlib.AddressToAccAddress(del), cosmlib.AddressToValAddress(val),
	)
}

// GetUnbondingDelegation0 implements the `getUnbondingDelegation(string)` method.
func (c *Contract) GetUnbondingDelegation0(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	bech32DelAddr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	del, err := sdk.AccAddressFromBech32(bech32DelAddr)
	if err != nil {
		return nil, err
	}
	bech32ValAddr, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	val, err := sdk.ValAddressFromBech32(bech32ValAddr)
	if err != nil {
		return nil, err
	}

	return c.getUnbondingDelegationHelper(ctx, del, val)
}

// GetRedelegations implements the `getRedelegations(address,address)` method.
func (c *Contract) GetRedelegations(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	del, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	srcVal, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	dstVal, ok := utils.GetAs[common.Address](args[2])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return c.getRedelegationsHelper(
		ctx,
		cosmlib.AddressToAccAddress(del),
		cosmlib.AddressToValAddress(srcVal),
		cosmlib.AddressToValAddress(dstVal),
	)
}

// GetRedelegations0 implements the `getRedelegations(string,string)` method.
func (c *Contract) GetRedelegations0(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	bech32DelAddr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	srcVal, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	dstVal, ok := utils.GetAs[string](args[2])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	del, err := sdk.AccAddressFromBech32(bech32DelAddr)
	if err != nil {
		return nil, err
	}
	src, err := sdk.ValAddressFromBech32(srcVal)
	if err != nil {
		return nil, err
	}
	dst, err := sdk.ValAddressFromBech32(dstVal)
	if err != nil {
		return nil, err
	}

	return c.getRedelegationsHelper(ctx, del, src, dst)
}

// Delegate implements the `delegate(address,uint256)` method.
func (c *Contract) Delegate(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	return c.delegateHelper(ctx, caller, amount, cosmlib.AddressToValAddress(val))
}

// Delegate0 implements the `delegate(string,uint256)` method.
func (c *Contract) Delegate0(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return c.delegateHelper(ctx, caller, amount, val)
}

// Undelegate implements the `undelegate(address,uint256)` method.
func (c *Contract) Undelegate(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	return c.undelegateHelper(ctx, caller, amount, cosmlib.AddressToValAddress(val))
}

// Undelegate0 implements the `undelegate(string,uint256)` method.
func (c *Contract) Undelegate0(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return c.undelegateHelper(ctx, caller, amount, val)
}

// BeginRedelegate implements the `beginRedelegate(address,address,uint256)` method.
func (c *Contract) BeginRedelegate(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	srcVal, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	dstVal, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[2])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	return c.beginRedelegateHelper(
		ctx,
		caller,
		amount,
		cosmlib.AddressToValAddress(srcVal),
		cosmlib.AddressToValAddress(dstVal),
	)
}

// BeginRedelegate0 implements the `beginRedelegate(string,string,uint256)` method.
func (c *Contract) BeginRedelegate0(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	srcVal, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	dstVal, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[2])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	src, err := sdk.ValAddressFromBech32(srcVal)
	if err != nil {
		return nil, err
	}
	dst, err := sdk.ValAddressFromBech32(dstVal)
	if err != nil {
		return nil, err
	}

	return c.beginRedelegateHelper(ctx, caller, amount, src, dst)
}

// CancelRedelegate implements the `cancelRedelegate(address,address,uint256,int64)` method.
func (c *Contract) CancelUnbondingDelegation(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}
	creationHeight, ok := utils.GetAs[int64](args[2])
	if !ok {
		return nil, precompile.ErrInvalidInt64
	}

	return c.cancelUnbondingDelegationHelper(ctx, caller, amount, cosmlib.AddressToValAddress(val), creationHeight)
}

// CancelRedelegate0 implements the `cancelRedelegate(string,string,uint256,int64)` method.
func (c *Contract) CancelUnbondingDelegation0(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}
	creationHeight, ok := utils.GetAs[int64](args[2])
	if !ok {
		return nil, precompile.ErrInvalidInt64
	}

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return c.cancelUnbondingDelegationHelper(ctx, caller, amount, val, creationHeight)
}

// GetActiveValidators implements the `getActiveValidators()` method.
func (c *Contract) GetActiveValidators(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	_ ...any,
) ([]any, error) {
	return c.activeValidatorsHelper(ctx)
}

// GetValidators implements the `getValidators()` method.
func (c *Contract) GetValidators(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	_ ...any,
) ([]any, error) {
	return c.validatorsHelper(ctx)
}

// GetValidators implements the `getValidator(address)` method.
func (c *Contract) GetValidator(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return c.validatorHelper(ctx, sdk.ValAddress(val[:]).String())
}

// GetValidators implements the `getValidator(string)` method.
func (c *Contract) GetValidator0(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	valBech32, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	return c.validatorHelper(ctx, valBech32)
}

// GetDelegatorValidators implements the `getDelegatorValidators(address)` method.
func (c *Contract) GetDelegatorValidators(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	del, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return c.delegatorValidatorsHelper(ctx, cosmlib.Bech32FromEthAddress(del))
}

// GetDelegatorValidators0 implements the `getDelegatorValidators(string)` method.
func (c *Contract) GetDelegatorValidators0(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	delBech32, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	return c.delegatorValidatorsHelper(ctx, delBech32)
}
