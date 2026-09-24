// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipamexternalresourceverificationtoken


type Ec2IpamExternalResourceVerificationTokenTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_external_resource_verification_token#key Ec2IpamExternalResourceVerificationToken#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_external_resource_verification_token#value Ec2IpamExternalResourceVerificationToken#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

