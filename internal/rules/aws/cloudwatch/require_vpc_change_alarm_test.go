package cloudwatch

import (
	"testing"

	"github.com/aquasecurity/defsec/internal/types"
	"github.com/aquasecurity/defsec/pkg/providers/aws/cloudtrail"
	"github.com/aquasecurity/defsec/pkg/providers/aws/cloudwatch"
	"github.com/aquasecurity/defsec/pkg/scan"
	"github.com/aquasecurity/defsec/pkg/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckVPCChangeAlarm(t *testing.T) {
	tests := []struct {
		name       string
		cloudtrail cloudtrail.CloudTrail
		cloudwatch cloudwatch.CloudWatch
		expected   bool
	}{
		{
			name: "Multi-region CloudTrail alarms on VPC changes",
			cloudtrail: cloudtrail.CloudTrail{
				Trails: []cloudtrail.Trail{
					{
						Metadata:                  types.NewTestMetadata(),
						CloudWatchLogsLogGroupArn: types.String("arn:aws:cloudwatch:us-east-1:123456789012:log-group:cloudtrail-logging", types.NewTestMetadata()),
						IsLogging:                 types.Bool(true, types.NewTestMetadata()),
						IsMultiRegion:             types.Bool(true, types.NewTestMetadata()),
					},
				},
			},
			cloudwatch: cloudwatch.CloudWatch{
				LogGroups: []cloudwatch.LogGroup{
					{
						Metadata: types.NewTestMetadata(),
						Arn:      types.String("arn:aws:cloudwatch:us-east-1:123456789012:log-group:cloudtrail-logging", types.NewTestMetadata()),
						MetricFilters: []cloudwatch.MetricFilter{
							{
								Metadata:   types.NewTestMetadata(),
								FilterName: types.String("VPCChange", types.NewTestMetadata()),
								FilterPattern: types.String(`{($.eventName=CreateVpc) || 
					($.eventName=DeleteVpc) || ($.eventName=ModifyVpcAttribute) || 
					($.eventName=AcceptVpcPeeringConnection) || ($.eventName=CreateVpcPeeringConnection) || 
					($.eventName=DeleteVpcPeeringConnection) || ($.eventName=RejectVpcPeeringConnection) || 
					($.eventName=AttachClassicLinkVpc) || ($.eventName=DetachClassicLinkVpc) || 
					($.eventName=DisableVpcClassicLink) || ($.eventName=EnableVpcClassicLink)}`, types.NewTestMetadata()),
							},
						},
					},
				},
				Alarms: []cloudwatch.Alarm{
					{
						Metadata:   types.NewTestMetadata(),
						AlarmName:  types.String("VPCChange", types.NewTestMetadata()),
						MetricName: types.String("VPCChange", types.NewTestMetadata()),
						Metrics: []cloudwatch.MetricDataQuery{
							{
								Metadata: types.NewTestMetadata(),
								ID:       types.String("VPCChange", types.NewTestMetadata()),
							},
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "Multi-region CloudTrail has no filter for VPC changes",
			cloudtrail: cloudtrail.CloudTrail{
				Trails: []cloudtrail.Trail{
					{
						Metadata:                  types.NewTestMetadata(),
						CloudWatchLogsLogGroupArn: types.String("arn:aws:cloudwatch:us-east-1:123456789012:log-group:cloudtrail-logging", types.NewTestMetadata()),
						IsLogging:                 types.Bool(true, types.NewTestMetadata()),
						IsMultiRegion:             types.Bool(true, types.NewTestMetadata()),
					},
				},
			},
			cloudwatch: cloudwatch.CloudWatch{
				LogGroups: []cloudwatch.LogGroup{
					{
						Metadata:      types.NewTestMetadata(),
						Arn:           types.String("arn:aws:cloudwatch:us-east-1:123456789012:log-group:cloudtrail-logging", types.NewTestMetadata()),
						MetricFilters: []cloudwatch.MetricFilter{},
					},
				},
				Alarms: []cloudwatch.Alarm{
					{
						Metadata:  types.NewTestMetadata(),
						AlarmName: types.String("VPCChange", types.NewTestMetadata()),
						Metrics: []cloudwatch.MetricDataQuery{
							{},
						},
					},
				},
			},
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.CloudWatch = test.cloudwatch
			testState.AWS.CloudTrail = test.cloudtrail
			results := requireVPCChangeAlarm.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == requireVPCChangeAlarm.Rule().LongID() {
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