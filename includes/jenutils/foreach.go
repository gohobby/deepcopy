package jenutils

import . "github.com/dave/jennifer/jen"

func Foreach(src Code, keyAndValue ...string) *Statement {
	key := "k"
	value := "v"

	switch len(keyAndValue) {
	case 1:
		key = keyAndValue[0]
	case 2:
		key = keyAndValue[0]
		value = keyAndValue[1]
	}

	return For(List(Id(key), Id(value)).Op(":=").Range().Add(src))
}
