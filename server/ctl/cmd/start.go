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

import "github.com/spf13/cobra"

// StartCmd -
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long: `starts the omni api service.
	Example: ./bin/apictl start`,
	RunE: func(_ *cobra.Command, _ []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(StartCmd)
}
