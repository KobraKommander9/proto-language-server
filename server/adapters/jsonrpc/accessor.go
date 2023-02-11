// Copyright 2023, Ed Kilgore
//
// This file is part of proto-language-server.
//
// proto-language-server is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// proto-language-server is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
// or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along with
// proto-language-server. If not, see <https://www.gnu.org/licenses/>.

// Package jsonrpc defines and implements types that interact with jsonrpc
package jsonrpc

import (
	"context"

	"go.lsp.dev/jsonrpc2"
)

// Accessor -
type Accessor interface {
	ListenAndServe(network, addr string, server jsonrpc2.StreamServer) error
}

// DefaultAccessor -
type DefaultAccessor struct{}

// ListenAndServe -
func (*DefaultAccessor) ListenAndServe(network, addr string, server jsonrpc2.StreamServer) error {
	return jsonrpc2.ListenAndServe(context.Background(), network, addr, server, 0)
}
