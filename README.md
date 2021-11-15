# DeepCopy

[![Go Reference](https://pkg.go.dev/badge/github.com/gohobby/deepcopy.svg)](https://pkg.go.dev/github.com/gohobby/deepcopy)
[![Go Report Card](https://goreportcard.com/badge/github.com/gohobby/deepcopy)](https://goreportcard.com/report/github.com/gohobby/deepcopy)
[![Go Cover](https://gocover.io/_badge/github.com/gohobby/deepcopy)](https://gocover.io/github.com/gohobby/deepcopy)
[![Github Action Test](https://github.com/gohobby/docs/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/gohobby/deepcopy/actions/workflows/golangci-lint.yml)

DeepCopy helps you create deep copies (clones) of your maps and slices.
Create deep copies (clones) of your objects
The package is based on [type assertions](https://golang.org/ref/spec#Type_assertions) and does not use reflection.

## Installation

Install DeepCopy with the [`go get`](https://pkg.go.dev/cmd/go#hdr-Add_dependencies_to_current_module_and_install_them)
command:

```shell
go get -u github.con/gohobby/deepcopy
```

## How it works

DeepCopy returns a new object with all recursively duplicated children. This means that changes made to the original
object will not affect the copied object and vice versa.

To copy a card or a slice:

```go
m := map[string]interface{}{"foo": []string{"bar", "baz"}}
cloneMap := deepcopy.Map(m).DeepCopy() // interface{}

s := []interface{}{1, 2, &m}
cloneSlice := deepcopy.Slice(s).DeepCopy() // interface{}
```

You can also use the Clone function to get the copy directly into the expected type, for example:

```go
m := map[string]interface{}{"foo": []string{"bar", "baz"}}
cloneMap := deepcopy.Map(m).CLone() // map[string]interface{}

s := []interface{}{1, 2, &m}
cloneSlice := deepcopy.Slice(s).Clone() // []interface{}
```

## Why?

### Mutability

Map types are pointers which make them mutable objects.

When you write the statement

`m := make(map[int]int)`

The compiler replaces it with a call
to [runtime.makemap](https://cs.opensource.google/go/go/+/master:src/runtime/map.go;l=304;drc=6f327f7b889b81549d551ce6963067267578bd70)
, which has the signature

    // makemap implements Go map creation for make(map[k]v, hint).
    // If the compiler has determined that the map or the first bucket
    // can be created on the stack, h and/or bucket may be non-nil.
    // If h != nil, the map can be created directly in h.
    // If h.buckets != nil, bucket pointed to can be used as the first bucket.
    func makemap(t *maptype, hint int, h *hmap) *hmap

As you can see, the type of the value returned by `runtime.makemap` is a pointer to
a [runtime.hmap](https://cs.opensource.google/go/go/+/master:src/runtime/map.go;l=116;drc=6f327f7b889b81549d551ce6963067267578bd70)
structure.

See [Dave Cheney's article](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it) for more
details.

Example:

```go
obj := map[string]int{"one": 1, "two": 2}

obj2 := obj

fmt.Printf("(obj)  %v\n(obj2) %v\n\n",
obj,  // mapcasegen[one:1 two:2]
obj2, // mapcasegen[one:1 two:2]
)

obj2["three"] = 3

fmt.Printf("(obj2) %v\n", obj2)
// mapcasegen[one:1 three:3 two:2] <-- ✅
fmt.Printf("(obj)  %v\n", obj)
// mapcasegen[one:1 three:3 two:2] <-- ❌
```

[Run this code in GoPlayground](https://play.golang.org/p/cLd5MJEagSI)

### How to create copies of your maps?

#### Shallow Copy

A shallow copy means that the first level is copied, deeper levels are referenced.

Example:

```go
type Map map[string]interface{}

var nestedMap = Map{
"flag": "🇫🇷",
"country": Map{
"city": "Paris",
},
}

// Shallow Copy
shallowClone := make(Map, len(nestedMap))

for k, v := range nestedMap {
shallowClone[k] = v
}

// Change of the cloned object
shallowClone["flag"] = "🇮🇹"
shallowClone["country"].(Map)["city"] = "Roma"

fmt.Printf("%v\n", shallowClone)
// mapcasegen[country:mapcasegen[city:Roma] flag:🇮🇹] <-- ✅

fmt.Printf("%v\n", nestedMap)
// mapcasegen[country:mapcasegen[city:Roma] flag:🇫🇷] <-- ❌

fmt.Printf("%p\n", nestedMap["country"]) // 0xc0000121e0
fmt.Printf("%p\n", shallowClone["country"]) // 0xc0000121e0
```

[Run this code in GoPlayground](https://play.golang.org/p/ZgGoyaC5hQa)

#### Deep Copy

A deep copy is a shallow copy applied recursively to all sub objects.

Example:

```go
deepClone := deepcopy.Map(nestedMap).Clone()

fmt.Printf("%v\n", deepClone)
// mapcasegen[country:mapcasegen[city:Roma] flag:🇮🇹] <-- ✅

fmt.Printf("%v\n", nestedMap)
// mapcasegen[country:mapcasegen[city:Paris] flag:🇫🇷] <-- ✅
```
