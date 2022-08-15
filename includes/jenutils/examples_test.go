package jenutils

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

func Example_dereference() {
	_, c := dereference(jen.Index().String(), jen.Id("object"))
	fmt.Printf("%s", c.GoString())
	// Output:
	// var value []string
	// if val, ok := object.(*[]string); ok {
	//	value = *val
	// } else {
	//	value = object.([]string)
	// }
}

func ExampleCaseType() {
	switchValue := jen.Id("t")
	c := jen.Switch(jen.Add(switchValue).Op(":=").Id("object").Assert(jen.Type())).Block(
		CaseType(jen.Map(jen.String()).Interface(), switchValue, func(value, t jen.Code) jen.Statement {
			return jen.Statement{jen.Line(), jen.Return(value)}
		}),
	)
	fmt.Printf("%#v", c)
	// Output:
	// switch t := object.(type) {
	// case map[string]interface{}, *map[string]interface{}:
	//	var value map[string]interface{}
	//
	//	if val, ok := t.(*map[string]interface{}); ok {
	//		value = *val
	//	} else {
	//		value = t.(map[string]interface{})
	//	}
	//
	//	return value
	// }
}

func ExampleCases() {
	switchValue := jen.Id("t")
	c := jen.Switch(jen.Add(switchValue).Op(":=").Id("object").Assert(jen.Type())).Block(
		Cases(jen.Statement{jen.Bool(), jen.Int(), jen.String()}, switchValue, func(value, t jen.Code) jen.Statement {
			return jen.Statement{jen.Line(), jen.Return(value)}
		}, jen.Index())...,
	)
	fmt.Printf("%#v", c)
	// Output:
	// switch t := object.(type) {
	// case []bool, *[]bool:
	//	var value []bool
	//
	//	if val, ok := t.(*[]bool); ok {
	//		value = *val
	//	} else {
	//		value = t.([]bool)
	//	}
	//
	//	return value
	// case []int, *[]int:
	//	var value []int
	//
	//	if val, ok := t.(*[]int); ok {
	//		value = *val
	//	} else {
	//		value = t.([]int)
	//	}
	//
	//	return value
	// case []string, *[]string:
	//	var value []string
	//
	//	if val, ok := t.(*[]string); ok {
	//		value = *val
	//	} else {
	//		value = t.([]string)
	//	}
	//
	//	return value
	// }
}

func ExampleCases_notBaseType() {
	switchValue := jen.Id("t")
	c := jen.Switch(jen.Add(switchValue).Op(":=").Id("object").Assert(jen.Type())).Block(
		Cases(jen.Statement{jen.Bool(), jen.Int(), jen.String()}, switchValue, func(value, t jen.Code) jen.Statement {
			return jen.Statement{jen.Line(), jen.Return(value)}
		})...,
	)
	fmt.Printf("%#v", c)
	// Output:
	// switch t := object.(type) {
	// case bool, *bool:
	//	var value bool
	//
	//	if val, ok := t.(*bool); ok {
	//		value = *val
	//	} else {
	//		value = t.(bool)
	//	}
	//
	//	return value
	// case int, *int:
	//	var value int
	//
	//	if val, ok := t.(*int); ok {
	//		value = *val
	//	} else {
	//		value = t.(int)
	//	}
	//
	//	return value
	// case string, *string:
	//	var value string
	//
	//	if val, ok := t.(*string); ok {
	//		value = *val
	//	} else {
	//		value = t.(string)
	//	}
	//
	//	return value
	// }
}

func ExampleForeach() {
	c := Foreach(jen.Id("object")).Block(
		jen.Id("newObject").Index(jen.Id("k")).Op("=").Id("v"),
	)
	fmt.Printf("%#v", c)
	// Output:
	// for k, v := range object {
	// 	newObject[k] = v
	// }
}

func ExampleForeach_customKeyValue() {
	c := Foreach(jen.Id("object"), "key", "value").Block(
		jen.Id("newObject").Index(jen.Id("key")).Op("=").Id("value"),
	)
	fmt.Printf("%#v", c)
	// Output:
	// for key, value := range object {
	// 	newObject[key] = value
	// }
}

func ExampleForeach_notKey() {
	c := Foreach(jen.Id("object"), "_").Block(
		jen.Id("newObject").Op("=").Append(jen.Id("newObject"), jen.Id("v")),
	)
	fmt.Printf("%#v", c)
	// Output:
	// for _, v := range object {
	// 	newObject = append(newObject, v)
	// }
}

func ExampleSwitchType() {
	c := SwitchType(jen.Id("object"), jen.Id("t"), func(value jen.Code) jen.Statement {
		return jen.Statement{jen.Case(jen.String()).Block(jen.Return(value))}
	})
	fmt.Printf("%#v", c)
	// Output:
	// switch t := object.(type) {
	// case string:
	//	return t
	// }
}

func ExampleShallowCopy() {
	c := ShallowCopy(jen.String(), jen.Id("origin"), jen.Id("clone"))
	fmt.Printf("%s", c.GoString())
	// Output:
	// clone := make(string, len(origin))
	// for k, v := range origin {
	//	clone[k] = v
	// }
}

func ExampleDeepCopy() {
	c := DeepCopy(jen.Map(jen.String()).Interface(), jen.Id("origin"), jen.Id("clone"), jen.Id("deepCopy"))
	fmt.Printf("%s", c.GoString())
	// Output:
	// clone := make(map[string]interface{}, len(origin))
	// for k, v := range origin {
	//	clone[k] = deepCopy(v).(interface{})
	// }
}

func ExampleInterfaceCopy() {
	c := InterfaceCopy(jen.Map(jen.String()).Interface(), jen.Id("origin"), jen.Id("clone"), jen.Id("deepCopy"))
	fmt.Printf("%s", c.GoString())
	// Output:
	// clone := make(map[string]interface{}, len(origin))
	// for k, v := range origin {
	//	clone[k] = deepCopy(v)
	// }
}
