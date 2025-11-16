// This file exists to force compilation of all code that doesn't have unit tests.
package main

import (
	_ "github.com/bubo-squared/temporal-go-sdk/activity"
	_ "github.com/bubo-squared/temporal-go-sdk/client"
	_ "github.com/bubo-squared/temporal-go-sdk/converter"
	_ "github.com/bubo-squared/temporal-go-sdk/log"
	_ "github.com/bubo-squared/temporal-go-sdk/temporal"
	_ "github.com/bubo-squared/temporal-go-sdk/testsuite"
	_ "github.com/bubo-squared/temporal-go-sdk/worker"
	_ "github.com/bubo-squared/temporal-go-sdk/workflow"
)

func main() {
}
