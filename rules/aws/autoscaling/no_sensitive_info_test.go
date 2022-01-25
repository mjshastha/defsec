package autoscaling

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/aws/autoscaling"
	"github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckNoSensitiveInfo(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name     string
		input    autoscaling.Autoscaling
		expected bool
	}{
		{
			name:     "positive result",
			input:    autoscaling.Autoscaling{},
			expected: true,
		},
		{
			name:     "negative result",
			input:    autoscaling.Autoscaling{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.Autoscaling = test.input
			results := CheckNoSensitiveInfo.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Rule().LongID() == CheckNoSensitiveInfo.Rule().LongID() {
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