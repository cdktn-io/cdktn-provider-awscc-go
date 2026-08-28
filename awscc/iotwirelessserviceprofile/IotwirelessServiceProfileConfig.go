// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelessserviceprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotwirelessServiceProfileConfig struct {
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
	// LoRaWAN supports all LoRa specific attributes for service profile for CreateServiceProfile operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_service_profile#lo_ra_wan IotwirelessServiceProfile#lo_ra_wan}
	LoRaWan *IotwirelessServiceProfileLoRaWan `field:"optional" json:"loRaWan" yaml:"loRaWan"`
	// Name of service profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_service_profile#name IotwirelessServiceProfile#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of key-value pairs that contain metadata for the service profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotwireless_service_profile#tags IotwirelessServiceProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

