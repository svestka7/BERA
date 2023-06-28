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

package localnet

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLocalnet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "testing:integration")
}

var _ = Describe("Fixture", func() {
	var (
		c   Localnet
		err error
	)

	BeforeEach(func() {
		_, err = NewDockerizedNetwork(
			"base",
			baseImageName,
			baseContext,
			baseDockerfile,
			"8545/tcp",
			"8546/tcp",
			map[string]string{
				"GO_VERSION":               "1.20.4",
				"PRECOMPILE_CONTRACTS_DIR": "./contracts",
				"GOOS":                     "linux",
				"GOARCH":                   "arm64",
			},
		)
		Expect(err).ToNot(HaveOccurred())

		c, err = NewDockerizedNetwork(
			"localnet1",
			localnetImageName,
			localnetContext,
			localnetDockerfile,
			"8545/tcp",
			"8546/tcp",
			map[string]string{
				"GO_VERSION":   "1.20.4",
				"BASE_IMAGE":   baseImageName,
				"GENESIS_PATH": "config",
			},
		)
		Expect(err).ToNot(HaveOccurred())
		Expect(c).ToNot(BeNil())
	})

	AfterEach(func() {
		Expect(c.Stop()).To(Succeed())

	})

	It("should create a container", func() {
		err = c.Start()
		Expect(err).ToNot(HaveOccurred())
	})
})
