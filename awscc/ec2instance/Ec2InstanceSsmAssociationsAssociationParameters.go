// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceSsmAssociationsAssociationParameters struct {
	// The name of an input parameter that is in the associated SSM document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_instance#key Ec2Instance#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of an input parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_instance#value Ec2Instance#value}
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

