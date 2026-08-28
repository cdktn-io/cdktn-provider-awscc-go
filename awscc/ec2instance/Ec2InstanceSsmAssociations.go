// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceSsmAssociations struct {
	// The input parameter values to use with the associated SSM document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#association_parameters Ec2Instance#association_parameters}
	AssociationParameters interface{} `field:"optional" json:"associationParameters" yaml:"associationParameters"`
	// The name of an SSM document to associate with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#document_name Ec2Instance#document_name}
	DocumentName *string `field:"optional" json:"documentName" yaml:"documentName"`
}

