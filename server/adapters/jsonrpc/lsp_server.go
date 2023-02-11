// Copyright 2023, Ed Kilgore
//
// This file is part of proto-language-server.
//
// proto-language-server is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option)  any later version.
//
// proto-language-server is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
// or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along with
// proto-language-server. If not, see <https://www.gnu.org/licenses/>.

// Package jsonrpc defines and implements types that interact with jsonrpc
package jsonrpc

import (
	"context"
	"fmt"

	"github.com/KobraKommander9/proto-language-server/server/ports/lsp"

	p "go.lsp.dev/protocol"
	"go.uber.org/zap"
)

// LspServer -
type LspServer struct {
	l      *zap.SugaredLogger
	engine lsp.Engine
}

// NewLspServer -
func NewLspServer(l *zap.SugaredLogger, engine lsp.Engine) *LspServer {
	return &LspServer{
		l:      l.Named("lsp_server"),
		engine: engine,
	}
}

// Initialize -
func (*LspServer) Initialize(context.Context, *p.InitializeParams) (*p.InitializeResult, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Initialized -
func (*LspServer) Initialized(context.Context, *p.InitializedParams) error {
	return fmt.Errorf("unimplemented")
}

// Shutdown -
func (*LspServer) Shutdown(context.Context) error {
	return fmt.Errorf("unimplemented")
}

// Exit -
func (*LspServer) Exit(context.Context) error {
	return fmt.Errorf("unimplemented")
}

// WorkDoneProgressCancel -
func (*LspServer) WorkDoneProgressCancel(ctx context.Context, params *p.WorkDoneProgressCancelParams) error {
	return fmt.Errorf("unimplemented")
}

// LogTrace -
func (*LspServer) LogTrace(ctx context.Context, params *p.LogTraceParams) error {
	return fmt.Errorf("unimplemented")
}

// SetTrace -
func (*LspServer) SetTrace(ctx context.Context, params *p.SetTraceParams) error {
	return fmt.Errorf("unimplemented")
}

// CodeAction -
func (*LspServer) CodeAction(ctx context.Context, params *p.CodeActionParams) ([]p.CodeAction, error) {
	return nil, fmt.Errorf("unimplemented")
}

// CodeLens -
func (*LspServer) CodeLens(ctx context.Context, params *p.CodeLensParams) ([]p.CodeLens, error) {
	return nil, fmt.Errorf("unimplemented")
}

// CodeLensResolve -
func (*LspServer) CodeLensResolve(ctx context.Context, params *p.CodeLens) (*p.CodeLens, error) {
	return nil, fmt.Errorf("unimplemented")
}

// ColorPresentation -
func (*LspServer) ColorPresentation(ctx context.Context, params *p.ColorPresentationParams) ([]p.ColorPresentation, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Completion -
func (*LspServer) Completion(ctx context.Context, params *p.CompletionParams) (*p.CompletionList, error) {
	return nil, fmt.Errorf("unimplemented")
}

// CompletionResolve -
func (*LspServer) CompletionResolve(ctx context.Context, params *p.CompletionItem) (*p.CompletionItem, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Declaration -
func (*LspServer) Declaration(ctx context.Context, params *p.DeclarationParams) ([]p.Location, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Definition -
func (*LspServer) Definition(ctx context.Context, params *p.DefinitionParams) ([]p.Location, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidChange -
func (*LspServer) DidChange(ctx context.Context, params *p.DidChangeTextDocumentParams) error {
	return fmt.Errorf("unimplemented")
}

// DidChangeConfiguration -
func (*LspServer) DidChangeConfiguration(ctx context.Context, params *p.DidChangeConfigurationParams) error {
	return fmt.Errorf("unimplemented")
}

// DidChangeWatchedFiles -
func (*LspServer) DidChangeWatchedFiles(ctx context.Context, params *p.DidChangeWatchedFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// DidChangeWorkspaceFolders -
func (*LspServer) DidChangeWorkspaceFolders(ctx context.Context, params *p.DidChangeWorkspaceFoldersParams) error {
	return fmt.Errorf("unimplemented")
}

// DidClose -
func (*LspServer) DidClose(ctx context.Context, params *p.DidCloseTextDocumentParams) error {
	return fmt.Errorf("unimplemented")
}

// DidOpen -
func (*LspServer) DidOpen(ctx context.Context, params *p.DidOpenTextDocumentParams) error {
	return fmt.Errorf("unimplemented")
}

// DidSave -
func (*LspServer) DidSave(ctx context.Context, params *p.DidSaveTextDocumentParams) error {
	return fmt.Errorf("unimplemented")
}

// DocumentColor -
func (*LspServer) DocumentColor(ctx context.Context, params *p.DocumentColorParams) ([]p.ColorInformation, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DocumentHighlight -
func (*LspServer) DocumentHighlight(ctx context.Context, params *p.DocumentHighlightParams) ([]p.DocumentHighlight, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DocumentLink -
func (*LspServer) DocumentLink(ctx context.Context, params *p.DocumentLinkParams) ([]p.DocumentLink, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DocumentLinkResolve -
func (*LspServer) DocumentLinkResolve(ctx context.Context, params *p.DocumentLink) (*p.DocumentLink, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DocumentSymbol -
func (*LspServer) DocumentSymbol(ctx context.Context, params *p.DocumentSymbolParams) ([]any, error) {
	return nil, fmt.Errorf("unimplemented")
}

// ExecuteCommand -
func (*LspServer) ExecuteCommand(ctx context.Context, params *p.ExecuteCommandParams) (any, error) {
	return nil, fmt.Errorf("unimplemented")
}

// FoldingRanges -
func (*LspServer) FoldingRanges(ctx context.Context, params *p.FoldingRangeParams) ([]p.FoldingRange, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Formatting -
func (*LspServer) Formatting(ctx context.Context, params *p.DocumentFormattingParams) ([]p.TextEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Hover -
func (*LspServer) Hover(ctx context.Context, params *p.HoverParams) (*p.Hover, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Implementation -
func (*LspServer) Implementation(ctx context.Context, params *p.ImplementationParams) ([]p.Location, error) {
	return nil, fmt.Errorf("unimplemented")
}

// OnTypeFormatting -
func (*LspServer) OnTypeFormatting(ctx context.Context, params *p.DocumentOnTypeFormattingParams) ([]p.TextEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// PrepareRename -
func (*LspServer) PrepareRename(ctx context.Context, params *p.PrepareRenameParams) (*p.Range, error) {
	return nil, fmt.Errorf("unimplemented")
}

// RangeFormatting -
func (*LspServer) RangeFormatting(ctx context.Context, params *p.DocumentRangeFormattingParams) ([]p.TextEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// References -
func (*LspServer) References(ctx context.Context, params *p.ReferenceParams) ([]p.Location, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Rename -
func (*LspServer) Rename(ctx context.Context, params *p.RenameParams) (*p.WorkspaceEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// SignatureHelp -
func (*LspServer) SignatureHelp(ctx context.Context, params *p.SignatureHelpParams) (*p.SignatureHelp, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Symbols -
func (*LspServer) Symbols(ctx context.Context, params *p.WorkspaceSymbolParams) ([]p.SymbolInformation, error) {
	return nil, fmt.Errorf("unimplemented")
}

// TypeDefinition -
func (*LspServer) TypeDefinition(ctx context.Context, params *p.TypeDefinitionParams) ([]p.Location, error) {
	return nil, fmt.Errorf("unimplemented")
}

// WillSave -
func (*LspServer) WillSave(ctx context.Context, params *p.WillSaveTextDocumentParams) error {
	return fmt.Errorf("unimplemented")
}

// WillSaveWaitUntil -
func (*LspServer) WillSaveWaitUntil(ctx context.Context, params *p.WillSaveTextDocumentParams) ([]p.TextEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// ShowDocument -
func (*LspServer) ShowDocument(ctx context.Context, params *p.ShowDocumentParams) (*p.ShowDocumentResult, error) {
	return nil, fmt.Errorf("unimplemented")
}

// WillCreateFiles -
func (*LspServer) WillCreateFiles(ctx context.Context, params *p.CreateFilesParams) (*p.WorkspaceEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidCreateFiles -
func (*LspServer) DidCreateFiles(ctx context.Context, params *p.CreateFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// WillRenameFiles -
func (*LspServer) WillRenameFiles(ctx context.Context, params *p.RenameFilesParams) (*p.WorkspaceEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidRenameFiles -
func (*LspServer) DidRenameFiles(ctx context.Context, params *p.RenameFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// WillDeleteFiles -
func (*LspServer) WillDeleteFiles(ctx context.Context, params *p.DeleteFilesParams) (*p.WorkspaceEdit, error) {
	return nil, fmt.Errorf("unimplemented")
}

// DidDeleteFiles -
func (*LspServer) DidDeleteFiles(ctx context.Context, params *p.DeleteFilesParams) error {
	return fmt.Errorf("unimplemented")
}

// CodeLensRefresh -
func (*LspServer) CodeLensRefresh(ctx context.Context) error {
	return fmt.Errorf("unimplemented")
}

// PrepareCallHierarchy -
func (*LspServer) PrepareCallHierarchy(ctx context.Context, params *p.CallHierarchyPrepareParams) ([]p.CallHierarchyItem, error) {
	return nil, fmt.Errorf("unimplemented")
}

// IncomingCalls -
func (*LspServer) IncomingCalls(ctx context.Context, params *p.CallHierarchyIncomingCallsParams) ([]p.CallHierarchyIncomingCall, error) {
	return nil, fmt.Errorf("unimplemented")
}

// OutgoingCalls -
func (*LspServer) OutgoingCalls(ctx context.Context, params *p.CallHierarchyOutgoingCallsParams) ([]p.CallHierarchyOutgoingCall, error) {
	return nil, fmt.Errorf("unimplemented")
}

// SemanticTokensFull -
func (*LspServer) SemanticTokensFull(ctx context.Context, params *p.SemanticTokensParams) (*p.SemanticTokens, error) {
	return nil, fmt.Errorf("unimplemented")
}

// SemanticTokensFullDelta -
func (*LspServer) SemanticTokensFullDelta(ctx context.Context, params *p.SemanticTokensDeltaParams) (any, error) {
	return nil, fmt.Errorf("unimplemented")
}

// SemanticTokensRange -
func (*LspServer) SemanticTokensRange(ctx context.Context, params *p.SemanticTokensRangeParams) (*p.SemanticTokens, error) {
	return nil, fmt.Errorf("unimplemented")
}

// SemanticTokensRefresh -
func (*LspServer) SemanticTokensRefresh(ctx context.Context) error {
	return fmt.Errorf("unimplemented")
}

// LinkedEditingRange -
func (*LspServer) LinkedEditingRange(ctx context.Context, params *p.LinkedEditingRangeParams) (*p.LinkedEditingRanges, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Moniker -
func (*LspServer) Moniker(ctx context.Context, params *p.MonikerParams) ([]p.Moniker, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Request -
func (*LspServer) Request(ctx context.Context, method string, params any) (any, error) {
	return nil, fmt.Errorf("unimplemented")
}
