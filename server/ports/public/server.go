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

// Package public defines how a public facing entity should behave
package public

import "context"

// Server -
type Server interface {
	Serve(ctx context.Context) error
}
