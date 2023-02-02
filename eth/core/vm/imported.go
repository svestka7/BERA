// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package vm

import (
	"github.com/ethereum/go-ethereum/core/vm"
)

type (
	AccountRef          = vm.AccountRef
	BlockContext        = vm.BlockContext
	CanTransferFunc     = vm.CanTransferFunc
	ContractRef         = vm.ContractRef
	Config              = vm.Config
	EVMLogger           = vm.EVMLogger
	GethEVM             = vm.EVM
	GethStateDB         = vm.StateDB
	GetHashFunc         = vm.GetHashFunc
	PrecompileContainer = vm.PrecompiledContract
	PrecompileManager   = vm.PrecompileManager
	TransferFunc        = vm.TransferFunc
	TxContext           = vm.TxContext
)

var (
	NewGethEVM                = vm.NewEVM
	NewGethEVMWithPrecompiles = vm.NewEVMWithPrecompiles
	ErrOutOfGas               = vm.ErrOutOfGas
)
