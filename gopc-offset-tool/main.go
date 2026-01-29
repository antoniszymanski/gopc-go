package main

import (
	"fmt"
	"go/types"
	"os"
	"slices"

	"github.com/antoniszymanski/loadpackage-go"
	"golang.org/x/tools/go/packages"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run() error {
	offset, err := getOffset()
	if err != nil {
		return err
	}
	fmt.Println("+--------+-----------+")
	fmt.Println("| Arch   | Offset    |")
	fmt.Println("+--------+-----------+")
	fmt.Printf("| 32-bit | %-9d |\n", offset.Arch32)
	fmt.Printf("| 64-bit | %-9d |\n", offset.Arch64)
	fmt.Println("+--------+-----------+")
	return nil
}

func getOffset() (Offset, error) {
	pkg, err := loadpackage.Load("runtime", &packages.Config{Mode: packages.LoadSyntax})
	if err != nil {
		return Offset{}, fmt.Errorf("failed to load the runtime package: %w", err)
	}
	var typ *types.Struct
	for _, obj := range pkg.TypesInfo.Defs {
		obj, _ := obj.(*types.TypeName)
		if obj == nil {
			continue
		}
		if obj.Name() != "g" {
			continue
		}
		typ, _ = obj.Type().Underlying().(*types.Struct)
		if typ == nil {
			continue
		}
	}
	if typ == nil {
		return Offset{}, fmt.Errorf("failed to find the runtime.g struct")
	}
	fields := slices.Collect(typ.Fields())
	for i, field := range fields {
		if field.Name() != "gopc" {
			continue
		}
		return Offset{
			Arch32: types.SizesFor("gc", "386").Offsetsof(fields)[i],
			Arch64: types.SizesFor("gc", "amd64").Offsetsof(fields)[i],
		}, nil
	}
	return Offset{}, fmt.Errorf("failed to find the gopc field in the runtime.g struct")
}

type Offset struct {
	Arch32, Arch64 int64
}
