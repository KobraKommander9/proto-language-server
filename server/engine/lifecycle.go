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

// Package engine defines how the lsp server behaves
package engine

import (
	"context"
	"fmt"

	"go.lsp.dev/protocol"
)

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
func (e *Engine) Initialize(_ context.Context, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	e.lock.Lock()
	defer e.lock.Unlock()

	if params == nil {
		e.l.Warnf("attempted to initialize server with nil params")
		return nil, fmt.Errorf("cannot initialize with nil params")
	}

	e.config = &EngineConfig{
		clientInfo:       params.ClientInfo,
		capabilities:     params.Capabilities,
		workspaceFolders: params.WorkspaceFolders,
		workDoneToken:    params.WorkDoneToken,
	}

	if params.ProcessID != 0 {
		e.config.processId = &params.ProcessID
	}

	if params.RootPath != "" {
		e.config.rootPath = &params.RootPath
	}

	if params.RootURI != "" {
		e.config.rootUri = &params.RootURI
	}

	if params.Trace != "" {
		e.config.trace = &params.Trace
	}

	return &protocol.InitializeResult{
		Capabilities: e.Capabilities(),
		ServerInfo: &protocol.ServerInfo{
			Name:    e.Name,
			Version: e.Version,
		},
	}, nil
}

// Initialized - The initialized notification is sent from the client to the server after the client
// received the result of the initialize request but before the client is sending any other request
// or notification to the server. The server can use the initialized notification for example to
// dynamically register capabilities. The initialized notification may only be sent once.
func (e *Engine) Initialized(context.Context) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	if e.config == nil {
		e.l.Warnf("never initialized server before calling Initialized")
		return fmt.Errorf("server must be initialized")
	}

	if e.initialized.Load() {
		e.l.Warnf("already checked if the server was initialized")
		return fmt.Errorf("cannot call Initialized more than once")
	}

	e.initialized.Store(true)
	return nil
}

// SetTrace - A notification that should be used by the client to modify the trace setting of the
// server.
func (e *Engine) SetTrace(_ context.Context, params *protocol.SetTraceParams) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.config.trace = &params.Value

	return nil
}

// Shutdown - The shutdown request is sent from the client to the server. It asks the server to shut
// down, but to not exit (otherwise the response might not be delivered correctly to the client).
// There is a separate exit notification that asks the server to exit. Clients must not send any
// notifications other than exit or requests to a server to which they have sent a shutdown request.
// Clients should also wait with sending the exit notification until they have received a response
// from the shutdown request.
//
// If a server receives requests after a shutdown request those requests should error with InvalidRequest.
func (e *Engine) Shutdown(context.Context) error {
	if e.shutdown.Load() {
		e.l.Warnf("server already notified of shutdown")
		return fmt.Errorf("server already notified of shutdown")
	}

	e.shutdown.Store(true)

	return nil
}
