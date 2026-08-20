// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package osispipeline


type OsisPipelineVpcOptionsVpcAttachmentOptions struct {
	// Whether the pipeline should be attached to the provided VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/osis_pipeline#attach_to_vpc OsisPipeline#attach_to_vpc}
	AttachToVpc interface{} `field:"optional" json:"attachToVpc" yaml:"attachToVpc"`
	// The CIDR block to be reserved for OpenSearch Ingestion to create elastic network interfaces (ENIs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/osis_pipeline#cidr_block OsisPipeline#cidr_block}
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
}

