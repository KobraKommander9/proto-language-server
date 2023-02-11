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

// Package accessor defines and implements accessors for testing
package accessor

import (
	"context"
	"fmt"
	"io"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
)

// JsonRpcAccessor -
type JsonRpcAccessor interface {
	ListenAndServe(ctx context.Context, port uint32, server protocol.Server) error
	NewConn(jsonrpc2.Stream) jsonrpc2.Conn
	NewStream(io.ReadWriteCloser) jsonrpc2.Stream
}

// DefaultJsonRpcAccessor -
type DefaultJsonRpcAccessor struct{}

// ListenAndServe -
func (*DefaultJsonRpcAccessor) ListenAndServe(ctx context.Context, port uint32, server protocol.Server) error {
	s := jsonrpc2.HandlerServer(protocol.ServerHandler(server, nil))
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	return jsonrpc2.ListenAndServe(ctx, "tcp", addr, s, 0)
}

// NewConn -
func (*DefaultJsonRpcAccessor) NewConn(stream jsonrpc2.Stream) jsonrpc2.Conn {
	return jsonrpc2.NewConn(stream)
}

// NewStream -
func (*DefaultJsonRpcAccessor) NewStream(conn io.ReadWriteCloser) jsonrpc2.Stream {
	return jsonrpc2.NewStream(conn)
}
