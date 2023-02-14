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

// Package public defines and implements public facing types for jsonrpc
package public

import (
	"github.com/KobraKommander9/proto-language-server/server/adapters/jsonrpc/accessor"
)

// ServerOptions -
type ServerOptions struct {
	method method
}

// Option -
type Option interface {
	apply(s *ServerOptions)
}

func newServerOptions() ServerOptions {
	return ServerOptions{
		method: newStdio(&accessor.DefaultJsonRpcAccessor{}, &accessor.DefaultOSAccessor{}),
	}
}

type funcOption struct {
	f func(*ServerOptions)
}

func (o *funcOption) apply(opts *ServerOptions) {
	o.f(opts)
}

func newFuncOption(f func(*ServerOptions)) Option {
	return &funcOption{
		f: f,
	}
}

// WithStdio -
func WithStdio(a1 accessor.JsonRpcAccessor, a2 accessor.OSAccessor) Option {
	return newFuncOption(func(opts *ServerOptions) {
		opts.method = newStdio(a1, a2)
	})
}

// WithSocket -
func WithSocket(a accessor.JsonRpcAccessor, port uint32) Option {
	return newFuncOption(func(opts *ServerOptions) {
		opts.method = newSocket(a, port)
	})
}
