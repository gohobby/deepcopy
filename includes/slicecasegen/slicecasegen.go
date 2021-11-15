package slicecasegen

import (
	"github.com/gohobby/deepcopy/includes/jenutils"

	. "github.com/dave/jennifer/jen"
)

func AllCases(types Statement, src, mainFunc Code) Statement {
	statement := make(Statement, 0)
	clone := Id("clone")

	statement = append(statement, Comment("Deep copy for types []T"))
	statement = append(statement, jenutils.Cases(types, src, func(value, t Code) Statement {
		return Statement{
			Line(),
			Add(clone).Op(":=").Make(t, Len(value)),
			Copy(clone, value),
			Line(),
			Return(clone),
		}
	}, Index())...)

	// []interface{}
	statement = append(statement, jenutils.CaseType(Index().Interface(), src, func(value, t Code) Statement {
		return Statement{
			Line(),
			Add(clone).Op(":=").Make(t, Len(value)),
			Copy(clone, value),
			Line(),
			jenutils.Foreach(clone).Block(
				Add(clone).Index(Id("k")).Op("=").Add(mainFunc).Parens(Id("v")),
			),
			Line(),
			Return(clone),
		}
	}))

	return statement
}
