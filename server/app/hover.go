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
	"context"
	"fmt"
	"regexp"
	"strings"

	"go.lsp.dev/protocol"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Hover - The hover request is sent from the client to the server to
// request hover information at a given text document position.
func (e *Engine) Hover(_ context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	uri := params.TextDocument.URI
	pos := params.Position
	registry := protoregistry.GlobalFiles
	_, err := registry.FindFileByPath(string(uri))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// getFullName - returns the full name to the message or field at the given position.
func (e *Engine) getFullName(file protoreflect.FileDescriptor, pos protocol.Position) (string, error) {
	line := pos.Line + 1
	col := int(pos.Character)

	data, err := e.clients.Filesystem.ReadFile(string(file.Path()))
	if err != nil {
		return "", err
	}

	lineStart := 0
	for i := 1; i < int(line); i++ {
		idx := strings.IndexByte(string(data[lineStart:]), '\n') + 1
		if idx == 0 {
			return "", fmt.Errorf("line %d not found", line)
		}
		lineStart += idx
	}

	lineEnd := lineStart + strings.IndexByte(string(data[lineStart:]), '\n')
	if lineEnd == -1 {
		lineEnd = len(data)
	}

	lineData := string(data[lineStart:lineEnd])

	identifiers := regexp.MustCompile(`[A-Za-z_][A-Za-z_0-9]*`)
	posInLine := 0
	for _, identifier := range identifiers.FindAllString(lineData, -1) {
		if col >= posInLine && col < posInLine+len(identifier) {
			return identifier, nil
		}
		posInLine += len(identifier)
	}

	return "", fmt.Errorf("identifier not found")
}
