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

type stdio struct {
	accessor.JsonRpcAccessor
	accessor.OSAccessor
	l *zap.SugaredLogger
}

func newStdio(a1 accessor.JsonRpcAccessor, a2 accessor.OSAccessor) *stdio {
	return &stdio{
		JsonRpcAccessor: a1,
		OSAccessor:      a2,
		l:               zap.S().Named("stdio"),
	}
}

func (s *stdio) apply(opts *ServerOptions) {
	opts.method = s
}

func (s *stdio) serve(ctx context.Context, server protocol.Server) error {
	s.l.Infof("serving jsonrpc server over stdio")

	conn := s.NewConn(s.NewStream(s))
	conn.Go(ctx, protocol.ServerHandler(server, nil))
	<-conn.Done()

	s.l.Infof("finished serving jsonrpc server over stdio")
	return nil
}

// Read -
func (s *stdio) Read(b []byte) (int, error) {
	return s.OSAccessor.Read(b)
}

// Write -
func (s *stdio) Write(b []byte) (int, error) {
	return s.OSAccessor.Write(b)
}

// Close -
func (*stdio) Close() error {
	return nil
}
