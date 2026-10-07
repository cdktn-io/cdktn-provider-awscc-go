// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloud9environmentec2


type Cloud9EnvironmentEc2Repositories struct {
	// The path within the development environment's default file system location to clone the AWS CodeCommit repository into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloud9_environment_ec2#path_component Cloud9EnvironmentEc2#path_component}
	PathComponent *string `field:"optional" json:"pathComponent" yaml:"pathComponent"`
	// The clone URL of the AWS CodeCommit repository to be cloned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloud9_environment_ec2#repository_url Cloud9EnvironmentEc2#repository_url}
	RepositoryUrl *string `field:"optional" json:"repositoryUrl" yaml:"repositoryUrl"`
}

