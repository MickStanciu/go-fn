package stack

import (
	"errors"
	"fmt"
)

type execBlock struct {
	name           string
	stepFn         func() error
	rollbackStepFn func() error
}

type Orchestrator struct {
	content []*execBlock
}

func NewOrchestrator() *Orchestrator {
	o := &Orchestrator{}
	o.content = []*execBlock{}
	return o
}

func (o *Orchestrator) AddStep(name string, stepFn func() error, rollbackFn func() error) {
	o.content = append(o.content, &execBlock{
		name:           name,
		stepFn:         stepFn,
		rollbackStepFn: rollbackFn,
	})
}

func (o *Orchestrator) StartExecution() error {
	var rollBackIdx = 0
	var lastErr error = nil

	for n := 0; n < len(o.content); n++ {
		err := o.content[n].stepFn()
		if err != nil {
			lastErr = fmt.Errorf("error executing step %q: %w", o.content[n].name, err)
			rollBackIdx = n
			break
		}
	}

	for n := rollBackIdx; n >= 0; n-- {
		// skip if we don't have a rollback fn
		if o.content[n].rollbackStepFn == nil {
			continue
		}

		err := o.content[n].rollbackStepFn()
		if err != nil {
			return errors.Join(lastErr, fmt.Errorf("error rolling back step %q: %w", o.content[n].name, err))
		}
	}

	return lastErr
}
