// Code generated from MyRule.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // MyRule

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseMyRuleVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMyRuleVisitor) VisitHasIntsExp(ctx *HasIntsExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitCompareIntExp(ctx *CompareIntExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitWithinExp(ctx *WithinExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitEqualStringExp(ctx *EqualStringExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitVersionExp(ctx *VersionExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitTimeRangeExp(ctx *TimeRangeExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitEmbbedExpr(ctx *EmbbedExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitInIntsExp(ctx *InIntsExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitEqualIntExp(ctx *EqualIntExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitHasStringsExp(ctx *HasStringsExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitInStringsExp(ctx *InStringsExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitLeftexpr(ctx *LeftexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitListStrings(ctx *ListStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitStringElements(ctx *StringElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitStringElement(ctx *StringElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitListInts(ctx *ListIntsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitIntElements(ctx *IntElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitIntElement(ctx *IntElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitListOfLists(ctx *ListOfListsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitFloatElement(ctx *FloatElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitListOfTimeLists(ctx *ListOfTimeListsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitTimeList(ctx *TimeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyRuleVisitor) VisitTimeStringElement(ctx *TimeStringElementContext) interface{} {
	return v.VisitChildren(ctx)
}
