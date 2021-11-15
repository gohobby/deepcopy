package main

import (
	
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gohobby/deepcopy/includes/jenutils"
	"github.com/gohobby/deepcopy/includes/mapcasegen"
	"github.com/gohobby/deepcopy/includes/slicecasegen"

	. "github.com/dave/jennifer/jen"
)

// Types slice
var allType = Statement{
	Bool(),
	Int(),
	Int8(),
	Int16(),
	Int64(),
	Uint(),
	Uint8(),
	Uint16(),
	Uint64(),
	Uintptr(),
	Float32(),
	Float64(),
	Complex64(),
	Complex128(),
	String(),
	// Interface(),
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", path, "deepcopy.go"))
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	err = render(file)
	if err != nil {
		log.Panicln(err)
	}
}

func render(w io.Writer) error {
	f := NewFile("deepcopy")

	f.HeaderComment("This file is generated - do not edit.")
	f.Line()

	f.Type().Id("Copyable").Interface(
		Id("DeepCopy").Params().Interface(),
	)

	mainFunc := Id("DeepCopy")
	object := Id("object")

	f.Comment("DeepCopy will create a deep copy of the source object.")
	f.Comment("Maps and slices will be taken into account when copying.")
	f.Func().Add(mainFunc).Params(Add(object).Interface()).Interface().
		Block(
			jenutils.SwitchType(Add(object), Id("t"), func(value Code) Statement {
				statement := make(Statement, 0)

				statement = append(statement, jenutils.CaseType(Id("Copyable"), value, func(value, t Code) Statement {
					return Statement{
						Line(),
						Return(Add(value).Dot("DeepCopy").Params()),
					}
				}))

				statement = append(statement, Line())
				statement = append(statement, mapcasegen.AllCases(allType, value, mainFunc)...)
				statement = append(statement, Line())
				statement = append(statement, slicecasegen.AllCases(allType, value, mainFunc)...)

				return statement
			}),
			Line(),
			Return(Add(object)),
		)

	return f.Render(w)
}
