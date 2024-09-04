// Code generated from Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Expression

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ExpressionParser.
type ExpressionVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExpressionParser#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#ComparisonExpr.
	VisitComparisonExpr(ctx *ComparisonExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#EqualsExpr.
	VisitEqualsExpr(ctx *EqualsExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#NotEqualsExpr.
	VisitNotEqualsExpr(ctx *NotEqualsExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#ContainsExpr.
	VisitContainsExpr(ctx *ContainsExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#BooleanOprand.
	VisitBooleanOprand(ctx *BooleanOprandContext) interface{}

	// Visit a parse tree produced by ExpressionParser#StringOprand.
	VisitStringOprand(ctx *StringOprandContext) interface{}
}
