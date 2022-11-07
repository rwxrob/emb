// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

package emb

import (
	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/fn/filt"
)

// comp fulfills the bonzai.Completer interface.
type comp struct{}

func (comp) Complete(x bonzai.Command, args ...string) []string {

	if len(args) > 1 {
		return []string{}
	}

	if args == nil || (len(args) > 0 && args[0] == "") {
		return RelPaths()
	}

	// catch edge cases
	if len(args) == 0 {
		if x != nil {
			return []string{x.GetName()} // will add tailing space
		}
		return RelPaths()
	}

	paths := RelPaths()
	return filt.HasPrefix(paths, args[0])

	return []string{}
}
