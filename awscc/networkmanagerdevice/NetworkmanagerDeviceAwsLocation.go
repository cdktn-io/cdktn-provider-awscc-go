// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerdevice


type NetworkmanagerDeviceAwsLocation struct {
	// The Amazon Resource Name (ARN) of the subnet that the device is located in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkmanager_device#subnet_arn NetworkmanagerDevice#subnet_arn}
	SubnetArn *string `field:"optional" json:"subnetArn" yaml:"subnetArn"`
	// The Zone that the device is located in.
	//
	// Specify the ID of an Availability Zone, Local Zone, Wavelength Zone, or an Outpost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkmanager_device#zone NetworkmanagerDevice#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

