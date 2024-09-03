// Code generated from Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Expression

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ExpressionParser struct {
	*antlr.BaseParser
}

var ExpressionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func expressionParserInit() {
	staticData := &ExpressionParserStaticData
	staticData.LiteralNames = []string{
		"", "'and'", "'or'", "'not'", "'='", "'!='", "'contains'", "", "", "",
		"'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "AND", "OR", "NOT", "EQUALS", "NOT_EQUALS", "CONTAINS", "BOOLEAN",
		"IDENTIFIER", "STRING", "LPAREN", "RPAREN", "WS",
	}
	staticData.RuleNames = []string{
		"expr", "comparison", "operand",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 43, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 15, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 5, 0, 23, 8, 0, 10, 0, 12, 0, 26, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 37, 8, 1, 1, 2, 1, 2, 3, 2, 41, 8, 2,
		1, 2, 0, 1, 0, 3, 0, 2, 4, 0, 0, 46, 0, 14, 1, 0, 0, 0, 2, 36, 1, 0, 0,
		0, 4, 40, 1, 0, 0, 0, 6, 7, 6, 0, -1, 0, 7, 8, 5, 3, 0, 0, 8, 15, 3, 0,
		0, 3, 9, 15, 3, 2, 1, 0, 10, 11, 5, 10, 0, 0, 11, 12, 3, 0, 0, 0, 12, 13,
		5, 11, 0, 0, 13, 15, 1, 0, 0, 0, 14, 6, 1, 0, 0, 0, 14, 9, 1, 0, 0, 0,
		14, 10, 1, 0, 0, 0, 15, 24, 1, 0, 0, 0, 16, 17, 10, 5, 0, 0, 17, 18, 5,
		1, 0, 0, 18, 23, 3, 0, 0, 6, 19, 20, 10, 4, 0, 0, 20, 21, 5, 2, 0, 0, 21,
		23, 3, 0, 0, 5, 22, 16, 1, 0, 0, 0, 22, 19, 1, 0, 0, 0, 23, 26, 1, 0, 0,
		0, 24, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 1, 1, 0, 0, 0, 26, 24, 1,
		0, 0, 0, 27, 28, 5, 8, 0, 0, 28, 29, 5, 4, 0, 0, 29, 37, 3, 4, 2, 0, 30,
		31, 5, 8, 0, 0, 31, 32, 5, 5, 0, 0, 32, 37, 3, 4, 2, 0, 33, 34, 5, 8, 0,
		0, 34, 35, 5, 6, 0, 0, 35, 37, 3, 4, 2, 0, 36, 27, 1, 0, 0, 0, 36, 30,
		1, 0, 0, 0, 36, 33, 1, 0, 0, 0, 37, 3, 1, 0, 0, 0, 38, 41, 5, 7, 0, 0,
		39, 41, 5, 9, 0, 0, 40, 38, 1, 0, 0, 0, 40, 39, 1, 0, 0, 0, 41, 5, 1, 0,
		0, 0, 5, 14, 22, 24, 36, 40,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ExpressionParserInit initializes any static state used to implement ExpressionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExpressionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExpressionParserInit() {
	staticData := &ExpressionParserStaticData
	staticData.once.Do(expressionParserInit)
}

// NewExpressionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExpressionParser(input antlr.TokenStream) *ExpressionParser {
	ExpressionParserInit()
	this := new(ExpressionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ExpressionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Expression.g4"

	return this
}

// ExpressionParser tokens.
const (
	ExpressionParserEOF        = antlr.TokenEOF
	ExpressionParserAND        = 1
	ExpressionParserOR         = 2
	ExpressionParserNOT        = 3
	ExpressionParserEQUALS     = 4
	ExpressionParserNOT_EQUALS = 5
	ExpressionParserCONTAINS   = 6
	ExpressionParserBOOLEAN    = 7
	ExpressionParserIDENTIFIER = 8
	ExpressionParserSTRING     = 9
	ExpressionParserLPAREN     = 10
	ExpressionParserRPAREN     = 11
	ExpressionParserWS         = 12
)

// ExpressionParser rules.
const (
	ExpressionParserRULE_expr       = 0
	ExpressionParserRULE_comparison = 1
	ExpressionParserRULE_operand    = 2
)

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AndExprContext struct {
	ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(ExpressionParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparisonExprContext struct {
	ExprContext
}

func NewComparisonExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonExprContext {
	var p = new(ComparisonExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ComparisonExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExprContext) Comparison() IComparisonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *ComparisonExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterComparisonExpr(s)
	}
}

func (s *ComparisonExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitComparisonExpr(s)
	}
}

func (s *ComparisonExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitComparisonExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(ExpressionParserNOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLPAREN, 0)
}

func (s *ParenExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(ExpressionParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExpressionParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ExpressionParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, ExpressionParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserNOT:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(7)
			p.Match(ExpressionParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(8)
			p.expr(3)
		}

	case ExpressionParserIDENTIFIER:
		localctx = NewComparisonExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(9)
			p.Comparison()
		}

	case ExpressionParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(10)
			p.Match(ExpressionParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(11)
			p.expr(0)
		}
		{
			p.SetState(12)
			p.Match(ExpressionParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(22)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expr)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(17)
					p.Match(ExpressionParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(18)
					p.expr(6)
				}

			case 2:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expr)
				p.SetState(19)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(20)
					p.Match(ExpressionParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(21)
					p.expr(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_comparison
	return p
}

func InitEmptyComparisonContext(p *ComparisonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_comparison
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) CopyAll(ctx *ComparisonContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EqualsExprContext struct {
	ComparisonContext
	op antlr.Token
}

func NewEqualsExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualsExprContext {
	var p = new(EqualsExprContext)

	InitEmptyComparisonContext(&p.ComparisonContext)
	p.parser = parser
	p.CopyAll(ctx.(*ComparisonContext))

	return p
}

func (s *EqualsExprContext) GetOp() antlr.Token { return s.op }

func (s *EqualsExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualsExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualsExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIDENTIFIER, 0)
}

func (s *EqualsExprContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *EqualsExprContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(ExpressionParserEQUALS, 0)
}

func (s *EqualsExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterEqualsExpr(s)
	}
}

func (s *EqualsExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitEqualsExpr(s)
	}
}

func (s *EqualsExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitEqualsExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContainsExprContext struct {
	ComparisonContext
	op antlr.Token
}

func NewContainsExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContainsExprContext {
	var p = new(ContainsExprContext)

	InitEmptyComparisonContext(&p.ComparisonContext)
	p.parser = parser
	p.CopyAll(ctx.(*ComparisonContext))

	return p
}

func (s *ContainsExprContext) GetOp() antlr.Token { return s.op }

func (s *ContainsExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ContainsExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainsExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIDENTIFIER, 0)
}

func (s *ContainsExprContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *ContainsExprContext) CONTAINS() antlr.TerminalNode {
	return s.GetToken(ExpressionParserCONTAINS, 0)
}

func (s *ContainsExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterContainsExpr(s)
	}
}

func (s *ContainsExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitContainsExpr(s)
	}
}

func (s *ContainsExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitContainsExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotEqualsExprContext struct {
	ComparisonContext
	op antlr.Token
}

func NewNotEqualsExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqualsExprContext {
	var p = new(NotEqualsExprContext)

	InitEmptyComparisonContext(&p.ComparisonContext)
	p.parser = parser
	p.CopyAll(ctx.(*ComparisonContext))

	return p
}

func (s *NotEqualsExprContext) GetOp() antlr.Token { return s.op }

func (s *NotEqualsExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *NotEqualsExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualsExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIDENTIFIER, 0)
}

func (s *NotEqualsExprContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *NotEqualsExprContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(ExpressionParserNOT_EQUALS, 0)
}

func (s *NotEqualsExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterNotEqualsExpr(s)
	}
}

func (s *NotEqualsExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitNotEqualsExpr(s)
	}
}

func (s *NotEqualsExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitNotEqualsExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExpressionParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExpressionParserRULE_comparison)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEqualsExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Match(ExpressionParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(28)

			var _m = p.Match(ExpressionParserEQUALS)

			localctx.(*EqualsExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.Operand()
		}

	case 2:
		localctx = NewNotEqualsExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Match(ExpressionParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)

			var _m = p.Match(ExpressionParserNOT_EQUALS)

			localctx.(*NotEqualsExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.Operand()
		}

	case 3:
		localctx = NewContainsExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.Match(ExpressionParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)

			var _m = p.Match(ExpressionParserCONTAINS)

			localctx.(*ContainsExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.Operand()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) CopyAll(ctx *OperandContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringOprandContext struct {
	OperandContext
}

func NewStringOprandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringOprandContext {
	var p = new(StringOprandContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *StringOprandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringOprandContext) STRING() antlr.TerminalNode {
	return s.GetToken(ExpressionParserSTRING, 0)
}

func (s *StringOprandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterStringOprand(s)
	}
}

func (s *StringOprandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitStringOprand(s)
	}
}

func (s *StringOprandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitStringOprand(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanOprandContext struct {
	OperandContext
}

func NewBooleanOprandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanOprandContext {
	var p = new(BooleanOprandContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *BooleanOprandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanOprandContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserBOOLEAN, 0)
}

func (s *BooleanOprandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterBooleanOprand(s)
	}
}

func (s *BooleanOprandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitBooleanOprand(s)
	}
}

func (s *BooleanOprandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitBooleanOprand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExpressionParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExpressionParserRULE_operand)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserBOOLEAN:
		localctx = NewBooleanOprandContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(ExpressionParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ExpressionParserSTRING:
		localctx = NewStringOprandContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Match(ExpressionParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ExpressionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ExpressionParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
