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
	"context"

	"github.com/KobraKommander9/proto-language-server/server/adapters/jsonrpc/accessor"

	"go.lsp.dev/protocol"
	"go.uber.org/zap"
)

type socket struct {
	accessor.JsonRpcAccessor
	l    *zap.SugaredLogger
	port uint32
}

func newSocket(a accessor.JsonRpcAccessor, port uint32) *socket {
	return &socket{
		JsonRpcAccessor: a,
		l:               zap.S().Named("socket"),
		port:            port,
	}
}

func (s *socket) apply(opts *ServerOptions) {
	opts.method = s
}

func (s *socket) serve(ctx context.Context, server protocol.Server) error {
	s.l.Infof("serving jsonrpc server over socket with port %d", s.port)
	return s.ListenAndServe(ctx, s.port, server)
}
