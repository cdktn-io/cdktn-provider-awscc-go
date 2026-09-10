// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceElasticInferenceAccelerators struct {
	// The number of elastic inference accelerators to attach to the instance. Amazon Elastic Inference is no longer available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#count Ec2Instance#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The type of elastic inference accelerator. Amazon Elastic Inference is no longer available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#type Ec2Instance#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

