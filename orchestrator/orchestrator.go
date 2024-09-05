package orchestrator

import (
	"errors"
	"fmt"
)

type stepFn func() error

// step - struct that will contain instructions of step execution and rollback
type step struct {
	name           string
	stepFn         stepFn
	rollbackStepFn stepFn
	//retryStrategy - nice to have
}

// Orchestrator - struct that will contain an array of steps to be executed
type Orchestrator struct {
	steps []*step
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
		o.steps = []*step{}
	}
	return o
}

type BuildOrchestratorOption func(*Orchestrator)

// AddStep - adds a new execution step
func (o *Orchestrator) AddStep(name string, stepFn stepFn, rollbackFn stepFn) {
	if stepFn == nil {
		panic(errors.New("stepFn cannot be nil"))
	}

	o.steps = append(o.steps, &step{
		name:           name,
		stepFn:         stepFn,
		rollbackStepFn: rollbackFn,
	})
}

// Run - starts running all the steps in order
func (o *Orchestrator) Run() error {
	var errStepIdx = 0
	var lastErr error = nil

	// executing steps one by one
	for idx, s := range o.steps {
		if err := s.stepFn(); err != nil {
			lastErr = fmt.Errorf("error executing step %q: %w", s.name, err)
			errStepIdx = idx
			break
		}
	}

	// if there was an error, we need to rollback previous steps
	for idx := errStepIdx; idx >= 0; idx-- {
		// skip if we don't have a rollback fn
		if o.steps[idx].rollbackStepFn == nil {
			continue
		}

		if err := o.steps[idx].rollbackStepFn(); err != nil {
			return errors.Join(lastErr, fmt.Errorf("error rolling back step %q: %w", o.steps[idx].name, err))
		}
	}

	return lastErr
}
