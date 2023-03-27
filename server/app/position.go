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

// Package app defines and implements the domain of the server
package app

import (
	"github.com/KobraKommander9/proto-language-server/server/core"

	"go.lsp.dev/protocol"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// getPositionInFile - returns the line and column of the given position in
// the protobuf file
func (e *Engine) getPositionInFile(file protoreflect.FileDescriptor, pos protocol.Position) (core.Location, error) {
	fset := e.accessors.Parser.NewFileSet()
	node, err := e.accessors.Parser.ParseFile(fset, "", file, 0)
}
