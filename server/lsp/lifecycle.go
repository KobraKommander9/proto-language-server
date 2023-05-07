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
func (s *LspServer) Initialize(ctx context.Context, params *p.InitializeParams) (*p.InitializeResult, error) {
	return s.engine.Initialize(ctx, params)
}

// Initialized -
func (s *LspServer) Initialized(ctx context.Context, _ *p.InitializedParams) error {
	return s.engine.Initialized(ctx)
}

// SetTrace -
func (s *LspServer) SetTrace(ctx context.Context, params *p.SetTraceParams) error {
	return s.engine.SetTrace(ctx, params)
}

// LogTrace -
func (*LspServer) LogTrace(ctx context.Context, params *p.LogTraceParams) error {
	return fmt.Errorf("unimplemented")
}

// Shutdown -
func (s *LspServer) Shutdown(ctx context.Context) error {
	return s.engine.Shutdown(ctx)
}

// Exit -
func (s *LspServer) Exit(context.Context) error {
	return s.engine.Close()
}
