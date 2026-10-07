// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloud9environmentec2


type Cloud9EnvironmentEc2Tags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#key Cloud9EnvironmentEc2#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#value Cloud9EnvironmentEc2#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

