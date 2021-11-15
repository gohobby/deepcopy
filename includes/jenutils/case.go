package jenutils

import . "github.com/dave/jennifer/jen"

// CaseType renders a case for the given type and its pointer.
// The closure renders a list of statements for this case.
//
// If the switch value is a pointer, it will be dereferenced.
//
// t is the type to check, its pointer will also be checked
// src is the variable to dereference
// fn is a closure taking as parameter the dereferenced value and returning a statement.
func CaseType(t, src Code, fn func(value, t Code) Statement) Code {
	value, statement := dereference(t, src)
	statement = append(statement, fn(value, t)...)

	return Case(t, Op("*").Add(t)).Block(statement...)
}

// Cases is a wrapper that renders a set of CaseType for each given type.
func Cases(types Statement, src Code, fn func(value, t Code) Statement, baseType ...*Statement) Statement {
	statement := make(Statement, 0)

	for _, t := range types {
		switch baseType {
		case nil:
			statement = append(statement, CaseType(t, src, fn))
		default:
			tt := append(*baseType[0], t)
			statement = append(statement, CaseType(&tt, src, fn))
		}
	}

	return statement
}
