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

// DidOpen -
func (*LspServer) DidOpen(ctx context.Context, params *p.DidOpenTextDocumentParams) error {
	return fmt.Errorf("unimplemented")
}

// DidChange -
func (*LspServer) DidChange(ctx context.Context, params *p.DidChangeTextDocumentParams) error {
	return fmt.Errorf("unimplemented")
}

// WillSave -
func (*LspServer) WillSave(ctx context.Context, params *p.WillSaveTextDocumentParams) error {
	return fmt.Errorf("unimplemented")
}

// WillSaveWaitUntil -
func (*LspServer) WillSaveWaitUntil(ctx context.Context, params *p.WillSaveTextDocumentParams) ([]p.TextEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidSave -
func (*LspServer) DidSave(ctx context.Context, params *p.DidSaveTextDocumentParams) error {
	return fmt.Errorf("unimplemented")
}

// DidClose -
func (*LspServer) DidClose(ctx context.Context, params *p.DidCloseTextDocumentParams) error {
	return fmt.Errorf("unimplemented")
}
