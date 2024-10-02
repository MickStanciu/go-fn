package orchestrator_test

import (
	"fmt"
	"github.com/MickStanciu/go-fn/orchestrator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrchestrator_StartExecution_ShouldExecuteTheStepsWithoutRollback(t *testing.T) {
	o := orchestrator.NewOrchestrator()

	step1 := func() error {
		fmt.Println("executing step 1")
		return nil
	}

	step2 := func() error {
		fmt.Println("executing step 2")
		return nil
	}

	step2RollBack := func() error {
		fmt.Println("executing rollback for step 2 ")
		return nil
	}

	o.AddStep(orchestrator.NewStep("step 1", step1))
	o.AddStep(orchestrator.NewStep("step 2", step2, orchestrator.WithRollbackStepFn(step2RollBack)))
	err := o.Run()
	require.NoError(t, err)
}

func TestOrchestrator_StartExecution_ShouldExecuteTheStepsWithRollback(t *testing.T) {
	o := orchestrator.NewOrchestrator()

	step1 := func() error {
		fmt.Println("executing step 1")
		return nil
	}

	step2 := func() error {
		fmt.Println("executing step 2")
		return nil
	}

	step2RollBack := func() error {
		fmt.Println("executing rollback for step 2 ")
		return nil
	}

	step3 := func() error {
		fmt.Println("executing step 3")
		return fmt.Errorf("error in step3")
	}

	o.AddStep(orchestrator.NewStep("step 1", step1))
	o.AddStep(orchestrator.NewStep("step 2", step2, orchestrator.WithRollbackStepFn(step2RollBack)))
	o.AddStep(orchestrator.NewStep("step 3", step3))
	err := o.Run()
	require.Error(t, err)
	assert.Equal(t, err.Error(), `error executing step "step 3": error in step3`)
}

func TestOrchestrator_StartExecution_WhenRollbackFails(t *testing.T) {
	o := orchestrator.NewOrchestrator()

	step1 := func() error {
		fmt.Println("executing step 1")
		return nil
	}

	step1RollBack := func() error {
		return fmt.Errorf("error rollback in step1")
	}

	step2 := func() error {
		fmt.Println("executing step 2")
		return nil
	}

	step2RollBack := func() error {
		fmt.Println("executing rollback for step 2 ")
		return nil
	}

	step3 := func() error {
		fmt.Println("executing step 3")
		return fmt.Errorf("error in step3")
	}

	o.AddStep(orchestrator.NewStep("step 1", step1, orchestrator.WithRollbackStepFn(step1RollBack)))
	o.AddStep(orchestrator.NewStep("step 2", step2, orchestrator.WithRollbackStepFn(step2RollBack)))
	o.AddStep(orchestrator.NewStep("step 3", step3))
	err := o.Run()
	require.Error(t, err)
	assert.Equal(t, err.Error(), "error executing step \"step 3\": error in step3\nerror rolling back step \"step 1\": error rollback in step1")
}

// Test disabled because of the execution time
//func TestOrchestrator_StartExecution_WithTimeOut(t *testing.T) {
//	o := orchestrator.NewOrchestrator()
//
//	step1 := func() error {
//		fmt.Println("executing step 1")
//		time.Sleep(3 * time.Second)
//		return nil
//	}
//
//	o.AddStep(orchestrator.NewStep("step 1", step1, orchestrator.WithTimeout(2)))
//	err := o.Run()
//	require.NoError(t, err)
//}
