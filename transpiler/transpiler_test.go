package transpiler_test

import (
	"testing"

	"github.com/dev-kess/dmg/transpiler"
)

func TestSelect(t *testing.T) {
	dataMapperDSL := "select p from Person p"
	transpilerDSL := transpiler.ProvideTranspiler()
	transpilerDSL.TranspileDSL(dataMapperDSL)
}
