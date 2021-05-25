package main

import (
	"github.com/fdymylja/tmos/pkg/modulecodegen"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(modulecodegen.PluginRunner)
}
