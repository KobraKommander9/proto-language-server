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

// Package lsp defines the language server protocol
package lsp

import (
	"context"
	"fmt"

	"github.com/KobraKommander9/proto-language-server/server/engine"

	p "go.lsp.dev/protocol"
	"go.uber.org/zap"
)

// Server -
type Server struct {
	l      *zap.SugaredLogger
	engine engine.LspEngine
}

// NewServer -
func NewServer(l *zap.SugaredLogger, engine engine.LspEngine) *Server {
	return &Server{
		l:      l.Named("lsp_server"),
		engine: engine,
	}
}

// Request -
func (*Server) Request(ctx context.Context, method string, params any) (any, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Declaration -
func (*Server) Declaration(ctx context.Context, params *p.DeclarationParams) ([]p.Location, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Definition -
func (*Server) Definition(ctx context.Context, params *p.DefinitionParams) ([]p.Location, error) {
	return nil, fmt.Errorf("unimplemented")
}

// TypeDefinition -
func (*Server) TypeDefinition(ctx context.Context, params *p.TypeDefinitionParams) ([]p.Location, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Implementation -
func (*Server) Implementation(ctx context.Context, params *p.ImplementationParams) ([]p.Location, error) {
	return nil, fmt.Errorf("unimplemented")
}

// References -
func (*Server) References(ctx context.Context, params *p.ReferenceParams) ([]p.Location, error) {
	return nil, fmt.Errorf("unimplemented")
}

// PrepareCallHierarchy -
func (*Server) PrepareCallHierarchy(ctx context.Context, params *p.CallHierarchyPrepareParams) ([]p.CallHierarchyItem, error) {
	return nil, fmt.Errorf("unimplemented")
}

// IncomingCalls -
func (*Server) IncomingCalls(ctx context.Context, params *p.CallHierarchyIncomingCallsParams) ([]p.CallHierarchyIncomingCall, error) {
	return nil, fmt.Errorf("unimplemented")
}

// OutgoingCalls -
func (*Server) OutgoingCalls(ctx context.Context, params *p.CallHierarchyOutgoingCallsParams) ([]p.CallHierarchyOutgoingCall, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DocumentHighlight -
func (*Server) DocumentHighlight(ctx context.Context, params *p.DocumentHighlightParams) ([]p.DocumentHighlight, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DocumentLink -
func (*Server) DocumentLink(ctx context.Context, params *p.DocumentLinkParams) ([]p.DocumentLink, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DocumentLinkResolve -
func (*Server) DocumentLinkResolve(ctx context.Context, params *p.DocumentLink) (*p.DocumentLink, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Hover -
func (*Server) Hover(ctx context.Context, params *p.HoverParams) (*p.Hover, error) {
	return nil, fmt.Errorf("unimplemented")
}

// CodeLens -
func (*Server) CodeLens(ctx context.Context, params *p.CodeLensParams) ([]p.CodeLens, error) {
	return nil, fmt.Errorf("unimplemented")
}

// CodeLensRefresh -
func (*Server) CodeLensRefresh(ctx context.Context) error {
	return fmt.Errorf("unimplemented")
}

// CodeLensResolve -
func (*Server) CodeLensResolve(ctx context.Context, params *p.CodeLens) (*p.CodeLens, error) {
	return nil, fmt.Errorf("unimplemented")
}

// FoldingRanges -
func (*Server) FoldingRanges(ctx context.Context, params *p.FoldingRangeParams) ([]p.FoldingRange, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DocumentSymbol -
func (*Server) DocumentSymbol(ctx context.Context, params *p.DocumentSymbolParams) ([]any, error) {
	return nil, fmt.Errorf("unimplemented")
}

// SemanticTokensFull -
func (*Server) SemanticTokensFull(ctx context.Context, params *p.SemanticTokensParams) (*p.SemanticTokens, error) {
	return nil, fmt.Errorf("unimplemented")
}

// SemanticTokensFullDelta -
func (*Server) SemanticTokensFullDelta(ctx context.Context, params *p.SemanticTokensDeltaParams) (any, error) {
	return nil, fmt.Errorf("unimplemented")
}

// SemanticTokensRange -
func (*Server) SemanticTokensRange(ctx context.Context, params *p.SemanticTokensRangeParams) (*p.SemanticTokens, error) {
	return nil, fmt.Errorf("unimplemented")
}

// SemanticTokensRefresh -
func (*Server) SemanticTokensRefresh(ctx context.Context) error {
	return fmt.Errorf("unimplemented")
}

// Moniker -
func (*Server) Moniker(ctx context.Context, params *p.MonikerParams) ([]p.Moniker, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Completion -
func (*Server) Completion(ctx context.Context, params *p.CompletionParams) (*p.CompletionList, error) {
	return nil, fmt.Errorf("unimplemented")
}

// CompletionResolve -
func (*Server) CompletionResolve(ctx context.Context, params *p.CompletionItem) (*p.CompletionItem, error) {
	return nil, fmt.Errorf("unimplemented")
}

// SignatureHelp -
func (*Server) SignatureHelp(ctx context.Context, params *p.SignatureHelpParams) (*p.SignatureHelp, error) {
	return nil, fmt.Errorf("unimplemented")
}

// CodeAction -
func (*Server) CodeAction(ctx context.Context, params *p.CodeActionParams) ([]p.CodeAction, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DocumentColor -
func (*Server) DocumentColor(ctx context.Context, params *p.DocumentColorParams) ([]p.ColorInformation, error) {
	return nil, fmt.Errorf("unimplemented")
}

// ColorPresentation -
func (*Server) ColorPresentation(ctx context.Context, params *p.ColorPresentationParams) ([]p.ColorPresentation, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Formatting -
func (*Server) Formatting(ctx context.Context, params *p.DocumentFormattingParams) ([]p.TextEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// RangeFormatting -
func (*Server) RangeFormatting(ctx context.Context, params *p.DocumentRangeFormattingParams) ([]p.TextEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// OnTypeFormatting -
func (*Server) OnTypeFormatting(ctx context.Context, params *p.DocumentOnTypeFormattingParams) ([]p.TextEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Rename -
func (*Server) Rename(ctx context.Context, params *p.RenameParams) (*p.WorkspaceEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// PrepareRename -
func (*Server) PrepareRename(ctx context.Context, params *p.PrepareRenameParams) (*p.Range, error) {
	return nil, fmt.Errorf("unimplemented")
}

// LinkedEditingRange -
func (*Server) LinkedEditingRange(ctx context.Context, params *p.LinkedEditingRangeParams) (*p.LinkedEditingRanges, error) {
	return nil, fmt.Errorf("unimplemented")
}
