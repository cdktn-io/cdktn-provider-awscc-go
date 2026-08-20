// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalinglaunchconfiguration


type AutoscalingLaunchConfigurationBlockDeviceMappings struct {
	// The device name exposed to the EC2 instance (for example, /dev/sdh or xvdh).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_launch_configuration#device_name AutoscalingLaunchConfiguration#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Parameters used to automatically set up EBS volumes when an instance is launched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_launch_configuration#ebs AutoscalingLaunchConfiguration#ebs}
	Ebs *AutoscalingLaunchConfigurationBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// Setting this value to true suppresses the specified device included in the block device mapping of the AMI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_launch_configuration#no_device AutoscalingLaunchConfiguration#no_device}
	NoDevice interface{} `field:"optional" json:"noDevice" yaml:"noDevice"`
	// The name of the virtual device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_launch_configuration#virtual_name AutoscalingLaunchConfiguration#virtual_name}
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

