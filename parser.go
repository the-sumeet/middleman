package main

// import (
// 	"fmt"
// 	"io"
// 	"middleman/parser"
// 	"net/http"
// 	"strings"

// 	"github.com/antlr4-go/antlr/v4"
// )

// const (
// 	BOOL = iota
// 	STRING
// 	IDENTIFIER
// 	// Operators
// 	EQUALS
// 	NOT_EQUALS
// 	CONTAINS
// 	// Identifiers
// 	URL
// 	HEADER
// 	QUERY
// 	BODY
// )

// type Expression struct {
// 	Op    string
// 	Value any
// 	Type  int
// }

// type CustomErrorListener struct {
// 	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
// 	Errors                      []error
// }

// func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
// 	c.Errors = append(c.Errors, fmt.Errorf("line %d:%d %s", line, column, msg))
// }

// type MyVisitor struct {
// 	Request http.Request
// 	parser.BaseExpressionVisitor
// }

// func (v *MyVisitor) Visit(tree antlr.Tree) any {
// 	switch val := tree.(type) {

// 	case *parser.AndExprContext:
// 		return v.VisitAndExpr(val)
// 	case *parser.OrExprContext:
// 		return v.VisitOrExpr(val)

// 	// Coparisons
// 	case *parser.ComparisonExprContext:
// 		return v.VisitComparisonExpr(val)
// 	case *parser.EqualsExprContext:
// 		return v.VisitEqualsExpr(val)
// 	case *parser.NotEqualsExprContext:
// 		return v.VisitNotEqualsExpr(val)
// 	case *parser.ContainsExprContext:
// 		return v.VisitContainsExpr(val)

// 	// Oprand
// 	case *parser.OperandContext:
// 		return v.OperandContext(val)
// 	case *parser.BooleanOprandContext:
// 		return v.VisitBooleanOprand(val)
// 	case *parser.StringOprandContext:
// 		return v.VisitStringOprand(val)
// 	default:
// 		panic("Unknown context")
// 	}
// }

// func (v *MyVisitor) getIdentifier(identifier string) int {

// 	switch identifier {
// 	case "url":
// 		return URL
// 	case "header":
// 		return HEADER
// 	case "query":
// 		return QUERY
// 	case "body":
// 		return BODY
// 	default:
// 		return URL
// 	}
// }

// func (v *MyVisitor) evaluateComparison(op int, identifier int, oprand string) bool {

// 	switch op {
// 	case EQUALS:
// 		return compareEuqual(v.Request, identifier, oprand)
// 	case NOT_EQUALS:
// 		return compareNotEuqual(v.Request, identifier, oprand)
// 	case CONTAINS:
// 		return compareContains(v.Request, identifier, oprand)
// 	default:
// 		panic("Unknown operator")
// 	}
// }

// // Comparisons
// func (v *MyVisitor) VisitComparisonExpr(ctx *parser.ComparisonExprContext) bool {
// 	child := v.Visit(ctx.GetChild(0))
// 	childBool := child.(bool)
// 	return childBool
// }
// func (v *MyVisitor) VisitEqualsExpr(ctx *parser.EqualsExprContext) bool {
// 	op := EQUALS
// 	identifier := v.getIdentifier(ctx.IDENTIFIER().GetText())
// 	oprand := v.Visit(ctx.GetChild(2))
// 	oprandStr := oprand.(string)
// 	result := v.evaluateComparison(op, identifier, oprandStr)
// 	return result
// }
// func (v *MyVisitor) VisitNotEqualsExpr(ctx *parser.NotEqualsExprContext) bool {
// 	op := NOT_EQUALS
// 	identifier := v.getIdentifier(ctx.IDENTIFIER().GetText())
// 	oprand := v.Visit(ctx.GetChild(2))
// 	oprandStr := oprand.(string)
// 	result := v.evaluateComparison(op, identifier, oprandStr)
// 	return result
// }
// func (v *MyVisitor) VisitContainsExpr(ctx *parser.ContainsExprContext) bool {
// 	op := CONTAINS
// 	identifier := v.getIdentifier(ctx.IDENTIFIER().GetText())
// 	oprand := v.Visit(ctx.GetChild(2))
// 	oprandStr := oprand.(string)
// 	result := v.evaluateComparison(op, identifier, oprandStr)
// 	return result
// }

// // Oprand
// func (v *MyVisitor) OperandContext(ctx *parser.OperandContext) string {
// 	return ctx.GetText()
// }
// func (v *MyVisitor) VisitOprandExpr(ctx *parser.ComparisonExprContext) Expression {
// 	identifier := v.Visit(ctx.GetChild(1))
// 	op := v.Visit(ctx.GetChild(0))
// 	oprand := v.Visit(ctx.GetChild(2))

// 	fmt.Println(identifier, op, oprand)

// 	return Expression{}
// }
// func (v *MyVisitor) VisitBooleanOprand(ctx *parser.BooleanOprandContext) Expression {
// 	val := ctx.GetText()
// 	return Expression{
// 		Op:    "",
// 		Value: val == "true",
// 		Type:  BOOL,
// 	}
// }

// // And
// func (v *MyVisitor) VisitAndExpr(ctx *parser.AndExprContext) bool {
// 	left := v.Visit(ctx.GetChild(0))
// 	right := v.Visit(ctx.GetChild(2))
// 	return left.(bool) && right.(bool)
// }

// // Or
// func (v *MyVisitor) VisitOrExpr(ctx *parser.OrExprContext) bool {
// 	left := v.Visit(ctx.GetChild(0))
// 	right := v.Visit(ctx.GetChild(2))
// 	return left.(bool) || right.(bool)
// }

// func (v *MyVisitor) VisitStringOprand(ctx *parser.StringOprandContext) string {
// 	stringVal := ctx.GetText()
// 	return stringVal[1 : len(stringVal)-1]
// }

// func getRequestValue(request http.Request, identifier int) any {
// 	switch identifier {
// 	case URL:
// 		return request.URL.Host
// 	case HEADER:
// 		return request.Header
// 	case QUERY:
// 		return ""
// 	case BODY:
// 		body, err := io.ReadAll(request.Body)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return string(body)
// 	default:
// 		return ""
// 	}
// }

// func compareEuqual(request http.Request, identifier int, oprand string) bool {

// 	switch identifier {
// 	case URL:
// 		return request.URL.String() == oprand
// 	case HEADER:
// 		header := getRequestValue(request, HEADER).(http.Header)
// 		// Split with color
// 		headerIn := strings.Split(oprand, ":")
// 		if len(headerIn) != 2 {
// 			return false
// 		}
// 		headerKey, headerValue := headerIn[0], headerIn[1]
// 		headerValues := header.Values(headerKey)
// 		for _, value := range headerValues {
// 			if value == headerValue {
// 				return true
// 			}
// 		}
// 		return false
// 	case QUERY:
// 		return false
// 	case BODY:
// 		requestValue := getRequestValue(request, BODY).(string)
// 		return requestValue == oprand
// 	default:
// 		return false
// 	}
// }

// func compareNotEuqual(request http.Request, identifier int, oprand string) bool {

// 	switch identifier {
// 	case URL:
// 		return request.URL.String() != oprand
// 	case HEADER:
// 		// Return true if any of the header values is not equal to the oprand
// 		header := getRequestValue(request, BODY).(http.Header)
// 		// Split with color
// 		headerIn := strings.Split(oprand, ":")
// 		if len(headerIn) != 2 {
// 			return false
// 		}
// 		headerKey, headerValue := headerIn[0], headerIn[1]
// 		headerValues := header.Values(headerKey)
// 		for _, value := range headerValues {
// 			if value != headerValue {
// 				return true
// 			}
// 		}
// 		return false
// 	case QUERY:
// 		return false
// 	case BODY:
// 		requestValue := getRequestValue(request, BODY).(string)
// 		return requestValue != oprand
// 	default:
// 		return false
// 	}
// }

// func compareContains(request http.Request, identifier int, oprand string) bool {

// 	switch identifier {
// 	case URL:
// 		return strings.Contains(request.URL.String(), oprand)
// 	case HEADER:
// 		// Return true if any of the header values is not equal to the oprand
// 		header := getRequestValue(request, BODY).(http.Header)
// 		// Split with color
// 		headerIn := strings.Split(oprand, ":")
// 		if len(headerIn) != 2 {
// 			return false
// 		}
// 		headerKey, headerValue := headerIn[0], headerIn[1]
// 		headerValues := header.Values(headerKey)
// 		for _, value := range headerValues {
// 			if strings.Contains(value, headerValue) {
// 				return true
// 			}
// 		}
// 		return false
// 	case QUERY:
// 		return false
// 	case BODY:
// 		requestValue := getRequestValue(request, BODY).(string)
// 		return strings.Contains(requestValue, oprand)
// 	default:
// 		return false
// 	}
// }

// func validateString(expr string) []error {
// 	input := antlr.NewInputStream(expr)
// 	lexer := parser.NewExpressionLexer(input)
// 	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

// 	parser := parser.NewExpressionParser(stream)
// 	parser.RemoveErrorListeners() // Remove default listeners
// 	errorListener := CustomErrorListener{}
// 	parser.AddErrorListener(&errorListener)

// 	_ = parser.Expr()

// 	if errorListener.Errors != nil {
// 		return errorListener.Errors
// 	}
// 	return nil
// }

// func parse(expr string, request http.Request) (bool, []error) {
// 	input := antlr.NewInputStream(expr)
// 	lexer := parser.NewExpressionLexer(input)
// 	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

// 	parser := parser.NewExpressionParser(stream)
// 	parser.RemoveErrorListeners() // Remove default listeners
// 	errorListener := CustomErrorListener{}
// 	parser.AddErrorListener(&errorListener)

// 	tree := parser.Expr()

// 	if errorListener.Errors != nil {
// 		return false, errorListener.Errors
// 	}

// 	visitor := &MyVisitor{Request: request}
// 	result := visitor.Visit(tree)
// 	resultBool := result.(bool)
// 	return resultBool, nil
// }
