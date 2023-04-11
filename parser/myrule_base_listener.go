// Code generated from MyRule.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // MyRule

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseMyRuleListener is a complete listener for a parse tree produced by MyRuleParser.
type BaseMyRuleListener struct{}

var _ MyRuleListener = &BaseMyRuleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMyRuleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMyRuleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMyRuleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMyRuleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterHasIntsExp is called when production hasIntsExp is entered.
func (s *BaseMyRuleListener) EnterHasIntsExp(ctx *HasIntsExpContext) {}

// ExitHasIntsExp is called when production hasIntsExp is exited.
func (s *BaseMyRuleListener) ExitHasIntsExp(ctx *HasIntsExpContext) {}

// EnterCompareIntExp is called when production compareIntExp is entered.
func (s *BaseMyRuleListener) EnterCompareIntExp(ctx *CompareIntExpContext) {}

// ExitCompareIntExp is called when production compareIntExp is exited.
func (s *BaseMyRuleListener) ExitCompareIntExp(ctx *CompareIntExpContext) {}

// EnterWithinExp is called when production withinExp is entered.
func (s *BaseMyRuleListener) EnterWithinExp(ctx *WithinExpContext) {}

// ExitWithinExp is called when production withinExp is exited.
func (s *BaseMyRuleListener) ExitWithinExp(ctx *WithinExpContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseMyRuleListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseMyRuleListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterEqualStringExp is called when production equalStringExp is entered.
func (s *BaseMyRuleListener) EnterEqualStringExp(ctx *EqualStringExpContext) {}

// ExitEqualStringExp is called when production equalStringExp is exited.
func (s *BaseMyRuleListener) ExitEqualStringExp(ctx *EqualStringExpContext) {}

// EnterVersionExp is called when production versionExp is entered.
func (s *BaseMyRuleListener) EnterVersionExp(ctx *VersionExpContext) {}

// ExitVersionExp is called when production versionExp is exited.
func (s *BaseMyRuleListener) ExitVersionExp(ctx *VersionExpContext) {}

// EnterTimeRangeExp is called when production timeRangeExp is entered.
func (s *BaseMyRuleListener) EnterTimeRangeExp(ctx *TimeRangeExpContext) {}

// ExitTimeRangeExp is called when production timeRangeExp is exited.
func (s *BaseMyRuleListener) ExitTimeRangeExp(ctx *TimeRangeExpContext) {}

// EnterEmbbedExpr is called when production embbedExpr is entered.
func (s *BaseMyRuleListener) EnterEmbbedExpr(ctx *EmbbedExprContext) {}

// ExitEmbbedExpr is called when production embbedExpr is exited.
func (s *BaseMyRuleListener) ExitEmbbedExpr(ctx *EmbbedExprContext) {}

// EnterInIntsExp is called when production inIntsExp is entered.
func (s *BaseMyRuleListener) EnterInIntsExp(ctx *InIntsExpContext) {}

// ExitInIntsExp is called when production inIntsExp is exited.
func (s *BaseMyRuleListener) ExitInIntsExp(ctx *InIntsExpContext) {}

// EnterEqualIntExp is called when production equalIntExp is entered.
func (s *BaseMyRuleListener) EnterEqualIntExp(ctx *EqualIntExpContext) {}

// ExitEqualIntExp is called when production equalIntExp is exited.
func (s *BaseMyRuleListener) ExitEqualIntExp(ctx *EqualIntExpContext) {}

// EnterHasStringsExp is called when production hasStringsExp is entered.
func (s *BaseMyRuleListener) EnterHasStringsExp(ctx *HasStringsExpContext) {}

// ExitHasStringsExp is called when production hasStringsExp is exited.
func (s *BaseMyRuleListener) ExitHasStringsExp(ctx *HasStringsExpContext) {}

// EnterInStringsExp is called when production inStringsExp is entered.
func (s *BaseMyRuleListener) EnterInStringsExp(ctx *InStringsExpContext) {}

// ExitInStringsExp is called when production inStringsExp is exited.
func (s *BaseMyRuleListener) ExitInStringsExp(ctx *InStringsExpContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseMyRuleListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseMyRuleListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterLeftexpr is called when production leftexpr is entered.
func (s *BaseMyRuleListener) EnterLeftexpr(ctx *LeftexprContext) {}

// ExitLeftexpr is called when production leftexpr is exited.
func (s *BaseMyRuleListener) ExitLeftexpr(ctx *LeftexprContext) {}

// EnterListStrings is called when production listStrings is entered.
func (s *BaseMyRuleListener) EnterListStrings(ctx *ListStringsContext) {}

// ExitListStrings is called when production listStrings is exited.
func (s *BaseMyRuleListener) ExitListStrings(ctx *ListStringsContext) {}

// EnterStringElements is called when production stringElements is entered.
func (s *BaseMyRuleListener) EnterStringElements(ctx *StringElementsContext) {}

// ExitStringElements is called when production stringElements is exited.
func (s *BaseMyRuleListener) ExitStringElements(ctx *StringElementsContext) {}

// EnterStringElement is called when production stringElement is entered.
func (s *BaseMyRuleListener) EnterStringElement(ctx *StringElementContext) {}

// ExitStringElement is called when production stringElement is exited.
func (s *BaseMyRuleListener) ExitStringElement(ctx *StringElementContext) {}

// EnterListInts is called when production listInts is entered.
func (s *BaseMyRuleListener) EnterListInts(ctx *ListIntsContext) {}

// ExitListInts is called when production listInts is exited.
func (s *BaseMyRuleListener) ExitListInts(ctx *ListIntsContext) {}

// EnterIntElements is called when production intElements is entered.
func (s *BaseMyRuleListener) EnterIntElements(ctx *IntElementsContext) {}

// ExitIntElements is called when production intElements is exited.
func (s *BaseMyRuleListener) ExitIntElements(ctx *IntElementsContext) {}

// EnterIntElement is called when production intElement is entered.
func (s *BaseMyRuleListener) EnterIntElement(ctx *IntElementContext) {}

// ExitIntElement is called when production intElement is exited.
func (s *BaseMyRuleListener) ExitIntElement(ctx *IntElementContext) {}

// EnterListOfLists is called when production listOfLists is entered.
func (s *BaseMyRuleListener) EnterListOfLists(ctx *ListOfListsContext) {}

// ExitListOfLists is called when production listOfLists is exited.
func (s *BaseMyRuleListener) ExitListOfLists(ctx *ListOfListsContext) {}

// EnterList is called when production list is entered.
func (s *BaseMyRuleListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseMyRuleListener) ExitList(ctx *ListContext) {}

// EnterFloatElement is called when production floatElement is entered.
func (s *BaseMyRuleListener) EnterFloatElement(ctx *FloatElementContext) {}

// ExitFloatElement is called when production floatElement is exited.
func (s *BaseMyRuleListener) ExitFloatElement(ctx *FloatElementContext) {}

// EnterListOfTimeLists is called when production listOfTimeLists is entered.
func (s *BaseMyRuleListener) EnterListOfTimeLists(ctx *ListOfTimeListsContext) {}

// ExitListOfTimeLists is called when production listOfTimeLists is exited.
func (s *BaseMyRuleListener) ExitListOfTimeLists(ctx *ListOfTimeListsContext) {}

// EnterTimeList is called when production timeList is entered.
func (s *BaseMyRuleListener) EnterTimeList(ctx *TimeListContext) {}

// ExitTimeList is called when production timeList is exited.
func (s *BaseMyRuleListener) ExitTimeList(ctx *TimeListContext) {}

// EnterTimeStringElement is called when production timeStringElement is entered.
func (s *BaseMyRuleListener) EnterTimeStringElement(ctx *TimeStringElementContext) {}

// ExitTimeStringElement is called when production timeStringElement is exited.
func (s *BaseMyRuleListener) ExitTimeStringElement(ctx *TimeStringElementContext) {}
