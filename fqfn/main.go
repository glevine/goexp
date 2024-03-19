package main

import (
	"fmt"

	"github.com/glevine/goexp/fqfn/pkg1"
	"github.com/glevine/goexp/fqfn/pkg2"
	"github.com/glevine/goexp/fqfn/pkg2/pkg3"
)

func main() {
	fmt.Println(pkg1.Func())
	fmt.Println(pkg2.Func())
	fmt.Println(pkg3.Func())
}
