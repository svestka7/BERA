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

package host

import (
	storetypes "cosmossdk.io/store/types"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	evmrpc "pkg.berachain.dev/polaris/cosmos/rpc"
	"pkg.berachain.dev/polaris/cosmos/store/offchain"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/block"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/configuration"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/gas"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/historical"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/vm"
	ethrpcconfig "pkg.berachain.dev/polaris/eth/rpc/config"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Compile-time interface assertion.
var _ core.PolarisHostChain = (*Host)(nil)

type Host struct {

	// We store a reference to the `rpcProvider` so that we can register it with
	// the cosmos mux router.
	rpcProvider evmrpc.Provider
	// The (unexposed) key used to access the store from the Context.
	storeKey storetypes.StoreKey
	// The offchain KV store.
	offChainKv *offchain.Store
	// authority is the bech32 address that is allowed to execute governance proposals.
	authority string

	// The various plugins that are are used to implement `core.PolarisHostChain`.
	bp  block.Plugin
	Cp  configuration.Plugin
	gp  gas.Plugin
	hp  historical.Plugin
	pp  precompile.Plugin
	sp  state.Plugin
	txp txpool.Plugin
}

// NewHost creates new instances of the plugin host.
func NewHost(
	storeKey storetypes.StoreKey,
	ak state.AccountKeeper,
	bk state.BankKeeper,
	authority string,
	appOpts servertypes.AppOptions,
	ethTxMempool sdkmempool.Mempool,
) *Host {
	// We setup the host with some Cosmos standard sauce.
	h := &Host{
		authority: authority,
		storeKey:  storeKey,
	}

	// TODO: parameterize kv store.
	if appOpts != nil {
		h.offChainKv = offchain.NewOffChainKVStore("eth_indexer", appOpts)
	}

	// Setup the RPC Service. // TODO: parameterize config.
	cfg := ethrpcconfig.DefaultServer()
	cfg.BaseRoute = "/eth/rpc"
	h.rpcProvider = evmrpc.NewProvider(cfg)

	// Build the Plugins
	h.bp = block.NewPlugin(storeKey)
	h.Cp = configuration.NewPlugin(storeKey)
	h.gp = gas.NewPlugin()
	h.hp = historical.NewPlugin(h.bp, h.offChainKv, storeKey)
	h.txp = txpool.NewPlugin(h.Cp, h.rpcProvider, utils.MustGetAs[*mempool.EthTxPool](ethTxMempool))
	return h
}

// Setup sets up the precompile and state plugins with the given precompiles and keepers. It also
// sets the query context function for the block and state plugins (to support historical queries).
func (h *Host) Setup(
	ak state.AccountKeeper,
	bk state.BankKeeper,
	precompiles []vm.RegistrablePrecompile,
	qc func(height int64, prove bool) (sdk.Context, error),
) {
	// Setup the precompile and state plugins
	h.pp = precompile.NewPlugin(precompiles)
	h.sp = state.NewPlugin(ak, bk, h.storeKey, h.Cp, h.pp)

	// Set the query context function for the block and state plugins
	h.sp.SetQueryContextFn(qc)
	h.bp.SetQueryContextFn(qc)
}

// GetBlockPlugin returns the header plugin.
func (h *Host) GetBlockPlugin() core.BlockPlugin {
	return h.bp
}

// GetConfigurationPlugin returns the configuration plugin.
func (h *Host) GetConfigurationPlugin() core.ConfigurationPlugin {
	return h.Cp
}

// GetGasPlugin returns the gas plugin.
func (h *Host) GetGasPlugin() core.GasPlugin {
	return h.gp
}

func (h *Host) GetHistoricalPlugin() core.HistoricalPlugin {
	return h.hp
}

// GetPrecompilePlugin returns the precompile plugin.
func (h *Host) GetPrecompilePlugin() core.PrecompilePlugin {
	return h.pp
}

// GetStatePlugin returns the state plugin.
func (h *Host) GetStatePlugin() core.StatePlugin {
	return h.sp
}

// GetTxPoolPlugin returns the txpool plugin.
func (h *Host) GetTxPoolPlugin() core.TxPoolPlugin {
	return h.txp
}

// GetAllPlugins returns all the plugins.
func (h *Host) GetAllPlugins() []plugins.BaseCosmosPolaris {
	return []plugins.BaseCosmosPolaris{h.hp, h.Cp, h.gp, h.pp, h.sp}
}
