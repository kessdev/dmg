package transpiler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/dev-kess/dmg/parser"
)

type TranspilerListener struct {
	*parser.BaseDMGParserListener
}

func NewTranspilerListener() *TranspilerListener {
	return new(TranspilerListener)
}

func (listener *TranspilerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}
