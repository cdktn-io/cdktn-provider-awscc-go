// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceElasticGpuSpecifications struct {
	// The type of Elastic Graphics accelerator. Amazon Elastic Graphics is no longer available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_instance#type Ec2Instance#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

