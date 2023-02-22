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

// Package app defines and implements the domain of the server
package app

import (
	"sync"

	"go.lsp.dev/protocol"
	"go.uber.org/zap"
)

// Engine -
type Engine struct {
	EngineInfo
	l      *zap.SugaredLogger
	lock   *sync.Mutex
	config *EngineConfig
}

// NewEngine -
func NewEngine(l *zap.SugaredLogger, info EngineInfo) *Engine {
	return &Engine{
		EngineInfo: info,
		l:          l.Named("engine"),
		lock:       &sync.Mutex{},
	}
}

// Close -
func (e *Engine) Close() error {
	e.l.Infof("closing engine...")
	e.l.Infof("successfully closed engine")
	return nil
}

// Initialize - The initialize request is the first request form the client to the server. If
// the server receives a request or notification before the initialize request it should act as
// follows:
//   - For a request the response should be an error with code: -32002. The message can be
//     picked by the server
//   - Notifications should be dropped, except for the exit notification. This will allow
//     the exit of a server without an initialization request.
//
// Until the server has responded to the initialize request with an InitializeResult, the client
// must not send any additional requests or notifications to the server. In addition the server is
// not allowed to send any requests or notifications to the client until it has responded with an
// InitializeResult, with the exception that during the initialize request the server is allowed to
// send the notifications window/showMessage, window/logMessage and telemetry/event as well as the
// window/showMessageRequest request to the client. In case the client sets up a progress token in
// the initialize params the server is also allowed to use that token (and only that token) using
// the $/progress notification sent from the server to the client.
//
// The initialize request may only be sent once.
func (e *Engine) Initialize() protocol.InitializeResult {
	e.lock.Lock()
	defer e.lock.Unlock()

	return protocol.InitializeResult{
		Capabilities: e.Capabilities(),
		ServerInfo: &protocol.ServerInfo{
			Name:    e.Name,
			Version: e.Version,
		},
	}
}
