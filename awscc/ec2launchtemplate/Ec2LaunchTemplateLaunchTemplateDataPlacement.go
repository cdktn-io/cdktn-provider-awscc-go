// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataPlacement struct {
	// The affinity setting for an instance on a Dedicated Host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#affinity Ec2LaunchTemplate#affinity}
	Affinity *string `field:"optional" json:"affinity" yaml:"affinity"`
	// The Availability Zone for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#availability_zone Ec2LaunchTemplate#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The Group Id of a placement group.
	//
	// You must specify the Placement Group *Group Id* to launch an instance in a shared placement group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#group_id Ec2LaunchTemplate#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The name of the placement group for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#group_name Ec2LaunchTemplate#group_name}
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The ID of the Dedicated Host for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#host_id Ec2LaunchTemplate#host_id}
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// The ARN of the host resource group in which to launch the instances.
	//
	// If you specify a host resource group ARN, omit the *Tenancy* parameter or set it to ``host``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#host_resource_group_arn Ec2LaunchTemplate#host_resource_group_arn}
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// The number of the partition the instance should launch in.
	//
	// Valid only if the placement group strategy is set to ``partition``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#partition_number Ec2LaunchTemplate#partition_number}
	PartitionNumber *float64 `field:"optional" json:"partitionNumber" yaml:"partitionNumber"`
	// Reserved for future use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#spread_domain Ec2LaunchTemplate#spread_domain}
	SpreadDomain *string `field:"optional" json:"spreadDomain" yaml:"spreadDomain"`
	// The tenancy of the instance. An instance with a tenancy of dedicated runs on single-tenant hardware.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#tenancy Ec2LaunchTemplate#tenancy}
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

