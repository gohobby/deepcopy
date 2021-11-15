package jenutils

import . "github.com/dave/jennifer/jen"

func SwitchType(src, dst Code, fn func(value Code) Statement) *Statement {
	return Switch(Add(dst).Op(":=").Add(src).Assert(Type())).Block(fn(dst)...)
}
