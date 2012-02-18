package main

import (
	"exp/types"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// Conversion between ast.Visitor and the simpler type func(ast.Node) (bool).
type SimpleVisitor struct {
	visitFunc func(node ast.Node) bool
}

// Returns itself if 'visitFunc' returns true - otherwise returns nil.
func (v SimpleVisitor) Visit(node ast.Node) ast.Visitor {
	if v.visitFunc(node) {
		return v
	}
	return nil
}

func visitor(visit func(node ast.Node) bool) ast.Visitor {
	return SimpleVisitor{visit}
}

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "./testdata/test.go", nil, 0)
	if err != nil {
		panic("parse file:" + err.Error())
	}

	pack, err := ast.NewPackage(fset, map[string]*ast.File{"./testdata/test.go": file}, types.GcImporter, types.Universe)
	if err != nil {
		panic("new package" + err.Error())
	}
	// validate the parse
	//fmt.Println("name:",pack.Name)
	//fmt.Println("println:",pack.Scope.Lookup("Println").Name)
	//fmt.Println("println:",pack.Scope.Lookup("Println").Type)
	//fmt.Printf("%#v\n",pack.Scope.Lookup("Println").Decl.(*ast.FuncDecl).Type.Results.List[0].Type.(*ast.Ident))
	for name, object := range pack.Scope.Objects {
		if f, ok := object.Decl.(*ast.FuncDecl); ok {
			if f.Type.Results == nil {
				continue
			}
			for _, returnValue := range f.Type.Results.List {
				if returnValue.Type.(*ast.Ident).Name == "error" {
					fmt.Println(name, ": function returns error")
				}
			}
		}
	}

	fmt.Println("imports:", pack.Imports)
	_ = pack

	//check that the file is valid
	valid := true
	ast.Walk(visitor(func(node ast.Node) bool {
		switch node.(type) {
		case *ast.BadDecl,
			*ast.BadExpr,
			*ast.BadStmt:
			valid = false
			return false
		}
		return true
	}), pack)
	if !valid {
		panic("invalid file")
	}

	//fmt.Printf("%#v\n",call.X.(*ast.CallExpr).Fun)
	//fmt.Printf("%#v\n",call.Rhs[0].(*ast.CallExpr).Fun.(*ast.Ident).Name)

	ast.Walk(visitor(func(node ast.Node) bool {
		switch a := node.(type) {
		case *ast.ExprStmt:
			fmt.Printf("%#v\n",a.X.(*ast.CallExpr).Fun)
		case *ast.AssignStmt:
			if call,ok := a.Rhs[0].(*ast.CallExpr); ok {
				if selector, ok := call.Fun.(*ast.SelectorExpr); ok {
					_ = selector
					//handle selector	
				} else {
					fmt.Printf("%#v\n",call.Fun.(*ast.Ident).Name)
				}
			}
		case *ast.GoStmt:
			fmt.Printf("%#v\n", a)
		case *ast.ForStmt:
			fmt.Printf("%#v\n", a)
		case *ast.DeferStmt:
			fmt.Printf("%#v\n", a)
		case *ast.RangeStmt:
			fmt.Printf("%#v\n", a)
		case *ast.ReturnStmt:
			fmt.Printf("%#v\n", a)
		}
		return true
	}), pack)

	//fmt.Println("fset:",fset)
	//fmt.Println("packages:",packages["main"])
	//m, err := types.Check(fset, packages["main"])
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(m)
}

/*
find all function calls
check return values against assignment

*/
