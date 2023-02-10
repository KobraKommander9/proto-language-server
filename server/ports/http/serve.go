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

// Package http defines and implements types for communicating over http
package http

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

const (
	maxHeaderByte = 1 * 1000000
	maxBodyByte   = 10 * 1000000
)

// ConnServer - an wrapper around http.Server for testing
type ConnServer interface {
	Close() error
	ListenAndServe() error
}

// Server -
type Server struct {
	serveMux *http.ServeMux
	server   ConnServer
}

// Serve -
func Serve(addr string, service Service) (*Server, error) {
	mux := http.NewServeMux()
	httpHandler := http.MaxBytesHandler(mux, maxBodyByte)

	s := &Server{
		serveMux: mux,
		server: &http.Server{
			Addr:           addr,
			Handler:        httpHandler,
			MaxHeaderBytes: maxHeaderByte,
		},
	}

	err := service.Register(mux)

	if err != nil {
		log.WithError(err).Warnf("failed to register %s http service", service.Name())
		return nil, err
	}

	log.WithError(err).Warnf("successfully registered %s http service with address %s", service.Name(), addr)

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		_ = s.server.Close()
	}()

	return s, nil
}

// Serve -
func (s *Server) Serve() error {
	return s.server.ListenAndServe()
}
