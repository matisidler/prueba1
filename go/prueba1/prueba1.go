// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package prueba1

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	if f.Params.Owner().Exists() {
		f.State.Owner().SetValue(f.Params.Owner().Value())
		return
	}
	f.State.Owner().SetValue(ctx.ContractCreator())
}

func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
	f.State.Owner().SetValue(f.Params.Owner().Value())
}

func funcSetText(ctx wasmlib.ScFuncContext, f *SetTextContext) {
	f.State.Text().SetValue(f.Params.Res().String())
}

func viewGetOwner(ctx wasmlib.ScViewContext, f *GetOwnerContext) {
	f.Results.Owner().SetValue(f.State.Owner().Value())
}

func viewGetText(ctx wasmlib.ScViewContext, f *GetTextContext) {
	if f.State.Text().Exists() {
		f.Results.Res().SetValue(f.State.Text().String())
	} else {
		f.Results.Res().SetValue("value doesn't exists")
	}
}
