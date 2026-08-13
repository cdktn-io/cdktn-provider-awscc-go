// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataEnclaveOptions struct {
	// If this parameter is set to ``true``, the instance is enabled for AWS Nitro Enclaves;
	//
	// otherwise, it is not enabled for AWS Nitro Enclaves.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_launch_template#enabled Ec2LaunchTemplate#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

