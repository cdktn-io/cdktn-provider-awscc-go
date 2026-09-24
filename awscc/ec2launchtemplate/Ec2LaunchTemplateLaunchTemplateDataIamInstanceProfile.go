// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfile struct {
	// The Amazon Resource Name (ARN) of the instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#arn Ec2LaunchTemplate#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The name of the instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#name Ec2LaunchTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

