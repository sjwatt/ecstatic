package main

import (
	"exp/types"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// Conversion between ast.Visitor and the simpler type func(ast.Node) (bool).
type SimpleVisitor struct {
	pack *ast.Package
	fset *token.FileSet
}

// Returns itself if 'visitFunc' returns true - otherwise returns nil.
func (v SimpleVisitor) Visit(node ast.Node) ast.Visitor {
		switch a := node.(type) {
		case *ast.ExprStmt:
			//fmt.Printf("%#v\n",a.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X)
			switch b := a.X.(*ast.CallExpr).Fun.(type){
			case *ast.Ident:
				ident:=v.pack.Scope.Lookup(b.Name)
				//fmt.Printf("%#v", ident)
				if ident != nil && ident.Decl.(*ast.FuncDecl).Type.Results != nil{
					fmt.Println(ident.Name," has unassigned return values",v.fset.Position(ident.Pos()))
				}
			case *ast.SelectorExpr:
				X := a.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X.(*ast.Ident).Name
				I := a.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name
				//fmt.Println(X+"."+I)
				f := v.pack.Imports[X].Data.(*ast.Scope).Lookup(I)
				//fmt.Printf("%#v\n",v.pack.Scope.Objects)
				//fmt.Printf("%#v\n",f)
				if f != nil && f.Type.(*types.Func).Results != nil{
					fmt.Println(X+"."+I," has unassigned return values",v.fset.Position(a.Pos()))
				}
			}
		}
		return v		
}
func main() {
	fset := token.NewFileSet()
	files := map[string]*ast.File{}

	for _,f := range os.Args[1:] {
		file, err := parser.ParseFile(fset, f, nil, 0)
		if err != nil {
			panic("parse file:" + err.Error())
		}
		files[f] = file
		
	}
	pack, err := ast.NewPackage(fset, files, types.GcImporter, types.Universe)
	if err != nil {
		panic("new package" + err.Error())
	}
	/*for name, object := range pack.Scope.Objects {
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
	}*/

	//check that the file is valid
	/*valid := true
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
	}*/

	ast.Walk(SimpleVisitor{pack,fset},pack)
}

/*
find all function calls
check return values against assignment

*/
