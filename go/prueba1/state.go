// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package prueba1

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type Immutableprueba1State struct {
	id int32
}

func (s Immutableprueba1State) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(StateOwner))
}

func (s Immutableprueba1State) Text() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, wasmlib.KeyID(StateText))
}

type Mutableprueba1State struct {
	id int32
}

func (s Mutableprueba1State) AsImmutable() Immutableprueba1State {
	return Immutableprueba1State(s)
}

func (s Mutableprueba1State) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(StateOwner))
}

func (s Mutableprueba1State) Text() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, wasmlib.KeyID(StateText))
}
