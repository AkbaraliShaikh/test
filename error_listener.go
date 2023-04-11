package rule

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	err      error
	hasError bool
}

func (e *ErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	e.err = fmt.Errorf("line %d:%d %s", line, column, msg)
	e.hasError = true
}
