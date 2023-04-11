// Code generated from MyRule.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // MyRule

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by MyRuleParser.
type MyRuleVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MyRuleParser#hasIntsExp.
	VisitHasIntsExp(ctx *HasIntsExpContext) interface{}

	// Visit a parse tree produced by MyRuleParser#compareIntExp.
	VisitCompareIntExp(ctx *CompareIntExpContext) interface{}

	// Visit a parse tree produced by MyRuleParser#withinExp.
	VisitWithinExp(ctx *WithinExpContext) interface{}

	// Visit a parse tree produced by MyRuleParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by MyRuleParser#equalStringExp.
	VisitEqualStringExp(ctx *EqualStringExpContext) interface{}

	// Visit a parse tree produced by MyRuleParser#versionExp.
	VisitVersionExp(ctx *VersionExpContext) interface{}

	// Visit a parse tree produced by MyRuleParser#timeRangeExp.
	VisitTimeRangeExp(ctx *TimeRangeExpContext) interface{}

	// Visit a parse tree produced by MyRuleParser#embbedExpr.
	VisitEmbbedExpr(ctx *EmbbedExprContext) interface{}

	// Visit a parse tree produced by MyRuleParser#inIntsExp.
	VisitInIntsExp(ctx *InIntsExpContext) interface{}

	// Visit a parse tree produced by MyRuleParser#equalIntExp.
	VisitEqualIntExp(ctx *EqualIntExpContext) interface{}

	// Visit a parse tree produced by MyRuleParser#hasStringsExp.
	VisitHasStringsExp(ctx *HasStringsExpContext) interface{}

	// Visit a parse tree produced by MyRuleParser#inStringsExp.
	VisitInStringsExp(ctx *InStringsExpContext) interface{}

	// Visit a parse tree produced by MyRuleParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by MyRuleParser#leftexpr.
	VisitLeftexpr(ctx *LeftexprContext) interface{}

	// Visit a parse tree produced by MyRuleParser#listStrings.
	VisitListStrings(ctx *ListStringsContext) interface{}

	// Visit a parse tree produced by MyRuleParser#stringElements.
	VisitStringElements(ctx *StringElementsContext) interface{}

	// Visit a parse tree produced by MyRuleParser#stringElement.
	VisitStringElement(ctx *StringElementContext) interface{}

	// Visit a parse tree produced by MyRuleParser#listInts.
	VisitListInts(ctx *ListIntsContext) interface{}

	// Visit a parse tree produced by MyRuleParser#intElements.
	VisitIntElements(ctx *IntElementsContext) interface{}

	// Visit a parse tree produced by MyRuleParser#intElement.
	VisitIntElement(ctx *IntElementContext) interface{}

	// Visit a parse tree produced by MyRuleParser#listOfLists.
	VisitListOfLists(ctx *ListOfListsContext) interface{}

	// Visit a parse tree produced by MyRuleParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by MyRuleParser#floatElement.
	VisitFloatElement(ctx *FloatElementContext) interface{}

	// Visit a parse tree produced by MyRuleParser#listOfTimeLists.
	VisitListOfTimeLists(ctx *ListOfTimeListsContext) interface{}

	// Visit a parse tree produced by MyRuleParser#timeList.
	VisitTimeList(ctx *TimeListContext) interface{}

	// Visit a parse tree produced by MyRuleParser#timeStringElement.
	VisitTimeStringElement(ctx *TimeStringElementContext) interface{}
}
