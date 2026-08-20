// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataTagSpecifications struct {
	// The type of resource to tag.
	//
	// You can specify tags for the following resource types only: ``instance`` | ``volume`` | ``network-interface`` | ``spot-instances-request``. If the instance does not include the resource type that you specify, the instance launch fails. For example, not all instance types include a volume.
	//  To tag a resource after it has been created, see [CreateTags](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_launch_template#resource_type Ec2LaunchTemplate#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags to apply to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_launch_template#tags Ec2LaunchTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

