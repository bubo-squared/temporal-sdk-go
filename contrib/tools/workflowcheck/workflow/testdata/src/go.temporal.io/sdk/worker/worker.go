package worker

import (
	"github.com/bubo-squared/temporal-go-sdk/activity"
	"github.com/bubo-squared/temporal-go-sdk/workflow"
)

type Worker interface {
	Registry
}

type Registry interface {
	WorkflowRegistry
	ActivityRegistry
}

type WorkflowRegistry interface {
	RegisterWorkflow(w interface{})
	RegisterWorkflowWithOptions(w interface{}, options workflow.RegisterOptions)
	RegisterDynamicWorkflow(w interface{}, options workflow.DynamicRegisterOptions)
}

type ActivityRegistry interface {
	RegisterActivity(a interface{})
	RegisterActivityWithOptions(a interface{}, options activity.RegisterOptions)
	RegisterDynamicActivity(a interface{}, options activity.DynamicRegisterOptions)
}
