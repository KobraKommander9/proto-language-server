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
	"io"

	"github.com/KobraKommander9/proto-language-server/server/ports/filesystem"

	"go.uber.org/zap"
)

// Clients -
type Clients struct {
	Filesystem filesystem.Client
}

// Close -
func (c *Clients) Close(l *zap.SugaredLogger) {
	c.logClose(l, "filesystem")
}

func (*Clients) logClose(l *zap.SugaredLogger, name string) {
	l.Infof("Closed %s client", name)
}

func (*Clients) closeClient(l *zap.SugaredLogger, name string, client io.Closer) {
	l.Infof("Closing %s client...", name)
	if err := client.Close(); err != nil {
		l.Warnf("Error closing %s client: %v", name, err)
	} else {
		l.Infof("Closed %s client", name)
	}
}
