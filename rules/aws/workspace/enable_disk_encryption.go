package workspace

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckEnableDiskEncryption = rules.Register(
	rules.Rule{
		Provider:    provider.AWSProvider,
		Service:     "workspace",
		ShortCode:   "enable-disk-encryption",
		Summary:     "Root and user volumes on Workspaces should be encrypted",
		Impact:      "Data can be freely read if compromised",
		Resolution:  "Root and user volume encryption should be enabled",
		Explanation: `Workspace volumes for both user and root should be encrypted to protect the data stored on them.`,
		Links: []string{ 
			"https://docs.aws.amazon.com/workspaces/latest/adminguide/encrypt-workspaces.html",
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
