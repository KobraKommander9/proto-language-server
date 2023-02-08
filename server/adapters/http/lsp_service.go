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

// Package http defines and implements types that interact with http
package http

import "net/http"

// LspService -
type LspService struct{}

// NewLspService -
func NewLspService() *LspService {
	return &LspService{}
}

// Close -
func (*LspService) Close() error {
	return nil
}

// Name -
func (*LspService) Name() string {
	return "proto-language-server"
}

// Register -
func (*LspService) Register(_ *http.ServeMux) error {
	return nil
}
