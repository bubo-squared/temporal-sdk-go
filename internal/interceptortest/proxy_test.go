package interceptortest_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/bubo-squared/temporal-sdk-go/activity"
	"github.com/bubo-squared/temporal-sdk-go/interceptor"
	"github.com/bubo-squared/temporal-sdk-go/internal/interceptortest"
	"github.com/bubo-squared/temporal-sdk-go/testsuite"
	"github.com/bubo-squared/temporal-sdk-go/worker"
	"github.com/bubo-squared/temporal-sdk-go/workflow"
)

func TestProxy(t *testing.T) {
	// Just a sanity check to make sure proxy works

	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	env.RegisterWorkflow(ProxyWorkflow)
	env.RegisterActivity(ProxyActivity)

	// Set recorder
	var rec interceptortest.CallRecordingInvoker
	proxy := interceptortest.NewProxy(&rec)
	env.SetWorkerOptions(worker.Options{Interceptors: []interceptor.WorkerInterceptor{proxy}})

	// Exec
	env.ExecuteWorkflow(ProxyWorkflow, "World")

	// Confirm result
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var result string
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, "Hello, World", result)

	// Make sure expected methods are present
	calls := rec.Calls()
	getCall := func(qualifiedMethod string) *interceptortest.RecordedCall {
		for _, call := range calls {
			if call.Interface.Name()+"."+call.Method.Name == qualifiedMethod {
				return call
			}
		}
		return nil
	}
	require.NotNil(t, getCall("ActivityInboundInterceptor.Init"))
	call := getCall("ActivityInboundInterceptor.ExecuteActivity")
	require.NotNil(t, call)
	require.Equal(t, "World", call.Args[1].Interface().(*interceptor.ExecuteActivityInput).Args[0])
	call = getCall("ActivityOutboundInterceptor.GetInfo")
	require.NotNil(t, call)
	require.Equal(t, "ProxyActivity", call.Results[0].Interface().(activity.Info).ActivityType.Name)
}

func ProxyWorkflow(ctx workflow.Context, suffix string) (ret string, err error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{StartToCloseTimeout: 10 * time.Second})
	err = workflow.ExecuteActivity(ctx, ProxyActivity, suffix).Get(ctx, &ret)
	return
}

func ProxyActivity(ctx context.Context, suffix string) (string, error) {
	activity.GetInfo(ctx)
	return "Hello, " + suffix, nil
}
