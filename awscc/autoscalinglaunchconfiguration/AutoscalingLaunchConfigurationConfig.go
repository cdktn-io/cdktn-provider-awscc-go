// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalinglaunchconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AutoscalingLaunchConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Provides the unique ID of the Amazon Machine Image (AMI) that was assigned during registration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#image_id AutoscalingLaunchConfiguration#image_id}
	ImageId *string `field:"required" json:"imageId" yaml:"imageId"`
	// Specifies the instance type of the EC2 instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#instance_type AutoscalingLaunchConfiguration#instance_type}
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// For Auto Scaling groups that are running in a virtual private cloud (VPC), specifies whether to assign a public IP address to the group's instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#associate_public_ip_address AutoscalingLaunchConfiguration#associate_public_ip_address}
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#block_device_mappings AutoscalingLaunchConfiguration#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The ID of a ClassicLink-enabled VPC to link your EC2-Classic instances to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#classic_link_vpc_id AutoscalingLaunchConfiguration#classic_link_vpc_id}
	ClassicLinkVpcId *string `field:"optional" json:"classicLinkVpcId" yaml:"classicLinkVpcId"`
	// The IDs of one or more security groups for the VPC that you specified in the ClassicLinkVPCId property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#classic_link_vpc_security_groups AutoscalingLaunchConfiguration#classic_link_vpc_security_groups}
	ClassicLinkVpcSecurityGroups *[]*string `field:"optional" json:"classicLinkVpcSecurityGroups" yaml:"classicLinkVpcSecurityGroups"`
	// Specifies whether the launch configuration is optimized for EBS I/O (true) or not (false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#ebs_optimized AutoscalingLaunchConfiguration#ebs_optimized}
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Provides the name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance.
	//
	// The instance profile contains the IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#iam_instance_profile AutoscalingLaunchConfiguration#iam_instance_profile}
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// The ID of the Amazon EC2 instance you want to use to create the launch configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#instance_id AutoscalingLaunchConfiguration#instance_id}
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Controls whether instances in this group are launched with detailed (true) or basic (false) monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#instance_monitoring AutoscalingLaunchConfiguration#instance_monitoring}
	InstanceMonitoring interface{} `field:"optional" json:"instanceMonitoring" yaml:"instanceMonitoring"`
	// Provides the ID of the kernel associated with the EC2 AMI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#kernel_id AutoscalingLaunchConfiguration#kernel_id}
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// Provides the name of the EC2 key pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#key_name AutoscalingLaunchConfiguration#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The name of the launch configuration. This name must be unique per Region per account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#launch_configuration_name AutoscalingLaunchConfiguration#launch_configuration_name}
	LaunchConfigurationName *string `field:"optional" json:"launchConfigurationName" yaml:"launchConfigurationName"`
	// The metadata options for the instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#metadata_options AutoscalingLaunchConfiguration#metadata_options}
	MetadataOptions *AutoscalingLaunchConfigurationMetadataOptions `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// The tenancy of the instance, either default or dedicated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#placement_tenancy AutoscalingLaunchConfiguration#placement_tenancy}
	PlacementTenancy *string `field:"optional" json:"placementTenancy" yaml:"placementTenancy"`
	// The ID of the RAM disk to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#ram_disk_id AutoscalingLaunchConfiguration#ram_disk_id}
	RamDiskId *string `field:"optional" json:"ramDiskId" yaml:"ramDiskId"`
	// A list that contains the security groups to assign to the instances in the Auto Scaling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#security_groups AutoscalingLaunchConfiguration#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The maximum hourly price you are willing to pay for any Spot Instances launched to fulfill the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#spot_price AutoscalingLaunchConfiguration#spot_price}
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// The Base64-encoded user data to make available to the launched EC2 instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_launch_configuration#user_data AutoscalingLaunchConfiguration#user_data}
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

