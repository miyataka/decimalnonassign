package main

import (
	"github.com/miyataka/decimalnonassign"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(decimalnonassign.Analyzer) }
