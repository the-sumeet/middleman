// Code generated from Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Expression

import "github.com/antlr4-go/antlr/v4"

type BaseExpressionVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExpressionVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitComparisonExpr(ctx *ComparisonExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitEqualsExpr(ctx *EqualsExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitNotEqualsExpr(ctx *NotEqualsExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitContainsExpr(ctx *ContainsExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitBooleanOprand(ctx *BooleanOprandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitStringOprand(ctx *StringOprandContext) interface{} {
	return v.VisitChildren(ctx)
}
