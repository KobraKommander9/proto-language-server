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

// Package cmd defines and implements commands for the server
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/KobraKommander9/proto-language-server/server/adapters/jsonrpc"
	"github.com/KobraKommander9/proto-language-server/server/adapters/jsonrpc/accessor"
	jp "github.com/KobraKommander9/proto-language-server/server/adapters/jsonrpc/public"
	"github.com/KobraKommander9/proto-language-server/server/app"
	"github.com/KobraKommander9/proto-language-server/server/ports/lsp"
	"github.com/KobraKommander9/proto-language-server/server/ports/public"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	lspType      = "lsp-type"
	stdioMethod  = "stdio"
	socketMethod = "port"
)

const (
	jsonrpcType = "jsonrpc://"
)

// StartCmd -
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long: `starts the omni api service.
	Example: ./bin/apictl start`,
	RunE: func(_ *cobra.Command, _ []string) error {
		ctx := context.Background()

		engine := app.NewEngine(zap.S())

		server, err := createLspServer(viper.GetString(lspType), engine)
		if err != nil {
			return fmt.Errorf("failed to create lsp server: %v", err)
		}

		return server.Serve(ctx)
	},
}

func init() {
	RootCmd.AddCommand(StartCmd)

	StartCmd.PersistentFlags().String(lspType, jsonrpcType, "type of lsp server")
	_ = viper.BindPFlag(lspType, StartCmd.PersistentFlags().Lookup(lspType))

	StartCmd.PersistentFlags().Bool(stdioMethod, false, "if the lsp server should handle requests over stdio")
	_ = viper.BindPFlag(stdioMethod, StartCmd.PersistentFlags().Lookup(stdioMethod))

	StartCmd.PersistentFlags().Uint32(socketMethod, 0, "the port when using sockets as the communication channel")
	_ = viper.BindPFlag(socketMethod, StartCmd.PersistentFlags().Lookup(socketMethod))
}

func createLspServer(t string, engine lsp.Engine) (server public.Server, err error) {
	switch {
	case strings.HasPrefix(t, jsonrpcType):
		lspServer := jsonrpc.NewLspServer(zap.S(), engine)
		server, err = createPublicServer(lspServer)

	default:
		return nil, fmt.Errorf("unknown lsp server type %s", t)
	}

	return server, err
}

func createPublicServer(lspServer *jsonrpc.LspServer) (server public.Server, err error) {
	port := viper.GetUint32(socketMethod)

	switch {
	case port != 0:
		zap.S().Infof("creating jsonrpc lsp server on port %d", port)
		server = jp.NewPublicServer(zap.S(), lspServer, jp.WithSocket(&accessor.DefaultJsonRpcAccessor{}, port))

	case viper.GetBool(stdioMethod):
		zap.S().Infof("creating jsonrpc lsp server with stdio communication")
		server = jp.NewPublicServer(zap.S(), lspServer, jp.WithStdio(&accessor.DefaultJsonRpcAccessor{}, &accessor.DefaultOSAccessor{}))

	default:
		return nil, fmt.Errorf("no method type provided for lsp server")
	}

	return server, err
}
