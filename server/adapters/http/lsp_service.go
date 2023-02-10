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
	"net/http"

	"github.com/KobraKommander9/proto-language-server/server/ports/lsp"

	m "github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	log "github.com/sirupsen/logrus"
)

// Header key constants
const (
	ContentTypeKey = "content-type"
)

var (
	allowedContentTypes = []string{
		"utf8", // for backwards compatibility
		"utf-8",
	}
)

// LspService -
type LspService struct {
	service lsp.Service
}

// NewLspService -
func NewLspService(service lsp.Service) *LspService {
	return &LspService{
		service: service,
	}
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
func (s *LspService) Register(mux *http.ServeMux) error {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	r := m.NewRouter()
	r.Use(s.loggingMiddleware)
	r.Use(s.contectCheckerMiddleware)

	r.HandleFunc("/", server)

	mux.Handle("/", r)

	return nil
}

func (*LspService) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"header":        r.Header,
			"url":           r.URL,
			"form":          r.Form,
			"contentLength": r.ContentLength,
			"host":          r.Host,
		}).Debugf("received http request to %v", r.URL)
		next.ServeHTTP(w, r)
	})
}

func (*LspService) contectCheckerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get(ContentTypeKey)

		found := false
		for _, ct := range allowedContentTypes {
			if contentType == "" || contentType == ct {
				found = true
				break
			}
		}

		if !found {
			log.WithField("contentType", contentType).Warnf("received request with unallowed content type %s", contentType)
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
	})
}
