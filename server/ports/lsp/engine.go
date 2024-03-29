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

// Package lsp defines how to interact with the lsp
package lsp

import (
	"context"

	"go.lsp.dev/protocol"
)

// Engine -
type Engine interface {
	Close() error

	Initialize(ctx context.Context, params *protocol.InitializeParams) (*protocol.InitializeResult, error)
	Initialized(ctx context.Context) error
	SetTrace(ctx context.Context, params *protocol.SetTraceParams) error
	Shutdown(ctx context.Context) error
}
