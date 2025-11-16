// This file exists to force compilation of all code that doesn't have unit tests.
package main

import (
	_ "github.com/bubo-squared/temporal-sdk-go/activity"
	_ "github.com/bubo-squared/temporal-sdk-go/client"
	_ "github.com/bubo-squared/temporal-sdk-go/converter"
	_ "github.com/bubo-squared/temporal-sdk-go/log"
	_ "github.com/bubo-squared/temporal-sdk-go/temporal"
	_ "github.com/bubo-squared/temporal-sdk-go/testsuite"
	_ "github.com/bubo-squared/temporal-sdk-go/worker"
	_ "github.com/bubo-squared/temporal-sdk-go/workflow"
)

func main() {
}
