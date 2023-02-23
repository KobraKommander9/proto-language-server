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
	"fmt"
	"sync"

	"go.uber.org/atomic"
	"go.uber.org/zap"
)

// Engine -
type Engine struct {
	EngineInfo
	l           *zap.SugaredLogger
	lock        *sync.Mutex
	config      *EngineConfig
	initialized *atomic.Bool
	shutdown    *atomic.Bool
}

// NewEngine -
func NewEngine(l *zap.SugaredLogger, info EngineInfo) *Engine {
	return &Engine{
		EngineInfo:  info,
		l:           l.Named("engine"),
		lock:        &sync.Mutex{},
		initialized: atomic.NewBool(false),
		shutdown:    atomic.NewBool(false),
	}
}

// Close -
func (e *Engine) Close() error {
	e.l.Infof("closing engine...")

	if e.shutdown.Load() {
		e.l.Infof("successfully closed engine")
		return nil
	}

	e.l.Warnf("failed to shutdown engine before closing")
	return fmt.Errorf("engine wasn't shutdown before closing")
}
