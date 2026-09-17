// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccessinstance


type Ec2VerifiedAccessInstanceLoggingConfigurationsKinesisDataFirehose struct {
	// The ID of the delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_instance#delivery_stream Ec2VerifiedAccessInstance#delivery_stream}
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
	// Indicates whether logging is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_instance#enabled Ec2VerifiedAccessInstance#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

