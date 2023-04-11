package rule

import (
	"fmt"
	"test/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func Parse(rule string) (parser.IQueryContext, bool, error) {
	return parse(rule)
}

func parse(rule string) (parser.IQueryContext, bool, error) {
	input := antlr.NewInputStream(rule)
	lex := parser.NewMyRuleLexer(input)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)

	p := parser.NewMyRuleParser(tokens)
	el := &ErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	tree := p.Query()

	if el.hasError {
		return nil, false, fmt.Errorf("parsing failed, error: %v", el.err)
	}

	return tree, true, nil
}
