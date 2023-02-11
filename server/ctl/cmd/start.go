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
	"fmt"
	"strings"

	"github.com/KobraKommander9/proto-language-server/server/adapters/jsonrpc"
	"github.com/KobraKommander9/proto-language-server/server/app"
	"github.com/KobraKommander9/proto-language-server/server/ports/lsp"
	"github.com/KobraKommander9/proto-language-server/server/ports/public"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	lspType = "lsp-type"
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
		engine := app.NewEngine()

		server, err := createLspServer(viper.GetString(lspType), engine)
		if err != nil {
			return fmt.Errorf("failed to create lsp server: %v", err)
		}

		return server.Serve()
	},
}

func init() {
	RootCmd.AddCommand(StartCmd)

	StartCmd.PersistentFlags().String(lspType, "jsonrpc://127.0.0.1:8000", "name and port for lsp service")
	_ = viper.BindPFlag(lspType, StartCmd.PersistentFlags().Lookup(lspType))
}

func createLspServer(t string, engine lsp.Engine) (server public.Server, err error) {
	switch {
	case strings.HasPrefix(t, jsonrpcType):
		log.Infof("creating jsonrpc lsp service for address %s", t)
		lspServer := jsonrpc.NewLspServer(engine)
		server = jsonrpc.NewPublicServer("tcp", strings.TrimPrefix(t, jsonrpcType), lspServer, &jsonrpc.DefaultAccessor{})

	default:
		return nil, fmt.Errorf("unknown lsp server type %s", t)
	}

	return server, err
}
