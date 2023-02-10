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

import (
	nhttp "net/http"

	"github.com/KobraKommander9/proto-language-server/server/ports/http"

	log "github.com/sirupsen/logrus"
)

// PublicServer -
type PublicServer struct {
	Accessor
	addr    string
	service http.Service
}

// NewPublicServer -
func NewPublicServer(addr string, service http.Service, accessor Accessor) *PublicServer {
	return &PublicServer{
		Accessor: accessor,
		addr:     addr,
		service:  service,
	}
}

// Serve -
func (s *PublicServer) Serve() error {
	server, err := s.Accessor.Serve(s.addr, s.service)
	if err != nil {
		log.WithError(err).Warnf("could not serve %s http service", s.service.Name())
		return err
	}

	log.Infof("%s http service started", s.service.Name())
	if err = server.Serve(); err != nil && err != nhttp.ErrServerClosed {
		log.WithError(err).Warnf("could not close %s server", s.service.Name())
		return err
	}

	return nil
}
