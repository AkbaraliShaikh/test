// Code generated from MyRule.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // MyRule

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MyRuleParser struct {
	*antlr.BaseParser
}

var myruleParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func myruleParserInit() {
	staticData := &myruleParserStaticData
	staticData.literalNames = []string{
		"", "'in-time-range'", "','", "'='", "'<'", "'>'", "'<='", "'>='", "'not='",
		"'in'", "'not-in'", "'has-all'", "'has-none'", "'has-any'", "'['", "']'",
		"'('", "')'", "'and'", "'or'", "'newer-than'", "'newer-than-or-equal-to'",
		"'older-than'", "'older-than-or-equal-to'", "'equals-version'", "'within-any-of'",
		"", "", "", "'true'", "'false'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "EQ", "LT", "GT", "LTEQ", "GTEQ", "NEQ", "IN", "NIN", "HALL",
		"HNONE", "HANY", "OBRACKET", "CBRACKET", "OPARAM", "CPARAM", "AND",
		"OR", "VNT", "VNTE", "VOT", "VOTE", "VE", "WAO", "INT", "FLOAT", "STRING",
		"TRUE", "FALSE", "SPACE", "WS", "ATTRNAME",
	}
	staticData.ruleNames = []string{
		"query", "leftexpr", "listStrings", "stringElements", "stringElement",
		"listInts", "intElements", "intElement", "listOfLists", "list", "floatElement",
		"listOfTimeLists", "timeList", "timeStringElement",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 33, 194, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 74, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 5, 0, 82, 8, 0, 10, 0, 12, 0, 85, 9, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 96, 8, 3, 10, 3, 12, 3, 99, 9, 3, 1, 3,
		5, 3, 102, 8, 3, 10, 3, 12, 3, 105, 9, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 116, 8, 6, 10, 6, 12, 6, 119, 9, 6, 1, 6,
		5, 6, 122, 8, 6, 10, 6, 12, 6, 125, 9, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 5, 8, 133, 8, 8, 10, 8, 12, 8, 136, 9, 8, 1, 8, 5, 8, 139, 8, 8,
		10, 8, 12, 8, 142, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 150,
		8, 9, 10, 9, 12, 9, 153, 9, 9, 1, 9, 5, 9, 156, 8, 9, 10, 9, 12, 9, 159,
		9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 169,
		8, 11, 10, 11, 12, 11, 172, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 5,
		12, 179, 8, 12, 10, 12, 12, 12, 182, 9, 12, 1, 12, 5, 12, 185, 8, 12, 10,
		12, 12, 12, 188, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 0, 1, 0, 14,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0, 6, 2, 0, 3, 3, 8,
		8, 1, 0, 4, 7, 1, 0, 9, 10, 1, 0, 11, 13, 1, 0, 20, 24, 1, 0, 26, 27, 202,
		0, 73, 1, 0, 0, 0, 2, 86, 1, 0, 0, 0, 4, 88, 1, 0, 0, 0, 6, 92, 1, 0, 0,
		0, 8, 106, 1, 0, 0, 0, 10, 108, 1, 0, 0, 0, 12, 112, 1, 0, 0, 0, 14, 126,
		1, 0, 0, 0, 16, 128, 1, 0, 0, 0, 18, 145, 1, 0, 0, 0, 20, 162, 1, 0, 0,
		0, 22, 164, 1, 0, 0, 0, 24, 175, 1, 0, 0, 0, 26, 191, 1, 0, 0, 0, 28, 29,
		6, 0, -1, 0, 29, 30, 5, 16, 0, 0, 30, 31, 3, 0, 0, 0, 31, 32, 5, 17, 0,
		0, 32, 74, 1, 0, 0, 0, 33, 34, 3, 2, 1, 0, 34, 35, 7, 0, 0, 0, 35, 36,
		5, 26, 0, 0, 36, 74, 1, 0, 0, 0, 37, 38, 3, 2, 1, 0, 38, 39, 7, 0, 0, 0,
		39, 40, 5, 28, 0, 0, 40, 74, 1, 0, 0, 0, 41, 42, 3, 2, 1, 0, 42, 43, 7,
		1, 0, 0, 43, 44, 5, 26, 0, 0, 44, 74, 1, 0, 0, 0, 45, 46, 3, 2, 1, 0, 46,
		47, 7, 2, 0, 0, 47, 48, 3, 10, 5, 0, 48, 74, 1, 0, 0, 0, 49, 50, 3, 2,
		1, 0, 50, 51, 7, 2, 0, 0, 51, 52, 3, 4, 2, 0, 52, 74, 1, 0, 0, 0, 53, 54,
		3, 2, 1, 0, 54, 55, 7, 3, 0, 0, 55, 56, 3, 10, 5, 0, 56, 74, 1, 0, 0, 0,
		57, 58, 3, 2, 1, 0, 58, 59, 7, 3, 0, 0, 59, 60, 3, 4, 2, 0, 60, 74, 1,
		0, 0, 0, 61, 62, 3, 2, 1, 0, 62, 63, 5, 25, 0, 0, 63, 64, 3, 16, 8, 0,
		64, 74, 1, 0, 0, 0, 65, 66, 3, 2, 1, 0, 66, 67, 5, 1, 0, 0, 67, 68, 3,
		22, 11, 0, 68, 74, 1, 0, 0, 0, 69, 70, 3, 2, 1, 0, 70, 71, 7, 4, 0, 0,
		71, 72, 5, 28, 0, 0, 72, 74, 1, 0, 0, 0, 73, 28, 1, 0, 0, 0, 73, 33, 1,
		0, 0, 0, 73, 37, 1, 0, 0, 0, 73, 41, 1, 0, 0, 0, 73, 45, 1, 0, 0, 0, 73,
		49, 1, 0, 0, 0, 73, 53, 1, 0, 0, 0, 73, 57, 1, 0, 0, 0, 73, 61, 1, 0, 0,
		0, 73, 65, 1, 0, 0, 0, 73, 69, 1, 0, 0, 0, 74, 83, 1, 0, 0, 0, 75, 76,
		10, 13, 0, 0, 76, 77, 5, 18, 0, 0, 77, 82, 3, 0, 0, 14, 78, 79, 10, 12,
		0, 0, 79, 80, 5, 19, 0, 0, 80, 82, 3, 0, 0, 13, 81, 75, 1, 0, 0, 0, 81,
		78, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0,
		0, 84, 1, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 87, 5, 33, 0, 0, 87, 3, 1,
		0, 0, 0, 88, 89, 5, 14, 0, 0, 89, 90, 3, 6, 3, 0, 90, 91, 5, 15, 0, 0,
		91, 5, 1, 0, 0, 0, 92, 103, 3, 8, 4, 0, 93, 97, 5, 2, 0, 0, 94, 96, 5,
		31, 0, 0, 95, 94, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97,
		98, 1, 0, 0, 0, 98, 100, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 100, 102, 3, 8,
		4, 0, 101, 93, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0,
		103, 104, 1, 0, 0, 0, 104, 7, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 107,
		5, 28, 0, 0, 107, 9, 1, 0, 0, 0, 108, 109, 5, 14, 0, 0, 109, 110, 3, 12,
		6, 0, 110, 111, 5, 15, 0, 0, 111, 11, 1, 0, 0, 0, 112, 123, 3, 14, 7, 0,
		113, 117, 5, 2, 0, 0, 114, 116, 5, 31, 0, 0, 115, 114, 1, 0, 0, 0, 116,
		119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 120,
		1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 122, 3, 14, 7, 0, 121, 113, 1, 0,
		0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0,
		124, 13, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 127, 5, 26, 0, 0, 127,
		15, 1, 0, 0, 0, 128, 129, 5, 14, 0, 0, 129, 140, 3, 18, 9, 0, 130, 134,
		5, 2, 0, 0, 131, 133, 5, 31, 0, 0, 132, 131, 1, 0, 0, 0, 133, 136, 1, 0,
		0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 137, 1, 0, 0, 0,
		136, 134, 1, 0, 0, 0, 137, 139, 3, 18, 9, 0, 138, 130, 1, 0, 0, 0, 139,
		142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 143,
		1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 144, 5, 15, 0, 0, 144, 17, 1, 0,
		0, 0, 145, 146, 5, 14, 0, 0, 146, 157, 3, 20, 10, 0, 147, 151, 5, 2, 0,
		0, 148, 150, 5, 31, 0, 0, 149, 148, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151,
		149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 154, 1, 0, 0, 0, 153, 151,
		1, 0, 0, 0, 154, 156, 3, 20, 10, 0, 155, 147, 1, 0, 0, 0, 156, 159, 1,
		0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 160, 1, 0, 0,
		0, 159, 157, 1, 0, 0, 0, 160, 161, 5, 15, 0, 0, 161, 19, 1, 0, 0, 0, 162,
		163, 7, 5, 0, 0, 163, 21, 1, 0, 0, 0, 164, 165, 5, 14, 0, 0, 165, 170,
		3, 24, 12, 0, 166, 167, 5, 2, 0, 0, 167, 169, 3, 24, 12, 0, 168, 166, 1,
		0, 0, 0, 169, 172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0,
		0, 171, 173, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 174, 5, 15, 0, 0, 174,
		23, 1, 0, 0, 0, 175, 176, 5, 14, 0, 0, 176, 186, 3, 26, 13, 0, 177, 179,
		5, 31, 0, 0, 178, 177, 1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 178, 1, 0,
		0, 0, 180, 181, 1, 0, 0, 0, 181, 183, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0,
		183, 185, 3, 26, 13, 0, 184, 180, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186,
		184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0, 188, 186,
		1, 0, 0, 0, 189, 190, 5, 15, 0, 0, 190, 25, 1, 0, 0, 0, 191, 192, 5, 28,
		0, 0, 192, 27, 1, 0, 0, 0, 14, 73, 81, 83, 97, 103, 117, 123, 134, 140,
		151, 157, 170, 180, 186,
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

// MyRuleParserInit initializes any static state used to implement MyRuleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMyRuleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MyRuleParserInit() {
	staticData := &myruleParserStaticData
	staticData.once.Do(myruleParserInit)
}

// NewMyRuleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMyRuleParser(input antlr.TokenStream) *MyRuleParser {
	MyRuleParserInit()
	this := new(MyRuleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &myruleParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MyRule.g4"

	return this
}

// MyRuleParser tokens.
const (
	MyRuleParserEOF      = antlr.TokenEOF
	MyRuleParserT__0     = 1
	MyRuleParserT__1     = 2
	MyRuleParserEQ       = 3
	MyRuleParserLT       = 4
	MyRuleParserGT       = 5
	MyRuleParserLTEQ     = 6
	MyRuleParserGTEQ     = 7
	MyRuleParserNEQ      = 8
	MyRuleParserIN       = 9
	MyRuleParserNIN      = 10
	MyRuleParserHALL     = 11
	MyRuleParserHNONE    = 12
	MyRuleParserHANY     = 13
	MyRuleParserOBRACKET = 14
	MyRuleParserCBRACKET = 15
	MyRuleParserOPARAM   = 16
	MyRuleParserCPARAM   = 17
	MyRuleParserAND      = 18
	MyRuleParserOR       = 19
	MyRuleParserVNT      = 20
	MyRuleParserVNTE     = 21
	MyRuleParserVOT      = 22
	MyRuleParserVOTE     = 23
	MyRuleParserVE       = 24
	MyRuleParserWAO      = 25
	MyRuleParserINT      = 26
	MyRuleParserFLOAT    = 27
	MyRuleParserSTRING   = 28
	MyRuleParserTRUE     = 29
	MyRuleParserFALSE    = 30
	MyRuleParserSPACE    = 31
	MyRuleParserWS       = 32
	MyRuleParserATTRNAME = 33
)

// MyRuleParser rules.
const (
	MyRuleParserRULE_query             = 0
	MyRuleParserRULE_leftexpr          = 1
	MyRuleParserRULE_listStrings       = 2
	MyRuleParserRULE_stringElements    = 3
	MyRuleParserRULE_stringElement     = 4
	MyRuleParserRULE_listInts          = 5
	MyRuleParserRULE_intElements       = 6
	MyRuleParserRULE_intElement        = 7
	MyRuleParserRULE_listOfLists       = 8
	MyRuleParserRULE_list              = 9
	MyRuleParserRULE_floatElement      = 10
	MyRuleParserRULE_listOfTimeLists   = 11
	MyRuleParserRULE_timeList          = 12
	MyRuleParserRULE_timeStringElement = 13
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyFrom(ctx *QueryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type HasIntsExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewHasIntsExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HasIntsExpContext {
	var p = new(HasIntsExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *HasIntsExpContext) GetOp() antlr.Token { return s.op }

func (s *HasIntsExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *HasIntsExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HasIntsExpContext) Leftexpr() ILeftexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftexprContext)
}

func (s *HasIntsExpContext) ListInts() IListIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListIntsContext)
}

func (s *HasIntsExpContext) HALL() antlr.TerminalNode {
	return s.GetToken(MyRuleParserHALL, 0)
}

func (s *HasIntsExpContext) HANY() antlr.TerminalNode {
	return s.GetToken(MyRuleParserHANY, 0)
}

func (s *HasIntsExpContext) HNONE() antlr.TerminalNode {
	return s.GetToken(MyRuleParserHNONE, 0)
}

func (s *HasIntsExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterHasIntsExp(s)
	}
}

func (s *HasIntsExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitHasIntsExp(s)
	}
}

func (s *HasIntsExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitHasIntsExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareIntExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewCompareIntExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareIntExpContext {
	var p = new(CompareIntExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *CompareIntExpContext) GetOp() antlr.Token { return s.op }

func (s *CompareIntExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareIntExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareIntExpContext) Leftexpr() ILeftexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftexprContext)
}

func (s *CompareIntExpContext) INT() antlr.TerminalNode {
	return s.GetToken(MyRuleParserINT, 0)
}

func (s *CompareIntExpContext) GT() antlr.TerminalNode {
	return s.GetToken(MyRuleParserGT, 0)
}

func (s *CompareIntExpContext) LT() antlr.TerminalNode {
	return s.GetToken(MyRuleParserLT, 0)
}

func (s *CompareIntExpContext) GTEQ() antlr.TerminalNode {
	return s.GetToken(MyRuleParserGTEQ, 0)
}

func (s *CompareIntExpContext) LTEQ() antlr.TerminalNode {
	return s.GetToken(MyRuleParserLTEQ, 0)
}

func (s *CompareIntExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterCompareIntExp(s)
	}
}

func (s *CompareIntExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitCompareIntExp(s)
	}
}

func (s *CompareIntExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitCompareIntExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type WithinExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewWithinExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WithinExpContext {
	var p = new(WithinExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *WithinExpContext) GetOp() antlr.Token { return s.op }

func (s *WithinExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *WithinExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithinExpContext) Leftexpr() ILeftexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftexprContext)
}

func (s *WithinExpContext) ListOfLists() IListOfListsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListOfListsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListOfListsContext)
}

func (s *WithinExpContext) WAO() antlr.TerminalNode {
	return s.GetToken(MyRuleParserWAO, 0)
}

func (s *WithinExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterWithinExp(s)
	}
}

func (s *WithinExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitWithinExp(s)
	}
}

func (s *WithinExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitWithinExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	*QueryContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllQuery() []IQueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQueryContext); ok {
			len++
		}
	}

	tst := make([]IQueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQueryContext); ok {
			tst[i] = t.(IQueryContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Query(i int) IQueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
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

	return t.(IQueryContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(MyRuleParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualStringExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewEqualStringExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualStringExpContext {
	var p = new(EqualStringExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *EqualStringExpContext) GetOp() antlr.Token { return s.op }

func (s *EqualStringExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualStringExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualStringExpContext) Leftexpr() ILeftexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftexprContext)
}

func (s *EqualStringExpContext) STRING() antlr.TerminalNode {
	return s.GetToken(MyRuleParserSTRING, 0)
}

func (s *EqualStringExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(MyRuleParserEQ, 0)
}

func (s *EqualStringExpContext) NEQ() antlr.TerminalNode {
	return s.GetToken(MyRuleParserNEQ, 0)
}

func (s *EqualStringExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterEqualStringExp(s)
	}
}

func (s *EqualStringExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitEqualStringExp(s)
	}
}

func (s *EqualStringExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitEqualStringExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type VersionExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewVersionExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VersionExpContext {
	var p = new(VersionExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *VersionExpContext) GetOp() antlr.Token { return s.op }

func (s *VersionExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *VersionExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionExpContext) Leftexpr() ILeftexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftexprContext)
}

func (s *VersionExpContext) STRING() antlr.TerminalNode {
	return s.GetToken(MyRuleParserSTRING, 0)
}

func (s *VersionExpContext) VNT() antlr.TerminalNode {
	return s.GetToken(MyRuleParserVNT, 0)
}

func (s *VersionExpContext) VNTE() antlr.TerminalNode {
	return s.GetToken(MyRuleParserVNTE, 0)
}

func (s *VersionExpContext) VOT() antlr.TerminalNode {
	return s.GetToken(MyRuleParserVOT, 0)
}

func (s *VersionExpContext) VOTE() antlr.TerminalNode {
	return s.GetToken(MyRuleParserVOTE, 0)
}

func (s *VersionExpContext) VE() antlr.TerminalNode {
	return s.GetToken(MyRuleParserVE, 0)
}

func (s *VersionExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterVersionExp(s)
	}
}

func (s *VersionExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitVersionExp(s)
	}
}

func (s *VersionExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitVersionExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type TimeRangeExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewTimeRangeExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimeRangeExpContext {
	var p = new(TimeRangeExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *TimeRangeExpContext) GetOp() antlr.Token { return s.op }

func (s *TimeRangeExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *TimeRangeExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeRangeExpContext) Leftexpr() ILeftexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftexprContext)
}

func (s *TimeRangeExpContext) ListOfTimeLists() IListOfTimeListsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListOfTimeListsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListOfTimeListsContext)
}

func (s *TimeRangeExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterTimeRangeExp(s)
	}
}

func (s *TimeRangeExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitTimeRangeExp(s)
	}
}

func (s *TimeRangeExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitTimeRangeExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmbbedExprContext struct {
	*QueryContext
}

func NewEmbbedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmbbedExprContext {
	var p = new(EmbbedExprContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *EmbbedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmbbedExprContext) OPARAM() antlr.TerminalNode {
	return s.GetToken(MyRuleParserOPARAM, 0)
}

func (s *EmbbedExprContext) Query() IQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *EmbbedExprContext) CPARAM() antlr.TerminalNode {
	return s.GetToken(MyRuleParserCPARAM, 0)
}

func (s *EmbbedExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterEmbbedExpr(s)
	}
}

func (s *EmbbedExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitEmbbedExpr(s)
	}
}

func (s *EmbbedExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitEmbbedExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type InIntsExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewInIntsExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InIntsExpContext {
	var p = new(InIntsExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *InIntsExpContext) GetOp() antlr.Token { return s.op }

func (s *InIntsExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *InIntsExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InIntsExpContext) Leftexpr() ILeftexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftexprContext)
}

func (s *InIntsExpContext) ListInts() IListIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListIntsContext)
}

func (s *InIntsExpContext) IN() antlr.TerminalNode {
	return s.GetToken(MyRuleParserIN, 0)
}

func (s *InIntsExpContext) NIN() antlr.TerminalNode {
	return s.GetToken(MyRuleParserNIN, 0)
}

func (s *InIntsExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterInIntsExp(s)
	}
}

func (s *InIntsExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitInIntsExp(s)
	}
}

func (s *InIntsExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitInIntsExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualIntExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewEqualIntExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualIntExpContext {
	var p = new(EqualIntExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *EqualIntExpContext) GetOp() antlr.Token { return s.op }

func (s *EqualIntExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualIntExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualIntExpContext) Leftexpr() ILeftexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftexprContext)
}

func (s *EqualIntExpContext) INT() antlr.TerminalNode {
	return s.GetToken(MyRuleParserINT, 0)
}

func (s *EqualIntExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(MyRuleParserEQ, 0)
}

func (s *EqualIntExpContext) NEQ() antlr.TerminalNode {
	return s.GetToken(MyRuleParserNEQ, 0)
}

func (s *EqualIntExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterEqualIntExp(s)
	}
}

func (s *EqualIntExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitEqualIntExp(s)
	}
}

func (s *EqualIntExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitEqualIntExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type HasStringsExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewHasStringsExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HasStringsExpContext {
	var p = new(HasStringsExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *HasStringsExpContext) GetOp() antlr.Token { return s.op }

func (s *HasStringsExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *HasStringsExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HasStringsExpContext) Leftexpr() ILeftexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftexprContext)
}

func (s *HasStringsExpContext) ListStrings() IListStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStringsContext)
}

func (s *HasStringsExpContext) HALL() antlr.TerminalNode {
	return s.GetToken(MyRuleParserHALL, 0)
}

func (s *HasStringsExpContext) HANY() antlr.TerminalNode {
	return s.GetToken(MyRuleParserHANY, 0)
}

func (s *HasStringsExpContext) HNONE() antlr.TerminalNode {
	return s.GetToken(MyRuleParserHNONE, 0)
}

func (s *HasStringsExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterHasStringsExp(s)
	}
}

func (s *HasStringsExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitHasStringsExp(s)
	}
}

func (s *HasStringsExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitHasStringsExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type InStringsExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewInStringsExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InStringsExpContext {
	var p = new(InStringsExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *InStringsExpContext) GetOp() antlr.Token { return s.op }

func (s *InStringsExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *InStringsExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InStringsExpContext) Leftexpr() ILeftexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftexprContext)
}

func (s *InStringsExpContext) ListStrings() IListStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStringsContext)
}

func (s *InStringsExpContext) IN() antlr.TerminalNode {
	return s.GetToken(MyRuleParserIN, 0)
}

func (s *InStringsExpContext) NIN() antlr.TerminalNode {
	return s.GetToken(MyRuleParserNIN, 0)
}

func (s *InStringsExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterInStringsExp(s)
	}
}

func (s *InStringsExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitInStringsExp(s)
	}
}

func (s *InStringsExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitInStringsExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	*QueryContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllQuery() []IQueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQueryContext); ok {
			len++
		}
	}

	tst := make([]IQueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQueryContext); ok {
			tst[i] = t.(IQueryContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Query(i int) IQueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
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

	return t.(IQueryContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(MyRuleParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) Query() (localctx IQueryContext) {
	return p.query(0)
}

func (p *MyRuleParser) query(_p int) (localctx IQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, MyRuleParserRULE_query, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEmbbedExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(29)
			p.Match(MyRuleParserOPARAM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.query(0)
		}
		{
			p.SetState(31)
			p.Match(MyRuleParserCPARAM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewEqualIntExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(33)
			p.Leftexpr()
		}
		{
			p.SetState(34)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*EqualIntExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MyRuleParserEQ || _la == MyRuleParserNEQ) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*EqualIntExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(35)
			p.Match(MyRuleParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewEqualStringExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(37)
			p.Leftexpr()
		}
		{
			p.SetState(38)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*EqualStringExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MyRuleParserEQ || _la == MyRuleParserNEQ) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*EqualStringExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(39)
			p.Match(MyRuleParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewCompareIntExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(41)
			p.Leftexpr()
		}
		{
			p.SetState(42)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareIntExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareIntExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(43)
			p.Match(MyRuleParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewInIntsExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.Leftexpr()
		}
		{
			p.SetState(46)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*InIntsExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MyRuleParserIN || _la == MyRuleParserNIN) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*InIntsExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(47)
			p.ListInts()
		}

	case 6:
		localctx = NewInStringsExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Leftexpr()
		}
		{
			p.SetState(50)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*InStringsExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MyRuleParserIN || _la == MyRuleParserNIN) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*InStringsExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(51)
			p.ListStrings()
		}

	case 7:
		localctx = NewHasIntsExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(53)
			p.Leftexpr()
		}
		{
			p.SetState(54)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*HasIntsExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14336) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*HasIntsExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(55)
			p.ListInts()
		}

	case 8:
		localctx = NewHasStringsExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(57)
			p.Leftexpr()
		}
		{
			p.SetState(58)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*HasStringsExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14336) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*HasStringsExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(59)
			p.ListStrings()
		}

	case 9:
		localctx = NewWithinExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Leftexpr()
		}
		{
			p.SetState(62)

			var _m = p.Match(MyRuleParserWAO)

			localctx.(*WithinExpContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.ListOfLists()
		}

	case 10:
		localctx = NewTimeRangeExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)
			p.Leftexpr()
		}
		{
			p.SetState(66)

			var _m = p.Match(MyRuleParserT__0)

			localctx.(*TimeRangeExpContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.ListOfTimeLists()
		}

	case 11:
		localctx = NewVersionExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Leftexpr()
		}
		{
			p.SetState(70)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VersionExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32505856) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VersionExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(71)
			p.Match(MyRuleParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(83)
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
			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndExprContext(p, NewQueryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyRuleParserRULE_query)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(76)
					p.Match(MyRuleParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(77)
					p.query(14)
				}

			case 2:
				localctx = NewOrExprContext(p, NewQueryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyRuleParserRULE_query)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(79)
					p.Match(MyRuleParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.query(13)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(85)
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

// ILeftexprContext is an interface to support dynamic dispatch.
type ILeftexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATTRNAME() antlr.TerminalNode

	// IsLeftexprContext differentiates from other interfaces.
	IsLeftexprContext()
}

type LeftexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeftexprContext() *LeftexprContext {
	var p = new(LeftexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_leftexpr
	return p
}

func (*LeftexprContext) IsLeftexprContext() {}

func NewLeftexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftexprContext {
	var p = new(LeftexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_leftexpr

	return p
}

func (s *LeftexprContext) GetParser() antlr.Parser { return s.parser }

func (s *LeftexprContext) ATTRNAME() antlr.TerminalNode {
	return s.GetToken(MyRuleParserATTRNAME, 0)
}

func (s *LeftexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeftexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterLeftexpr(s)
	}
}

func (s *LeftexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitLeftexpr(s)
	}
}

func (s *LeftexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitLeftexpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) Leftexpr() (localctx ILeftexprContext) {
	localctx = NewLeftexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MyRuleParserRULE_leftexpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(MyRuleParserATTRNAME)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListStringsContext is an interface to support dynamic dispatch.
type IListStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBRACKET() antlr.TerminalNode
	StringElements() IStringElementsContext
	CBRACKET() antlr.TerminalNode

	// IsListStringsContext differentiates from other interfaces.
	IsListStringsContext()
}

type ListStringsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListStringsContext() *ListStringsContext {
	var p = new(ListStringsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_listStrings
	return p
}

func (*ListStringsContext) IsListStringsContext() {}

func NewListStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStringsContext {
	var p = new(ListStringsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_listStrings

	return p
}

func (s *ListStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStringsContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserOBRACKET, 0)
}

func (s *ListStringsContext) StringElements() IStringElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringElementsContext)
}

func (s *ListStringsContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserCBRACKET, 0)
}

func (s *ListStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStringsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterListStrings(s)
	}
}

func (s *ListStringsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitListStrings(s)
	}
}

func (s *ListStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitListStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) ListStrings() (localctx IListStringsContext) {
	localctx = NewListStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MyRuleParserRULE_listStrings)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(MyRuleParserOBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.StringElements()
	}
	{
		p.SetState(90)
		p.Match(MyRuleParserCBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringElementsContext is an interface to support dynamic dispatch.
type IStringElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStringElement() []IStringElementContext
	StringElement(i int) IStringElementContext
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode

	// IsStringElementsContext differentiates from other interfaces.
	IsStringElementsContext()
}

type StringElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringElementsContext() *StringElementsContext {
	var p = new(StringElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_stringElements
	return p
}

func (*StringElementsContext) IsStringElementsContext() {}

func NewStringElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringElementsContext {
	var p = new(StringElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_stringElements

	return p
}

func (s *StringElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StringElementsContext) AllStringElement() []IStringElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringElementContext); ok {
			len++
		}
	}

	tst := make([]IStringElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringElementContext); ok {
			tst[i] = t.(IStringElementContext)
			i++
		}
	}

	return tst
}

func (s *StringElementsContext) StringElement(i int) IStringElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringElementContext); ok {
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

	return t.(IStringElementContext)
}

func (s *StringElementsContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(MyRuleParserSPACE)
}

func (s *StringElementsContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(MyRuleParserSPACE, i)
}

func (s *StringElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterStringElements(s)
	}
}

func (s *StringElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitStringElements(s)
	}
}

func (s *StringElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitStringElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) StringElements() (localctx IStringElementsContext) {
	localctx = NewStringElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MyRuleParserRULE_stringElements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.StringElement()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MyRuleParserT__1 {
		{
			p.SetState(93)
			p.Match(MyRuleParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MyRuleParserSPACE {
			{
				p.SetState(94)
				p.Match(MyRuleParserSPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(99)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(100)
			p.StringElement()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IStringElementContext is an interface to support dynamic dispatch.
type IStringElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsStringElementContext differentiates from other interfaces.
	IsStringElementContext()
}

type StringElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringElementContext() *StringElementContext {
	var p = new(StringElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_stringElement
	return p
}

func (*StringElementContext) IsStringElementContext() {}

func NewStringElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringElementContext {
	var p = new(StringElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_stringElement

	return p
}

func (s *StringElementContext) GetParser() antlr.Parser { return s.parser }

func (s *StringElementContext) STRING() antlr.TerminalNode {
	return s.GetToken(MyRuleParserSTRING, 0)
}

func (s *StringElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterStringElement(s)
	}
}

func (s *StringElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitStringElement(s)
	}
}

func (s *StringElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitStringElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) StringElement() (localctx IStringElementContext) {
	localctx = NewStringElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MyRuleParserRULE_stringElement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(MyRuleParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListIntsContext is an interface to support dynamic dispatch.
type IListIntsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBRACKET() antlr.TerminalNode
	IntElements() IIntElementsContext
	CBRACKET() antlr.TerminalNode

	// IsListIntsContext differentiates from other interfaces.
	IsListIntsContext()
}

type ListIntsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListIntsContext() *ListIntsContext {
	var p = new(ListIntsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_listInts
	return p
}

func (*ListIntsContext) IsListIntsContext() {}

func NewListIntsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListIntsContext {
	var p = new(ListIntsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_listInts

	return p
}

func (s *ListIntsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListIntsContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserOBRACKET, 0)
}

func (s *ListIntsContext) IntElements() IIntElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntElementsContext)
}

func (s *ListIntsContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserCBRACKET, 0)
}

func (s *ListIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListIntsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListIntsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterListInts(s)
	}
}

func (s *ListIntsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitListInts(s)
	}
}

func (s *ListIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitListInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) ListInts() (localctx IListIntsContext) {
	localctx = NewListIntsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MyRuleParserRULE_listInts)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(MyRuleParserOBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.IntElements()
	}
	{
		p.SetState(110)
		p.Match(MyRuleParserCBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntElementsContext is an interface to support dynamic dispatch.
type IIntElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIntElement() []IIntElementContext
	IntElement(i int) IIntElementContext
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode

	// IsIntElementsContext differentiates from other interfaces.
	IsIntElementsContext()
}

type IntElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntElementsContext() *IntElementsContext {
	var p = new(IntElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_intElements
	return p
}

func (*IntElementsContext) IsIntElementsContext() {}

func NewIntElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntElementsContext {
	var p = new(IntElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_intElements

	return p
}

func (s *IntElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *IntElementsContext) AllIntElement() []IIntElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIntElementContext); ok {
			len++
		}
	}

	tst := make([]IIntElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIntElementContext); ok {
			tst[i] = t.(IIntElementContext)
			i++
		}
	}

	return tst
}

func (s *IntElementsContext) IntElement(i int) IIntElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntElementContext); ok {
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

	return t.(IIntElementContext)
}

func (s *IntElementsContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(MyRuleParserSPACE)
}

func (s *IntElementsContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(MyRuleParserSPACE, i)
}

func (s *IntElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterIntElements(s)
	}
}

func (s *IntElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitIntElements(s)
	}
}

func (s *IntElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitIntElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) IntElements() (localctx IIntElementsContext) {
	localctx = NewIntElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MyRuleParserRULE_intElements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.IntElement()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MyRuleParserT__1 {
		{
			p.SetState(113)
			p.Match(MyRuleParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MyRuleParserSPACE {
			{
				p.SetState(114)
				p.Match(MyRuleParserSPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(120)
			p.IntElement()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IIntElementContext is an interface to support dynamic dispatch.
type IIntElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsIntElementContext differentiates from other interfaces.
	IsIntElementContext()
}

type IntElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntElementContext() *IntElementContext {
	var p = new(IntElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_intElement
	return p
}

func (*IntElementContext) IsIntElementContext() {}

func NewIntElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntElementContext {
	var p = new(IntElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_intElement

	return p
}

func (s *IntElementContext) GetParser() antlr.Parser { return s.parser }

func (s *IntElementContext) INT() antlr.TerminalNode {
	return s.GetToken(MyRuleParserINT, 0)
}

func (s *IntElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterIntElement(s)
	}
}

func (s *IntElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitIntElement(s)
	}
}

func (s *IntElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitIntElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) IntElement() (localctx IIntElementContext) {
	localctx = NewIntElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MyRuleParserRULE_intElement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(MyRuleParserINT)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListOfListsContext is an interface to support dynamic dispatch.
type IListOfListsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBRACKET() antlr.TerminalNode
	AllList() []IListContext
	List(i int) IListContext
	CBRACKET() antlr.TerminalNode
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode

	// IsListOfListsContext differentiates from other interfaces.
	IsListOfListsContext()
}

type ListOfListsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListOfListsContext() *ListOfListsContext {
	var p = new(ListOfListsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_listOfLists
	return p
}

func (*ListOfListsContext) IsListOfListsContext() {}

func NewListOfListsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListOfListsContext {
	var p = new(ListOfListsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_listOfLists

	return p
}

func (s *ListOfListsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListOfListsContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserOBRACKET, 0)
}

func (s *ListOfListsContext) AllList() []IListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListContext); ok {
			len++
		}
	}

	tst := make([]IListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListContext); ok {
			tst[i] = t.(IListContext)
			i++
		}
	}

	return tst
}

func (s *ListOfListsContext) List(i int) IListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
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

	return t.(IListContext)
}

func (s *ListOfListsContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserCBRACKET, 0)
}

func (s *ListOfListsContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(MyRuleParserSPACE)
}

func (s *ListOfListsContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(MyRuleParserSPACE, i)
}

func (s *ListOfListsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfListsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListOfListsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterListOfLists(s)
	}
}

func (s *ListOfListsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitListOfLists(s)
	}
}

func (s *ListOfListsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitListOfLists(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) ListOfLists() (localctx IListOfListsContext) {
	localctx = NewListOfListsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MyRuleParserRULE_listOfLists)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(MyRuleParserOBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.List()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MyRuleParserT__1 {
		{
			p.SetState(130)
			p.Match(MyRuleParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MyRuleParserSPACE {
			{
				p.SetState(131)
				p.Match(MyRuleParserSPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(137)
			p.List()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
		p.Match(MyRuleParserCBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBRACKET() antlr.TerminalNode
	AllFloatElement() []IFloatElementContext
	FloatElement(i int) IFloatElementContext
	CBRACKET() antlr.TerminalNode
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserOBRACKET, 0)
}

func (s *ListContext) AllFloatElement() []IFloatElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFloatElementContext); ok {
			len++
		}
	}

	tst := make([]IFloatElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFloatElementContext); ok {
			tst[i] = t.(IFloatElementContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) FloatElement(i int) IFloatElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatElementContext); ok {
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

	return t.(IFloatElementContext)
}

func (s *ListContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserCBRACKET, 0)
}

func (s *ListContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(MyRuleParserSPACE)
}

func (s *ListContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(MyRuleParserSPACE, i)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MyRuleParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(MyRuleParserOBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.FloatElement()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MyRuleParserT__1 {
		{
			p.SetState(147)
			p.Match(MyRuleParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MyRuleParserSPACE {
			{
				p.SetState(148)
				p.Match(MyRuleParserSPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(154)
			p.FloatElement()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(160)
		p.Match(MyRuleParserCBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloatElementContext is an interface to support dynamic dispatch.
type IFloatElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsFloatElementContext differentiates from other interfaces.
	IsFloatElementContext()
}

type FloatElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatElementContext() *FloatElementContext {
	var p = new(FloatElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_floatElement
	return p
}

func (*FloatElementContext) IsFloatElementContext() {}

func NewFloatElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatElementContext {
	var p = new(FloatElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_floatElement

	return p
}

func (s *FloatElementContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatElementContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(MyRuleParserFLOAT, 0)
}

func (s *FloatElementContext) INT() antlr.TerminalNode {
	return s.GetToken(MyRuleParserINT, 0)
}

func (s *FloatElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterFloatElement(s)
	}
}

func (s *FloatElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitFloatElement(s)
	}
}

func (s *FloatElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitFloatElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) FloatElement() (localctx IFloatElementContext) {
	localctx = NewFloatElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MyRuleParserRULE_floatElement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MyRuleParserINT || _la == MyRuleParserFLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListOfTimeListsContext is an interface to support dynamic dispatch.
type IListOfTimeListsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBRACKET() antlr.TerminalNode
	AllTimeList() []ITimeListContext
	TimeList(i int) ITimeListContext
	CBRACKET() antlr.TerminalNode

	// IsListOfTimeListsContext differentiates from other interfaces.
	IsListOfTimeListsContext()
}

type ListOfTimeListsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListOfTimeListsContext() *ListOfTimeListsContext {
	var p = new(ListOfTimeListsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_listOfTimeLists
	return p
}

func (*ListOfTimeListsContext) IsListOfTimeListsContext() {}

func NewListOfTimeListsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListOfTimeListsContext {
	var p = new(ListOfTimeListsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_listOfTimeLists

	return p
}

func (s *ListOfTimeListsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListOfTimeListsContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserOBRACKET, 0)
}

func (s *ListOfTimeListsContext) AllTimeList() []ITimeListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITimeListContext); ok {
			len++
		}
	}

	tst := make([]ITimeListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITimeListContext); ok {
			tst[i] = t.(ITimeListContext)
			i++
		}
	}

	return tst
}

func (s *ListOfTimeListsContext) TimeList(i int) ITimeListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeListContext); ok {
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

	return t.(ITimeListContext)
}

func (s *ListOfTimeListsContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserCBRACKET, 0)
}

func (s *ListOfTimeListsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfTimeListsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListOfTimeListsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterListOfTimeLists(s)
	}
}

func (s *ListOfTimeListsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitListOfTimeLists(s)
	}
}

func (s *ListOfTimeListsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitListOfTimeLists(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) ListOfTimeLists() (localctx IListOfTimeListsContext) {
	localctx = NewListOfTimeListsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MyRuleParserRULE_listOfTimeLists)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(MyRuleParserOBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.TimeList()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MyRuleParserT__1 {
		{
			p.SetState(166)
			p.Match(MyRuleParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.TimeList()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(173)
		p.Match(MyRuleParserCBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimeListContext is an interface to support dynamic dispatch.
type ITimeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBRACKET() antlr.TerminalNode
	AllTimeStringElement() []ITimeStringElementContext
	TimeStringElement(i int) ITimeStringElementContext
	CBRACKET() antlr.TerminalNode
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode

	// IsTimeListContext differentiates from other interfaces.
	IsTimeListContext()
}

type TimeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeListContext() *TimeListContext {
	var p = new(TimeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_timeList
	return p
}

func (*TimeListContext) IsTimeListContext() {}

func NewTimeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeListContext {
	var p = new(TimeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_timeList

	return p
}

func (s *TimeListContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeListContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserOBRACKET, 0)
}

func (s *TimeListContext) AllTimeStringElement() []ITimeStringElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITimeStringElementContext); ok {
			len++
		}
	}

	tst := make([]ITimeStringElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITimeStringElementContext); ok {
			tst[i] = t.(ITimeStringElementContext)
			i++
		}
	}

	return tst
}

func (s *TimeListContext) TimeStringElement(i int) ITimeStringElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeStringElementContext); ok {
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

	return t.(ITimeStringElementContext)
}

func (s *TimeListContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(MyRuleParserCBRACKET, 0)
}

func (s *TimeListContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(MyRuleParserSPACE)
}

func (s *TimeListContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(MyRuleParserSPACE, i)
}

func (s *TimeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterTimeList(s)
	}
}

func (s *TimeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitTimeList(s)
	}
}

func (s *TimeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitTimeList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) TimeList() (localctx ITimeListContext) {
	localctx = NewTimeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MyRuleParserRULE_timeList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(MyRuleParserOBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.TimeStringElement()
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MyRuleParserSTRING || _la == MyRuleParserSPACE {
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MyRuleParserSPACE {
			{
				p.SetState(177)
				p.Match(MyRuleParserSPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(183)
			p.TimeStringElement()
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(189)
		p.Match(MyRuleParserCBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimeStringElementContext is an interface to support dynamic dispatch.
type ITimeStringElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsTimeStringElementContext differentiates from other interfaces.
	IsTimeStringElementContext()
}

type TimeStringElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeStringElementContext() *TimeStringElementContext {
	var p = new(TimeStringElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyRuleParserRULE_timeStringElement
	return p
}

func (*TimeStringElementContext) IsTimeStringElementContext() {}

func NewTimeStringElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeStringElementContext {
	var p = new(TimeStringElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyRuleParserRULE_timeStringElement

	return p
}

func (s *TimeStringElementContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeStringElementContext) STRING() antlr.TerminalNode {
	return s.GetToken(MyRuleParserSTRING, 0)
}

func (s *TimeStringElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeStringElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeStringElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.EnterTimeStringElement(s)
	}
}

func (s *TimeStringElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyRuleListener); ok {
		listenerT.ExitTimeStringElement(s)
	}
}

func (s *TimeStringElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyRuleVisitor:
		return t.VisitTimeStringElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyRuleParser) TimeStringElement() (localctx ITimeStringElementContext) {
	localctx = NewTimeStringElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MyRuleParserRULE_timeStringElement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(MyRuleParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *MyRuleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *QueryContext = nil
		if localctx != nil {
			t = localctx.(*QueryContext)
		}
		return p.Query_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MyRuleParser) Query_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
