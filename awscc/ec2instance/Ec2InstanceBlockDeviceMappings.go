// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceBlockDeviceMappings struct {
	// The device name (for example, /dev/sdh or xvdh).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#device_name Ec2Instance#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Parameters used to automatically set up EBS volumes when the instance is launched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#ebs Ec2Instance#ebs}
	Ebs *Ec2InstanceBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#no_device Ec2Instance#no_device}.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#virtual_name Ec2Instance#virtual_name}.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

