// Code generated from MyRule.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // MyRule

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// MyRuleListener is a complete listener for a parse tree produced by MyRuleParser.
type MyRuleListener interface {
	antlr.ParseTreeListener

	// EnterHasIntsExp is called when entering the hasIntsExp production.
	EnterHasIntsExp(c *HasIntsExpContext)

	// EnterCompareIntExp is called when entering the compareIntExp production.
	EnterCompareIntExp(c *CompareIntExpContext)

	// EnterWithinExp is called when entering the withinExp production.
	EnterWithinExp(c *WithinExpContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterEqualStringExp is called when entering the equalStringExp production.
	EnterEqualStringExp(c *EqualStringExpContext)

	// EnterVersionExp is called when entering the versionExp production.
	EnterVersionExp(c *VersionExpContext)

	// EnterTimeRangeExp is called when entering the timeRangeExp production.
	EnterTimeRangeExp(c *TimeRangeExpContext)

	// EnterEmbbedExpr is called when entering the embbedExpr production.
	EnterEmbbedExpr(c *EmbbedExprContext)

	// EnterInIntsExp is called when entering the inIntsExp production.
	EnterInIntsExp(c *InIntsExpContext)

	// EnterEqualIntExp is called when entering the equalIntExp production.
	EnterEqualIntExp(c *EqualIntExpContext)

	// EnterHasStringsExp is called when entering the hasStringsExp production.
	EnterHasStringsExp(c *HasStringsExpContext)

	// EnterInStringsExp is called when entering the inStringsExp production.
	EnterInStringsExp(c *InStringsExpContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterLeftexpr is called when entering the leftexpr production.
	EnterLeftexpr(c *LeftexprContext)

	// EnterListStrings is called when entering the listStrings production.
	EnterListStrings(c *ListStringsContext)

	// EnterStringElements is called when entering the stringElements production.
	EnterStringElements(c *StringElementsContext)

	// EnterStringElement is called when entering the stringElement production.
	EnterStringElement(c *StringElementContext)

	// EnterListInts is called when entering the listInts production.
	EnterListInts(c *ListIntsContext)

	// EnterIntElements is called when entering the intElements production.
	EnterIntElements(c *IntElementsContext)

	// EnterIntElement is called when entering the intElement production.
	EnterIntElement(c *IntElementContext)

	// EnterListOfLists is called when entering the listOfLists production.
	EnterListOfLists(c *ListOfListsContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterFloatElement is called when entering the floatElement production.
	EnterFloatElement(c *FloatElementContext)

	// EnterListOfTimeLists is called when entering the listOfTimeLists production.
	EnterListOfTimeLists(c *ListOfTimeListsContext)

	// EnterTimeList is called when entering the timeList production.
	EnterTimeList(c *TimeListContext)

	// EnterTimeStringElement is called when entering the timeStringElement production.
	EnterTimeStringElement(c *TimeStringElementContext)

	// ExitHasIntsExp is called when exiting the hasIntsExp production.
	ExitHasIntsExp(c *HasIntsExpContext)

	// ExitCompareIntExp is called when exiting the compareIntExp production.
	ExitCompareIntExp(c *CompareIntExpContext)

	// ExitWithinExp is called when exiting the withinExp production.
	ExitWithinExp(c *WithinExpContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitEqualStringExp is called when exiting the equalStringExp production.
	ExitEqualStringExp(c *EqualStringExpContext)

	// ExitVersionExp is called when exiting the versionExp production.
	ExitVersionExp(c *VersionExpContext)

	// ExitTimeRangeExp is called when exiting the timeRangeExp production.
	ExitTimeRangeExp(c *TimeRangeExpContext)

	// ExitEmbbedExpr is called when exiting the embbedExpr production.
	ExitEmbbedExpr(c *EmbbedExprContext)

	// ExitInIntsExp is called when exiting the inIntsExp production.
	ExitInIntsExp(c *InIntsExpContext)

	// ExitEqualIntExp is called when exiting the equalIntExp production.
	ExitEqualIntExp(c *EqualIntExpContext)

	// ExitHasStringsExp is called when exiting the hasStringsExp production.
	ExitHasStringsExp(c *HasStringsExpContext)

	// ExitInStringsExp is called when exiting the inStringsExp production.
	ExitInStringsExp(c *InStringsExpContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitLeftexpr is called when exiting the leftexpr production.
	ExitLeftexpr(c *LeftexprContext)

	// ExitListStrings is called when exiting the listStrings production.
	ExitListStrings(c *ListStringsContext)

	// ExitStringElements is called when exiting the stringElements production.
	ExitStringElements(c *StringElementsContext)

	// ExitStringElement is called when exiting the stringElement production.
	ExitStringElement(c *StringElementContext)

	// ExitListInts is called when exiting the listInts production.
	ExitListInts(c *ListIntsContext)

	// ExitIntElements is called when exiting the intElements production.
	ExitIntElements(c *IntElementsContext)

	// ExitIntElement is called when exiting the intElement production.
	ExitIntElement(c *IntElementContext)

	// ExitListOfLists is called when exiting the listOfLists production.
	ExitListOfLists(c *ListOfListsContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitFloatElement is called when exiting the floatElement production.
	ExitFloatElement(c *FloatElementContext)

	// ExitListOfTimeLists is called when exiting the listOfTimeLists production.
	ExitListOfTimeLists(c *ListOfTimeListsContext)

	// ExitTimeList is called when exiting the timeList production.
	ExitTimeList(c *TimeListContext)

	// ExitTimeStringElement is called when exiting the timeStringElement production.
	ExitTimeStringElement(c *TimeStringElementContext)
}
