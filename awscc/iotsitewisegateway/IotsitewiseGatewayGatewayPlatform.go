// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisegateway


type IotsitewiseGatewayGatewayPlatform struct {
	// A gateway that runs on AWS IoT Greengrass V2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_gateway#greengrass_v2 IotsitewiseGateway#greengrass_v2}
	GreengrassV2 *IotsitewiseGatewayGatewayPlatformGreengrassV2 `field:"optional" json:"greengrassV2" yaml:"greengrassV2"`
	// A gateway that runs on Siemens Industrial Edge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_gateway#siemens_ie IotsitewiseGateway#siemens_ie}
	SiemensIe *IotsitewiseGatewayGatewayPlatformSiemensIe `field:"optional" json:"siemensIe" yaml:"siemensIe"`
}

