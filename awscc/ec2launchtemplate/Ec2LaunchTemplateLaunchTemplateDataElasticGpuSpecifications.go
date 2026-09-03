// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecifications struct {
	// The type of Elastic Graphics accelerator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_launch_template#type Ec2LaunchTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

