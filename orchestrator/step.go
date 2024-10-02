package orchestrator

// Step - struct that will contain instructions of Step execution and rollback
type Step struct {
	name           string
	stepFn         execFn
	rollbackStepFn execFn
	//retryStrategy - nice to have
	timeOutSeconds int
}

type BuildStepOption func(*Step)

// WithTimeout - adds a timeout to the function to be executed
func WithTimeout(timeOutSeconds int) BuildStepOption {
	return func(s *Step) {
		s.timeOutSeconds = timeOutSeconds
	}
}

// WithRollbackStepFn - adds a rollback function
func WithRollbackStepFn(fn execFn) BuildStepOption {
	return func(s *Step) {
		if fn != nil {
			s.rollbackStepFn = fn
		}
	}
}

// NewStep - builds a new step
func NewStep(name string, fn execFn, opts ...BuildStepOption) *Step {
	s := &Step{
		name:           name,
		stepFn:         fn,
		timeOutSeconds: 0,
	}

	// process each option
	for _, opt := range opts {
		opt(s)
	}

	return s
}

// ExecStep - executes the step
func (s *Step) ExecStep() error {
	if s.timeOutSeconds == 0 {
		return s.stepFn()
	}

	// executing with timeout
	return ExecuteWithTimeout(s.stepFn, s.name, s.timeOutSeconds)
}

// ExecRollback - executes the rollback step
func (s *Step) ExecRollback() error {
	if s.rollbackStepFn == nil {
		return nil
	}

	if s.timeOutSeconds == 0 {
		return s.rollbackStepFn()
	}

	// executing with timeout
	return ExecuteWithTimeout(s.rollbackStepFn, s.name, s.timeOutSeconds)
}
