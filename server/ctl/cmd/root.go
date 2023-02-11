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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const logKey = "log"

// RootCmd -
var RootCmd = &cobra.Command{
	Use:          "server",
	Short:        "cli for proto-language-server",
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		log.SetFormatter(&log.JSONFormatter{})

		levelString, err := cmd.Flags().GetString(logKey)
		if err != nil {
			log.Info("no log flag passed in - using Info log level")
			levelString = "INFO"
		}

		level, err := log.ParseLevel(levelString)
		if err != nil {
			return fmt.Errorf("could not parse log level: %v", err)
		}
		log.SetLevel(level)

		return nil
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// Set default log level to error
	RootCmd.PersistentFlags().String(logKey, "info", "set log level")
}

func initConfig() {}
