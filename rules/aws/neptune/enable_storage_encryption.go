package neptune

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckEnableStorageEncryption = rules.Register(
	rules.Rule{
		Provider:    provider.AWSProvider,
		Service:     "neptune",
		ShortCode:   "enable-storage-encryption",
		Summary:     "Neptune storage must be encrypted at rest",
		Impact:      "Unencrypted sensitive data is vulnerable to compromise.",
		Resolution:  "Enable encryption of Neptune storage",
		Explanation: `Encryption of Neptune storage ensures that if their is compromise of the disks, the data is still protected.`,
		Links: []string{ 
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
