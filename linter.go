// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"

	"golang.org/x/exp/slices"
)

func lint(keywords []string, sources []string) int {
	totalErrors := 0
	files := resolveFiles(sources)
	for _, file := range files {
		src, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, file, src, 0)
		if err != nil {
			log.Fatal(err)
		}

		visitor := &Visitor{
			fset:     fset,
			keywords: keywords,
			errors:   0,
		}
		ast.Walk(visitor, f)
		totalErrors = totalErrors + visitor.errors
	}
	return totalErrors
}

type Visitor struct {
	fset     *token.FileSet
	keywords []string
	errors   int
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	switch x := node.(type) {
	case *ast.CallExpr:
		name := getCallName(x)
		if slices.Contains(v.keywords, name) {
			pos := v.fset.Position(node.Pos())
			for _, arg := range x.Args {
				if lit, ok := arg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
					for _, rule := range allRules {
						filePos := fmt.Sprintf("%s:%d", pos.Filename, pos.Line)
						if matching, message := rule.Check(lit.Value, filePos); matching {
							fmt.Fprintf(os.Stderr, "%s (%s) %s%s\n", filePos, rule.Name(), lit.Value, message)
							v.errors = v.errors + 1
						}
					}
				}
			}
		}
	}
	return v
}

func getCallName(call *ast.CallExpr) string {
	switch x := call.Fun.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.SelectorExpr:
		return x.Sel.Name
	}
	return ""
}
