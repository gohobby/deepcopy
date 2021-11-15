package jenutils

import . "github.com/dave/jennifer/jen"

// dereference the source variable if it is a pointer
// t is the type of the variable
// src is the variable to dereference
// return the dereferenced value and the dereferencing statement.
func dereference(t, src Code) (value *Statement, statement Statement) {
	value = Id("value")

	return value, Statement{
		Var().Add(value).Add(t),
		Line(),
		If(
			List(Id("val"), Id("ok")).
				Op(":=").
				Add(src).
				Assert(Add(Op("*")).Add(t)),
			Id("ok"),
		).Block(
			Add(value).Op("=").Add(Op("*")).Id("val"),
		).Else().Block(
			Add(value).Op("=").Add(src).Assert(t),
		),
	}
}
