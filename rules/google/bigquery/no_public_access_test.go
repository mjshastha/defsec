package bigquery

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/google/bigquery"
	"github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckNoPublicAccess(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name     string
		input    bigquery.BigQuery
		expected bool
	}{
		{
			name:     "positive result",
			input:    bigquery.BigQuery{},
			expected: true,
		},
		{
			name:     "negative result",
			input:    bigquery.BigQuery{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.Google.BigQuery = test.input
			results := CheckNoPublicAccess.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Rule().LongID() == CheckNoPublicAccess.Rule().LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}