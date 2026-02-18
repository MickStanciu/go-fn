package e2e

import "github.com/MickStanciu/go-fn/e2e/internal"

type StoreRequirementResultFn func(key string, value string)
type GetRequirementResultFn func(key string) (string, bool)

type TestSuite struct {
	name                  string
	tests                 []*E2ETest
	prerequisites         []*E2EDependency
	prerequisitesExecuted map[string]bool
	storage               *internal.Storage
}

func (ts *TestSuite) GetRequirementResult(key string) (string, bool) {
	return ts.storage.GetArtifact(key)
}

func (ts *TestSuite) StoreRequirementResult(key string, value string) {
	ts.storage.SetArtifact(key, value)
}

type TestSuiteOption func(*TestSuite)

func NewTestSuite(suiteName string, opts ...TestSuiteOption) *TestSuite {
	ts := &TestSuite{
		name:                  suiteName,
		tests:                 make([]*E2ETest, 0),
		prerequisites:         make([]*E2EDependency, 0),
		prerequisitesExecuted: make(map[string]bool),
		storage:               internal.NewStorage(),
	}

	for _, opt := range opts {
		opt(ts)
	}

	return ts
}

func WithVariables(vars map[string]string) TestSuiteOption {
	return func(ts *TestSuite) {
		for k, v := range vars {
			ts.storage.SetArtifact(k, v)
		}
	}
}

func WithPreRequirement(name string, configuration *E2EDependencyConfiguration) TestSuiteOption {
	return func(ts *TestSuite) {
		ts.prerequisites = append(ts.prerequisites, &E2EDependency{
			name:          name,
			execFn:        configuration.ExecFn,
			configuration: configuration.Configuration,
		})
		ts.prerequisitesExecuted[name] = false
	}
}

func WithTest(name string, configuration *E2ETestConfiguration) TestSuiteOption {
	return func(ts *TestSuite) {
		ts.tests = append(ts.tests, &E2ETest{
			name:          name,
			enabled:       configuration.Enabled,
			execFn:        configuration.ExecFn,
			configuration: configuration.Configuration,
			dependsOn:     configuration.DependsOnResults,
		})
	}
}

type E2ETest struct {
	name          string
	enabled       bool
	execFn        func(*E2ETestParams) error
	configuration map[string]string
	dependsOn     []string
}

// E2ETestParams - passed into the test execution function
type E2ETestParams struct {
	GetValueFromSuiteStoreFn GetRequirementResultFn
	SetValueInSuiteStoreFn   StoreRequirementResultFn
	LocalConfiguration       map[string]string
}

// E2ETestConfiguration - configuration for an end-to-end test
type E2ETestConfiguration struct {
	Enabled          bool
	Configuration    map[string]string
	DependsOnResults []string
	ExecFn           func(*E2ETestParams) error
}

// E2EDependency - represents a prerequisite that must be executed before tests
type E2EDependency struct {
	name          string
	configuration map[string]string
	execFn        func(*E2ETestParams) error
}

// E2EDependencyConfiguration - configuration for an end-to-end dependency
type E2EDependencyConfiguration struct {
	Configuration map[string]string
	ExecFn        func(*E2ETestParams) error
}
