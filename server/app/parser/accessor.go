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

// Package parser defines and implements how to parse a proto file
package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Accessor -
type Accessor interface {
	NewFileSet() *token.FileSet
	ParseFile(fset *token.FileSet, filename string, src any, mode parser.Mode) (*ast.File, error)
}

// DefaultAccessor -
type DefaultAccessor struct{}

// NewFileSet -
func (*DefaultAccessor) NewFileSet() *token.FileSet {
	return token.NewFileSet()
}

// ParseFile -
func (*DefaultAccessor) ParseFile(fset *token.FileSet, filename string, src any, mode parser.Mode) (*ast.File, error) {
	return parser.ParseFile(fset, filename, src, mode)
}
