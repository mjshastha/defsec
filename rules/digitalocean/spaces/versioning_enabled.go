package spaces

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckVersioningEnabled = rules.Register(
	rules.Rule{
		Provider:    provider.DigitalOceanProvider,
		Service:     "spaces",
		ShortCode:   "versioning-enabled",
		Summary:     "Spaces buckets should have versioning enabled",
		Impact:      "Deleted or modified data would not be recoverable",
		Resolution:  "Enable versioning to protect against accidental or malicious removal or modification",
		Explanation: `Versioning is a means of keeping multiple variants of an object in the same bucket. You can use the Spaces (S3) Versioning feature to preserve, retrieve, and restore every version of every object stored in your buckets. With versioning you can recover more easily from both unintended user actions and application failures.`,
		Links: []string{ 
			"https://docs.aws.amazon.com/AmazonS3/latest/userguide/Versioning.html",
		},
		Severity: severity.Medium,
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
