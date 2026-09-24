// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationSourceAwsResources struct {
	// The Amazon Resource Name (ARN) of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#resource_arn DevopsagentAssociation#resource_arn}
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// Additional metadata for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#resource_metadata DevopsagentAssociation#resource_metadata}
	ResourceMetadata *string `field:"optional" json:"resourceMetadata" yaml:"resourceMetadata"`
	// Resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#resource_type DevopsagentAssociation#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

