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

	"github.com/KobraKommander9/proto-language-server/server/ports/lsp"

	"go.lsp.dev/jsonrpc2"
)

// LspServer -
type LspServer struct {
	engine lsp.Engine
}

// NewLspServer -
func NewLspServer(engine lsp.Engine) *LspServer {
	return &LspServer{
		engine: engine,
	}
}

// HandleRequest -
func (*LspServer) HandleRequest(_ context.Context, _ jsonrpc2.Replier, _ jsonrpc2.Request) error {
	return nil
}
