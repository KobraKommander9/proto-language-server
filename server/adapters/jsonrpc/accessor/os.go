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

// Package accessor defines and implements accessors for testing
package accessor

import "os"

// OSAccessor -
type OSAccessor interface {
	Read(b []byte) (int, error)
	Write(b []byte) (int, error)
	Close() error
}

// DefaultOSAccessor -
type DefaultOSAccessor struct{}

// Read -
func (*DefaultOSAccessor) Read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}

// Write -
func (*DefaultOSAccessor) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}

// Close -
func (*DefaultOSAccessor) Close() error {
	return nil
}
