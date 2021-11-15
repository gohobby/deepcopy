package mapcasegen

import (
	"deepcopy/includes/jenutils"

	. "github.com/dave/jennifer/jen"
)

func AllCases(types Statement, src, mainFunc Code) Statement {
	statement := make(Statement, 0)
	clone := Id("clone")

	statement = append(statement, Comment("Deep copy for types map[string]T"))
	statement = append(statement, jenutils.Cases(types, src, func(value, t Code) Statement {
		s := Statement{Line()}
		s = append(s, jenutils.ShallowCopy(t, value, clone)...)
		s = append(s, Line(), Return(clone))

		return s
	}, Map(String()))...)

	// map[string]interface{}
	statement = append(statement, jenutils.CaseType(Map(String()).Interface(), src, func(value, t Code) Statement {
		return deepCopy(t, value, clone, mainFunc)
	}))

	statement = append(statement, Comment("Deep copy for types map[string][]T"))
	statement = append(statement, jenutils.Cases(types, src, func(value, t Code) Statement {
		return deepCopy(t, value, clone, mainFunc)
	}, Map(String()).Index())...)

	// map[string][]interface{}
	statement = append(statement, jenutils.CaseType(Map(String()).Index().Interface(), src, func(value, t Code) Statement {
		return deepCopy(t, value, clone, mainFunc)
	}))

	return statement
}

func deepCopy(t, src, dst, recursiveFunc Code) Statement {
	s := Statement{Line()}
	s = append(s, jenutils.DeepCopy(t, src, dst, recursiveFunc)...)
	s = append(s, Line(), Return(dst))

	return s
}
