package outer

import (
	_ "unsafe"

	_ "github.com/he2121/demos/linkname_example/inner"
)

//go:linkname A github.com/he2121/demos/linkname_example/inner.a
var A int

//go:linkname Hello github.com/he2121/demos/linkname_example/inner.hello
func Hello()
