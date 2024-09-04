// Code generated from Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ExpressionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ExpressionLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func expressionlexerLexerInit() {
	staticData := &ExpressionLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'and'", "'or'", "'not'", "'='", "'!='", "'contains'", "", "", "",
		"'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "AND", "OR", "NOT", "EQUALS", "NOT_EQUALS", "CONTAINS", "BOOLEAN",
		"IDENTIFIER", "STRING", "LPAREN", "RPAREN", "WS",
	}
	staticData.RuleNames = []string{
		"AND", "OR", "NOT", "EQUALS", "NOT_EQUALS", "CONTAINS", "BOOLEAN", "IDENTIFIER",
		"STRING", "LPAREN", "RPAREN", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 96, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 60, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 75, 8, 7, 1, 8, 1, 8, 5, 8, 79,
		8, 8, 10, 8, 12, 8, 82, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		11, 4, 11, 91, 8, 11, 11, 11, 12, 11, 92, 1, 11, 1, 11, 0, 0, 12, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
		12, 1, 0, 12, 2, 0, 66, 66, 98, 98, 2, 0, 79, 79, 111, 111, 2, 0, 68, 68,
		100, 100, 2, 0, 89, 89, 121, 121, 2, 0, 85, 85, 117, 117, 2, 0, 82, 82,
		114, 114, 2, 0, 76, 76, 108, 108, 2, 0, 72, 72, 104, 104, 2, 0, 69, 69,
		101, 101, 2, 0, 65, 65, 97, 97, 3, 0, 10, 10, 13, 13, 34, 34, 3, 0, 9,
		10, 13, 13, 32, 32, 100, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 1, 25, 1, 0, 0, 0, 3, 29, 1, 0, 0, 0, 5,
		32, 1, 0, 0, 0, 7, 36, 1, 0, 0, 0, 9, 38, 1, 0, 0, 0, 11, 41, 1, 0, 0,
		0, 13, 59, 1, 0, 0, 0, 15, 74, 1, 0, 0, 0, 17, 76, 1, 0, 0, 0, 19, 85,
		1, 0, 0, 0, 21, 87, 1, 0, 0, 0, 23, 90, 1, 0, 0, 0, 25, 26, 5, 97, 0, 0,
		26, 27, 5, 110, 0, 0, 27, 28, 5, 100, 0, 0, 28, 2, 1, 0, 0, 0, 29, 30,
		5, 111, 0, 0, 30, 31, 5, 114, 0, 0, 31, 4, 1, 0, 0, 0, 32, 33, 5, 110,
		0, 0, 33, 34, 5, 111, 0, 0, 34, 35, 5, 116, 0, 0, 35, 6, 1, 0, 0, 0, 36,
		37, 5, 61, 0, 0, 37, 8, 1, 0, 0, 0, 38, 39, 5, 33, 0, 0, 39, 40, 5, 61,
		0, 0, 40, 10, 1, 0, 0, 0, 41, 42, 5, 99, 0, 0, 42, 43, 5, 111, 0, 0, 43,
		44, 5, 110, 0, 0, 44, 45, 5, 116, 0, 0, 45, 46, 5, 97, 0, 0, 46, 47, 5,
		105, 0, 0, 47, 48, 5, 110, 0, 0, 48, 49, 5, 115, 0, 0, 49, 12, 1, 0, 0,
		0, 50, 51, 5, 116, 0, 0, 51, 52, 5, 114, 0, 0, 52, 53, 5, 117, 0, 0, 53,
		60, 5, 101, 0, 0, 54, 55, 5, 102, 0, 0, 55, 56, 5, 97, 0, 0, 56, 57, 5,
		108, 0, 0, 57, 58, 5, 115, 0, 0, 58, 60, 5, 101, 0, 0, 59, 50, 1, 0, 0,
		0, 59, 54, 1, 0, 0, 0, 60, 14, 1, 0, 0, 0, 61, 62, 7, 0, 0, 0, 62, 63,
		7, 1, 0, 0, 63, 64, 7, 2, 0, 0, 64, 75, 7, 3, 0, 0, 65, 66, 7, 4, 0, 0,
		66, 67, 7, 5, 0, 0, 67, 75, 7, 6, 0, 0, 68, 69, 7, 7, 0, 0, 69, 70, 7,
		8, 0, 0, 70, 71, 7, 9, 0, 0, 71, 72, 7, 2, 0, 0, 72, 73, 7, 8, 0, 0, 73,
		75, 7, 5, 0, 0, 74, 61, 1, 0, 0, 0, 74, 65, 1, 0, 0, 0, 74, 68, 1, 0, 0,
		0, 75, 16, 1, 0, 0, 0, 76, 80, 5, 34, 0, 0, 77, 79, 8, 10, 0, 0, 78, 77,
		1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0,
		81, 83, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 84, 5, 34, 0, 0, 84, 18, 1,
		0, 0, 0, 85, 86, 5, 40, 0, 0, 86, 20, 1, 0, 0, 0, 87, 88, 5, 41, 0, 0,
		88, 22, 1, 0, 0, 0, 89, 91, 7, 11, 0, 0, 90, 89, 1, 0, 0, 0, 91, 92, 1,
		0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94,
		95, 6, 11, 0, 0, 95, 24, 1, 0, 0, 0, 5, 0, 59, 74, 80, 92, 1, 6, 0, 0,
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

// ExpressionLexerInit initializes any static state used to implement ExpressionLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewExpressionLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExpressionLexerInit() {
	staticData := &ExpressionLexerLexerStaticData
	staticData.once.Do(expressionlexerLexerInit)
}

// NewExpressionLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewExpressionLexer(input antlr.CharStream) *ExpressionLexer {
	ExpressionLexerInit()
	l := new(ExpressionLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ExpressionLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Expression.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExpressionLexer tokens.
const (
	ExpressionLexerAND        = 1
	ExpressionLexerOR         = 2
	ExpressionLexerNOT        = 3
	ExpressionLexerEQUALS     = 4
	ExpressionLexerNOT_EQUALS = 5
	ExpressionLexerCONTAINS   = 6
	ExpressionLexerBOOLEAN    = 7
	ExpressionLexerIDENTIFIER = 8
	ExpressionLexerSTRING     = 9
	ExpressionLexerLPAREN     = 10
	ExpressionLexerRPAREN     = 11
	ExpressionLexerWS         = 12
)
