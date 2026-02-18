package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (ts *TestService) Run(t *testing.T) {
	for _, suite := range ts.suites {
		t.Log("Running collection: ", suite.name)

		// execute pre-requirements
		for _, req := range suite.prerequisites {
			executed := suite.prerequisitesExecuted[req.name]
			if executed {
				continue
			}

			t.Log("> running pre-requirement:", req.name)
			err := req.execFn(&E2ETestParams{
				GetValueFromSuiteStoreFn: suite.GetRequirementResult,
				SetValueInSuiteStoreFn:   suite.StoreRequirementResult,
				LocalConfiguration:       req.configuration,
			})
			assert.NoError(t, err, "pre-requirement:", req.name)
			suite.prerequisitesExecuted[req.name] = true
		}

		for _, tc := range suite.tests {
			// check if test is enabled
			if tc.enabled == false {
				t.Log("> skipping test:", tc.name)
				continue
			}

			// check prerequirements
			for _, reqName := range tc.dependsOn {
				executed := suite.prerequisitesExecuted[reqName]
				require.True(t, executed, "missing prerequisite result:", reqName)
			}

			t.Log("> running test:", tc.name)
			t.Run(tc.name, func(t *testing.T) {
				err := tc.execFn(&E2ETestParams{
					GetValueFromSuiteStoreFn: suite.GetRequirementResult,
					SetValueInSuiteStoreFn:   suite.StoreRequirementResult,
					LocalConfiguration:       tc.configuration,
				})
				assert.NoError(t, err, "test:", tc.name)
			})

		}
	}
}
