// Copyright 2023, Ed Kilgore
//
// This file is part of protobuf-language-server.
//
// protobuf-language-server is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// protobuf-language-server is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
// or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along with
// protobuf-language-server. If not, see <https://www.gnu.org/licenses/>.

// Package jsonrpc defines how to interact with jsonrpc
package jsonrpc

import (
	"context"

	"go.lsp.dev/jsonrpc2"
)

// Server -
type Server interface {
	HandleRequest(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error
}
