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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const debugKey = "debug"

var logger *zap.Logger

// RootCmd -
var RootCmd = &cobra.Command{
	Use:          "server",
	Short:        "cli for proto-language-server",
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var logger *zap.Logger
		if viper.GetBool(debugKey) {
			logger, _ = zap.NewDevelopment()
		} else {
			logger, _ = zap.NewProduction(zap.WithCaller(false))
		}

		zap.ReplaceGlobals(logger)
		return nil
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().Bool(debugKey, false, "set logger to debug level")
	_ = viper.BindPFlag(debugKey, RootCmd.PersistentFlags().Lookup(debugKey))
}

func initConfig() {}
