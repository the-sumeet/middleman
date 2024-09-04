// Code generated from Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Expression

import "github.com/antlr4-go/antlr/v4"

// ExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type ExpressionListener interface {
	antlr.ParseTreeListener

	// EnterAndExpr is called when entering the AndExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterComparisonExpr is called when entering the ComparisonExpr production.
	EnterComparisonExpr(c *ComparisonExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterOrExpr is called when entering the OrExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterEqualsExpr is called when entering the EqualsExpr production.
	EnterEqualsExpr(c *EqualsExprContext)

	// EnterNotEqualsExpr is called when entering the NotEqualsExpr production.
	EnterNotEqualsExpr(c *NotEqualsExprContext)

	// EnterContainsExpr is called when entering the ContainsExpr production.
	EnterContainsExpr(c *ContainsExprContext)

	// EnterBooleanOprand is called when entering the BooleanOprand production.
	EnterBooleanOprand(c *BooleanOprandContext)

	// EnterStringOprand is called when entering the StringOprand production.
	EnterStringOprand(c *StringOprandContext)

	// ExitAndExpr is called when exiting the AndExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitComparisonExpr is called when exiting the ComparisonExpr production.
	ExitComparisonExpr(c *ComparisonExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitOrExpr is called when exiting the OrExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitEqualsExpr is called when exiting the EqualsExpr production.
	ExitEqualsExpr(c *EqualsExprContext)

	// ExitNotEqualsExpr is called when exiting the NotEqualsExpr production.
	ExitNotEqualsExpr(c *NotEqualsExprContext)

	// ExitContainsExpr is called when exiting the ContainsExpr production.
	ExitContainsExpr(c *ContainsExprContext)

	// ExitBooleanOprand is called when exiting the BooleanOprand production.
	ExitBooleanOprand(c *BooleanOprandContext)

	// ExitStringOprand is called when exiting the StringOprand production.
	ExitStringOprand(c *StringOprandContext)
}
