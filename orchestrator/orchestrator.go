package orchestrator

import (
	"errors"
	"fmt"
)

// Orchestrator - struct that will contain an array of steps to be executed
type Orchestrator struct {
	steps []*Step
}

// NewOrchestrator - creates an Orchestrator
func NewOrchestrator(opts ...BuildOrchestratorOption) *Orchestrator {
	o := &Orchestrator{}

	// process each option
	for _, opt := range opts {
		opt(o)
	}

	// fall back
	if o.steps == nil {
		o.steps = []*Step{}
	}
	return o
}

type BuildOrchestratorOption func(*Orchestrator)

// AddStep - adds a new execution step
func (o *Orchestrator) AddStep(step *Step) {
	if step == nil {
		return
	}

	if step.stepFn == nil {
		panic(errors.New("stepFn cannot be nil"))
	}

	o.steps = append(o.steps, step)
}

// Run - starts running all the steps in order
func (o *Orchestrator) Run() error {
	var errStepIdx = 0
	var lastErr error = nil

	// executing steps one by one
	for idx, s := range o.steps {
		if err := s.ExecStep(); err != nil {
			lastErr = fmt.Errorf("error executing step %q: %w", s.name, err)
			errStepIdx = idx
			break
		}
	}

	// if there was an error, we need to rollback previous steps
	for idx := errStepIdx; idx >= 0; idx-- {
		if err := o.steps[idx].ExecRollback(); err != nil {
			return errors.Join(lastErr, fmt.Errorf("error rolling back step %q: %w", o.steps[idx].name, err))
		}
	}

	return lastErr
}
