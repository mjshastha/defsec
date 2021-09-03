package config

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckAggregateAllRegions = rules.Register(
	rules.Rule{
		Provider:    provider.AWSProvider,
		Service:     "config",
		ShortCode:   "aggregate-all-regions",
		Summary:     "Config configuration aggregator should be using all regions for source",
		Impact:      "Sources that aren't covered by the aggregator are not include in the configuration",
		Resolution:  "Set the aggregator to cover all regions",
		Explanation: `The configuration aggregator should be configured with all_regions for the source. 

This will help limit the risk of any unmonitored configuration in regions that are thought to be unused.`,
		Links: []string{ 
			"https://docs.aws.amazon.com/config/latest/developerguide/aggregate-data.html",
		},
		Severity: severity.High,
	},
	func(s *state.State) (results rules.Results) {
		for _, x := range s.AWS.S3.Buckets {
			if x.Encryption.Enabled.IsFalse() {
				results.Add(
					"",
					x.Encryption.Enabled.Metadata(),
					x.Encryption.Enabled.Value(),
				)
			}
		}
		return
	},
)
