package network

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckNoPublicIngress = rules.Register(
	rules.Rule{
		Provider:    provider.AzureProvider,
		Service:     "network",
		ShortCode:   "no-public-ingress",
		Summary:     "An inbound network security rule allows traffic from /0.",
		Impact:      "The port is exposed for ingress from the internet",
		Resolution:  "Set a more restrictive cidr range",
		Explanation: `Network security rules should not use very broad subnets.

Where possible, segments should be broken into smaller subnets.`,
		Links: []string{ 
			"https://docs.microsoft.com/en-us/azure/security/fundamentals/network-best-practices",
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
