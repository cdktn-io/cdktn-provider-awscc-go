// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerdevice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerDeviceConfig struct {
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
	// The ID of the global network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_device#global_network_id NetworkmanagerDevice#global_network_id}
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The Amazon Web Services location of the device, if applicable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_device#aws_location NetworkmanagerDevice#aws_location}
	AwsLocation *NetworkmanagerDeviceAwsLocation `field:"optional" json:"awsLocation" yaml:"awsLocation"`
	// The description of the device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_device#description NetworkmanagerDevice#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The site location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_device#location NetworkmanagerDevice#location}
	Location *NetworkmanagerDeviceLocation `field:"optional" json:"location" yaml:"location"`
	// The device model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_device#model NetworkmanagerDevice#model}
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The device serial number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_device#serial_number NetworkmanagerDevice#serial_number}
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// The site ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_device#site_id NetworkmanagerDevice#site_id}
	SiteId *string `field:"optional" json:"siteId" yaml:"siteId"`
	// The tags for the device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_device#tags NetworkmanagerDevice#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The device type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_device#type NetworkmanagerDevice#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The device vendor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_device#vendor NetworkmanagerDevice#vendor}
	Vendor *string `field:"optional" json:"vendor" yaml:"vendor"`
}

