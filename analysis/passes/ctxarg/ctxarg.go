package ctxarg

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `check for parameters order while receiving context as parameter

The ctxarg checker walks functions checking context parameters 
and ensures the context parameter is always the first received argument.`

var Analyzer = &analysis.Analyzer{
	Name:     "ctxarg",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

var name, paramorder *bool

func init() {
	name = Analyzer.Flags.Bool("name", true, "will ensure context as parameter is named ctx")
	paramorder = Analyzer.Flags.Bool("paramorder", true, "will ensure context is the first argument of the functions")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncType)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		function := n.(*ast.FuncType)

		numContext := 0
		for _, f := range function.Params.List {
			t := pass.TypesInfo.TypeOf(f.Type)
			if `context.Context` == t.String() {
				numContext++
			}
		}
		if numContext > 1 {
			pass.Reportf(function.Pos(), "the function has more than one context defined in the parameters.")
			return
		}

		for key, f := range function.Params.List {
			// if the type is not a context, we do not need to go further
			t := pass.TypesInfo.TypeOf(f.Type)
			if `context.Context` != t.String() {
				continue
			}

			// if it is a context we will check two things:
			// - it is the first parameter
			// - the variable is named ctx
			if *paramorder {
				if 0 != key {
					pass.Reportf(function.Pos(), "the function has a context that is not the first argument.")
				}
			}
			if *name {
				if len(f.Names) > 1 {
					continue
				}
				if `ctx` != f.Names[0].Name {
					pass.Reportf(function.Pos(), "the function has a context that is not named 'ctx'.")
				}
			}
		}
	})
	return nil, nil
}
