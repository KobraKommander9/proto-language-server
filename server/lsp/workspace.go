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

// Package lsp defines the language server protocol
package lsp

import (
	"context"
	"fmt"

	p "go.lsp.dev/protocol"
)

// Symbols -
func (*Server) Symbols(ctx context.Context, params *p.WorkspaceSymbolParams) ([]p.SymbolInformation, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidChangeConfiguration -
func (*Server) DidChangeConfiguration(ctx context.Context, params *p.DidChangeConfigurationParams) error {
	return fmt.Errorf("unimplemented")
}

// DidChangeWorkspaceFolders -
func (*Server) DidChangeWorkspaceFolders(ctx context.Context, params *p.DidChangeWorkspaceFoldersParams) error {
	return fmt.Errorf("unimplemented")
}

// WillCreateFiles -
func (*Server) WillCreateFiles(ctx context.Context, params *p.CreateFilesParams) (*p.WorkspaceEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidCreateFiles -
func (*Server) DidCreateFiles(ctx context.Context, params *p.CreateFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// WillRenameFiles -
func (*Server) WillRenameFiles(ctx context.Context, params *p.RenameFilesParams) (*p.WorkspaceEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidRenameFiles -
func (*Server) DidRenameFiles(ctx context.Context, params *p.RenameFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// WillDeleteFiles -
func (*Server) WillDeleteFiles(ctx context.Context, params *p.DeleteFilesParams) (*p.WorkspaceEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidDeleteFiles -
func (*Server) DidDeleteFiles(ctx context.Context, params *p.DeleteFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// DidChangeWatchedFiles -
func (*Server) DidChangeWatchedFiles(ctx context.Context, params *p.DidChangeWatchedFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// ExecuteCommand -
func (*Server) ExecuteCommand(ctx context.Context, params *p.ExecuteCommandParams) (any, error) {
	return nil, fmt.Errorf("unimplemented")
}
