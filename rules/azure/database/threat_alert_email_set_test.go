package database

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/azure/database"
	"github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckThreatAlertEmailSet(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name     string
		input    database.Database
		expected bool
	}{
		{
			name:     "positive result",
			input:    database.Database{},
			expected: true,
		},
		{
			name:     "negative result",
			input:    database.Database{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.Azure.Database = test.input
			results := CheckThreatAlertEmailSet.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Rule().LongID() == CheckThreatAlertEmailSet.Rule().LongID() {
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