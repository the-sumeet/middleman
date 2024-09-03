package main

import (
	"fmt"
	"middleman/parser"
	"net/http"

	"github.com/antlr4-go/antlr/v4"
)

const (
	BOOL = iota
	STRING
	IDENTIFIER
	// Operators
	EQUALS
	NOT_EQUALS
	CONTAINS
)

type Expression struct {
	Op    string
	Value any
	Type  int
}

type MyVisitor struct {
	Request http.Request
	parser.BaseExpressionVisitor
}

func (v *MyVisitor) Visit(tree antlr.Tree) any {
	switch val := tree.(type) {

	// Coparisons
	case *parser.ComparisonExprContext:
		return v.VisitComparisonExpr(val)
	case *parser.EqualsExprContext:
		return v.VisitEqualsExpr(val)
	case *parser.NotEqualsExprContext:
		return v.VisitNotEqualsExpr(val)
	case *parser.ContainsExprContext:
		return v.VisitContainsExpr(val)
	// Oprand
	case *parser.BooleanOprandContext:
		return v.VisitBooleanOprand(val)
	case *parser.StringOprandContext:
		return v.VisitStringOprand(val)
	default:
		panic("Unknown context")
	}
}

// Comparisons
func (v *MyVisitor) VisitComparisonExpr(ctx *parser.ComparisonExprContext) bool {
	// op := v.Visit(ctx.GetChild(0)).(int)
	// identifier := v.Visit(ctx.GetChild(1))
	// oprand := v.Visit(ctx.GetChild(2))
	// fmt.Println(identifier, op, oprand)

	// Visit children
	for _, child := range ctx.GetChildren() {
		v.Visit(child)
	}
	return true
}

func (v *MyVisitor) VisitEqualsExpr(ctx *parser.EqualsExprContext) int {
	op := EQUALS
	identifier := ctx.IDENTIFIER().GetText()
	oprand := v.Visit(ctx.GetChild(2))
	fmt.Println(identifier, op, oprand)
	return EQUALS
}

func (v *MyVisitor) VisitNotEqualsExpr(ctx *parser.NotEqualsExprContext) int {
	return EQUALS
}

func (v *MyVisitor) VisitContainsExpr(ctx *parser.ContainsExprContext) int {
	return EQUALS
}

func (v *MyVisitor) VisitOprandExpr(ctx *parser.ComparisonExprContext) Expression {
	identifier := v.Visit(ctx.GetChild(1))
	op := v.Visit(ctx.GetChild(0))
	oprand := v.Visit(ctx.GetChild(2))

	fmt.Println(identifier, op, oprand)

	return Expression{}
}

func (v *MyVisitor) VisitBooleanOprand(ctx *parser.BooleanOprandContext) Expression {
	val := ctx.GetText()
	return Expression{
		Op:    "",
		Value: val == "true",
		Type:  BOOL,
	}
}

func (v *MyVisitor) VisitStringOprand(ctx *parser.StringOprandContext) Expression {
	stringVal := ctx.GetText()
	val := stringVal[1 : len(stringVal)-1]
	return Expression{
		Op:    "",
		Value: val,
		Type:  STRING,
	}
}

func parse(expr string) any {
	input := antlr.NewInputStream(expr)
	lexer := parser.NewBooleanExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewExpressionParser(stream)
	tree := parser.Expr()

	visitor := &MyVisitor{}
	result := visitor.Visit(tree)
	return result
}
