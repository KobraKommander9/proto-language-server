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

	"go.lsp.dev/protocol"
	"go.uber.org/zap"
)

// PublicServer -
type PublicServer struct {
	l       *zap.SugaredLogger
	options ServerOptions
	server  protocol.Server
}

// NewPublicServer -
func NewPublicServer(l *zap.SugaredLogger, server protocol.Server, opts ...Option) *PublicServer {
	options := newServerOptions()
	for _, opt := range opts {
		opt.apply(&options)
	}

	return &PublicServer{
		l:       l.Named("public"),
		options: options,
	}
}

// Serve -
func (s *PublicServer) Serve(ctx context.Context) error {
	if s.options.method == nil {
		s.l.DPanicf("no method found in options when serving jsonrpc content")
	}

	return s.options.method.serve(ctx, s.server)
}
