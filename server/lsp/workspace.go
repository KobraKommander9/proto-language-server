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

// Symbols -
func (*LspServer) Symbols(ctx context.Context, params *p.WorkspaceSymbolParams) ([]p.SymbolInformation, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidChangeConfiguration -
func (*LspServer) DidChangeConfiguration(ctx context.Context, params *p.DidChangeConfigurationParams) error {
	return fmt.Errorf("unimplemented")
}

// DidChangeWorkspaceFolders -
func (*LspServer) DidChangeWorkspaceFolders(ctx context.Context, params *p.DidChangeWorkspaceFoldersParams) error {
	return fmt.Errorf("unimplemented")
}

// WillCreateFiles -
func (*LspServer) WillCreateFiles(ctx context.Context, params *p.CreateFilesParams) (*p.WorkspaceEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidCreateFiles -
func (*LspServer) DidCreateFiles(ctx context.Context, params *p.CreateFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// WillRenameFiles -
func (*LspServer) WillRenameFiles(ctx context.Context, params *p.RenameFilesParams) (*p.WorkspaceEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidRenameFiles -
func (*LspServer) DidRenameFiles(ctx context.Context, params *p.RenameFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// WillDeleteFiles -
func (*LspServer) WillDeleteFiles(ctx context.Context, params *p.DeleteFilesParams) (*p.WorkspaceEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidDeleteFiles -
func (*LspServer) DidDeleteFiles(ctx context.Context, params *p.DeleteFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// DidChangeWatchedFiles -
func (*LspServer) DidChangeWatchedFiles(ctx context.Context, params *p.DidChangeWatchedFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// ExecuteCommand -
func (*LspServer) ExecuteCommand(ctx context.Context, params *p.ExecuteCommandParams) (any, error) {
	return nil, fmt.Errorf("unimplemented")
}
