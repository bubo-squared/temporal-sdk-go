package main

import (
	"github.com/bubo-squared/temporal-go-sdk/contrib/tools/workflowcheck/workflow"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(workflow.NewChecker(workflow.Config{}).NewAnalyzer())
}
