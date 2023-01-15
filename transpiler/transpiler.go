package transpiler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/dev-kess/dmg/parser"
)

type Transpiler interface {
	TranspileDSL(dsl string) error
	TranspileFile(file string) error
}

type transpilerImpl struct {
}

func (transpiler transpilerImpl) TranspileDSL(dsl string) error {
	input := antlr.NewInputStream(dsl)
	transpileData(input)
	return nil
}

func (transpiler transpilerImpl) TranspileFile(file string) error {
	input, _ := antlr.NewFileStream(file)
	transpileData(input)
	return nil
}

func transpileData(input antlr.CharStream) {
	lexer := parser.NewDMGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	dataMapperParser := parser.NewDMGParser(stream)
	dataMapperParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	dataMapperParser.BuildParseTrees = true
	tree := dataMapperParser.Select_statement()
	listener := NewTranspilerListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

func ProvideTranspiler() Transpiler {
	return transpilerImpl{}
}
