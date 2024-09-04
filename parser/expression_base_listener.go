// Code generated from Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Expression

import "github.com/antlr4-go/antlr/v4"

// BaseExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type BaseExpressionListener struct{}

var _ ExpressionListener = &BaseExpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAndExpr is called when production AndExpr is entered.
func (s *BaseExpressionListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production AndExpr is exited.
func (s *BaseExpressionListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterComparisonExpr is called when production ComparisonExpr is entered.
func (s *BaseExpressionListener) EnterComparisonExpr(ctx *ComparisonExprContext) {}

// ExitComparisonExpr is called when production ComparisonExpr is exited.
func (s *BaseExpressionListener) ExitComparisonExpr(ctx *ComparisonExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseExpressionListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseExpressionListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseExpressionListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseExpressionListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterOrExpr is called when production OrExpr is entered.
func (s *BaseExpressionListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production OrExpr is exited.
func (s *BaseExpressionListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterEqualsExpr is called when production EqualsExpr is entered.
func (s *BaseExpressionListener) EnterEqualsExpr(ctx *EqualsExprContext) {}

// ExitEqualsExpr is called when production EqualsExpr is exited.
func (s *BaseExpressionListener) ExitEqualsExpr(ctx *EqualsExprContext) {}

// EnterNotEqualsExpr is called when production NotEqualsExpr is entered.
func (s *BaseExpressionListener) EnterNotEqualsExpr(ctx *NotEqualsExprContext) {}

// ExitNotEqualsExpr is called when production NotEqualsExpr is exited.
func (s *BaseExpressionListener) ExitNotEqualsExpr(ctx *NotEqualsExprContext) {}

// EnterContainsExpr is called when production ContainsExpr is entered.
func (s *BaseExpressionListener) EnterContainsExpr(ctx *ContainsExprContext) {}

// ExitContainsExpr is called when production ContainsExpr is exited.
func (s *BaseExpressionListener) ExitContainsExpr(ctx *ContainsExprContext) {}

// EnterBooleanOprand is called when production BooleanOprand is entered.
func (s *BaseExpressionListener) EnterBooleanOprand(ctx *BooleanOprandContext) {}

// ExitBooleanOprand is called when production BooleanOprand is exited.
func (s *BaseExpressionListener) ExitBooleanOprand(ctx *BooleanOprandContext) {}

// EnterStringOprand is called when production StringOprand is entered.
func (s *BaseExpressionListener) EnterStringOprand(ctx *StringOprandContext) {}

// ExitStringOprand is called when production StringOprand is exited.
func (s *BaseExpressionListener) ExitStringOprand(ctx *StringOprandContext) {}
