package main

import (
	"github.com/blanchonvincent/ctxarg/analysis/passes/ctxarg"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(ctxarg.Analyzer) }
