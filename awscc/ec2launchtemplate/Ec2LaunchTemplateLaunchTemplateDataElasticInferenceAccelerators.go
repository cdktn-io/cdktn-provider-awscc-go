// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAccelerators struct {
	// The number of elastic inference accelerators to attach to the instance.   Default: 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#count Ec2LaunchTemplate#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The type of elastic inference accelerator. The possible values are eia1.medium, eia1.large, and eia1.xlarge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_launch_template#type Ec2LaunchTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

