// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelessserviceprofile


type IotwirelessServiceProfileLoRaWan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotwireless_service_profile#add_gw_metadata IotwirelessServiceProfile#add_gw_metadata}.
	AddGwMetadata interface{} `field:"optional" json:"addGwMetadata" yaml:"addGwMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotwireless_service_profile#pr_allowed IotwirelessServiceProfile#pr_allowed}.
	PrAllowed interface{} `field:"optional" json:"prAllowed" yaml:"prAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotwireless_service_profile#ra_allowed IotwirelessServiceProfile#ra_allowed}.
	RaAllowed interface{} `field:"optional" json:"raAllowed" yaml:"raAllowed"`
}

