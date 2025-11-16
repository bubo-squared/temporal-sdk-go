package mocks

import (
	"github.com/bubo-squared/temporal-go-sdk/client"
	"github.com/bubo-squared/temporal-go-sdk/converter"
)

// make sure mocks are in sync with interfaces
var (
	_ client.Client               = (*Client)(nil)
	_ client.HistoryEventIterator = (*HistoryEventIterator)(nil)
	_ client.NamespaceClient      = (*NamespaceClient)(nil)
	_ converter.EncodedValue      = (*Value)(nil)
	_ client.WorkflowRun          = (*WorkflowRun)(nil)
)
