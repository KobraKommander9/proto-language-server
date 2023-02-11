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

// Package jsonrpc defines and implements types that interact with jsonrpc
package jsonrpc

import (
	"github.com/KobraKommander9/proto-language-server/server/ports/jsonrpc"

	log "github.com/sirupsen/logrus"
)

// PublicServer -
type PublicServer struct {
	Accessor
	network string
	addr    string
	server  jsonrpc.Server
}

// NewPublicServer -
func NewPublicServer(network, addr string, server jsonrpc.Server, accessor Accessor) *PublicServer {
	return &PublicServer{
		Accessor: accessor,
		network:  network,
		addr:     addr,
		server:   server,
	}
}

// Serve -
func (s *PublicServer) Serve() error {
	log.Infof("serving jsonrpc server with %s on %s", s.network, s.addr)
	return s.ListenAndServe(s.network, s.addr, s.HandlerServer(s.server.HandleRequest))
}
