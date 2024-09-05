package stack_test

import (
	"fmt"
	"github.com/MickStanciu/go-fn/stack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrchestrator_StartExecution_ShouldExecuteTheStepsWithoutRollback(t *testing.T) {
	o := stack.NewOrchestrator()

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

	o.AddStep("step 1", step1, nil)
	o.AddStep("step 2", step2, step2RollBack)
	err := o.StartExecution()
	require.NoError(t, err)
}

func TestOrchestrator_StartExecution_ShouldExecuteTheStepsWithRollback(t *testing.T) {
	o := stack.NewOrchestrator()

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

	o.AddStep("step 1", step1, nil)
	o.AddStep("step 2", step2, step2RollBack)
	o.AddStep("step 3", step3, nil)
	err := o.StartExecution()
	require.Error(t, err)
	assert.Equal(t, err.Error(), `error executing step "step 3": error in step3`)
}

func TestOrchestrator_StartExecution_WhenRollbackFails(t *testing.T) {
	o := stack.NewOrchestrator()

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

	o.AddStep("step 1", step1, step1RollBack)
	o.AddStep("step 2", step2, step2RollBack)
	o.AddStep("step 3", step3, nil)
	err := o.StartExecution()
	require.Error(t, err)
	assert.Equal(t, err.Error(), "error executing step \"step 3\": error in step3\nerror rolling back step \"step 1\": error rollback in step1")
}
