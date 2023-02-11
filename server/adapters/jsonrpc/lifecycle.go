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
	"fmt"

	p "go.lsp.dev/protocol"
)

// Initialize -
func (*LspServer) Initialize(context.Context, *p.InitializeParams) (*p.InitializeResult, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Initialized -
func (*LspServer) Initialized(context.Context, *p.InitializedParams) error {
	return fmt.Errorf("unimplemented")
}

// SetTrace -
func (*LspServer) SetTrace(ctx context.Context, params *p.SetTraceParams) error {
	return fmt.Errorf("unimplemented")
}

// LogTrace -
func (*LspServer) LogTrace(ctx context.Context, params *p.LogTraceParams) error {
	return fmt.Errorf("unimplemented")
}

// Shutdown -
func (*LspServer) Shutdown(context.Context) error {
	return fmt.Errorf("unimplemented")
}

// Exit -
func (*LspServer) Exit(context.Context) error {
	return fmt.Errorf("unimplemented")
}
