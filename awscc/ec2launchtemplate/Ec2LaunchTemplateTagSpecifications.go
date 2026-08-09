// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateTagSpecifications struct {
	// The type of resource. To tag a launch template, ``ResourceType`` must be ``launch-template``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#resource_type Ec2LaunchTemplate#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#tags Ec2LaunchTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

