// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/matisidler/prueba1/go/prueba1"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, prueba1.ScName, prueba1.OnLoad)
	require.NoError(t, ctx.ContractExists(prueba1.ScName))
}
