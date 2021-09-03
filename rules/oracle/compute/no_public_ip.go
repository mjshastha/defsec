package compute

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckNoPublicIp = rules.Register(
	rules.Rule{
		Provider:    provider.OracleProvider,
		Service:     "compute",
		ShortCode:   "no-public-ip",
		Summary:     "Compute instance requests an IP reservation from a public pool",
		Impact:      "The compute instance has the ability to be reached from outside",
		Resolution:  "Reconsider the use of an public IP",
		Explanation: `Compute instance requests an IP reservation from a public pool

The compute instance has the ability to be reached from outside, you might want to sonder the use of a non public IP.`,
		Links: []string{ 
		},
		Severity: severity.Critical,
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
