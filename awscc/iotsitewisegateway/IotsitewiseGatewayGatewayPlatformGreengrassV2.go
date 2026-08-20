// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisegateway


type IotsitewiseGatewayGatewayPlatformGreengrassV2 struct {
	// The operating system of the core device in AWS IoT Greengrass V2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotsitewise_gateway#core_device_operating_system IotsitewiseGateway#core_device_operating_system}
	CoreDeviceOperatingSystem *string `field:"optional" json:"coreDeviceOperatingSystem" yaml:"coreDeviceOperatingSystem"`
	// The name of the CoreDevice in GreenGrass V2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotsitewise_gateway#core_device_thing_name IotsitewiseGateway#core_device_thing_name}
	CoreDeviceThingName *string `field:"optional" json:"coreDeviceThingName" yaml:"coreDeviceThingName"`
}

