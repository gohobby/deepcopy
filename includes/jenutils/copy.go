package jenutils

import . "github.com/dave/jennifer/jen"

func ShallowCopy(t, src, dst Code) Statement {
	return Statement{
		Add(dst).Op(":=").Make(t, Len(src)),
		Line(),
		Foreach(src).Block(
			Add(dst).Index(Id("k")).Op("=").Id("v"),
		),
	}
}

func DeepCopy(t, src, dst, recursiveFunc Code) Statement {
	deepType := (*t.(*Statement))[1:]

	return Statement{
		Add(dst).Op(":=").Make(t, Len(src)),
		Line(),
		Foreach(src).Block(
			Add(dst).Index(Id("k")).Op("=").Add(recursiveFunc).Parens(Id("v")).Assert(&deepType),
		),
	}
}

func InterfaceCopy(t, src, dst, recursiveFunc Code) Statement {
	return Statement{
		Add(dst).Op(":=").Make(t, Len(src)),
		Line(),
		Foreach(src).Block(
			Add(dst).Index(Id("k")).Op("=").Add(recursiveFunc).Parens(Id("v")),
		),
	}
}
