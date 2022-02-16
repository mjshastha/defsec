package efs

import (
	"github.com/aquasecurity/defsec/provider/aws/efs"
	"github.com/aquasecurity/trivy-config-parsers/cloudformation/parser"
)

func getFileSystems(ctx parser.FileContext) (filesystems []efs.FileSystem) {

	filesystemResources := ctx.GetResourceByType("AWS::EFS::FileSystem")

	for _, r := range filesystemResources {

		filesystem := efs.FileSystem{
			Metadata:  r.Metadata(),
			Encrypted: r.GetBoolProperty("Encrypted"),
		}

		filesystems = append(filesystems, filesystem)
	}

	return filesystems
}
