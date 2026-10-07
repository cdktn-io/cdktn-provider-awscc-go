// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow


type ScnDataIntegrationFlowSourcesS3Source struct {
	// The S3 bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#bucket_name ScnDataIntegrationFlow#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The Amazon S3 options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#options ScnDataIntegrationFlow#options}
	Options *ScnDataIntegrationFlowSourcesS3SourceOptions `field:"optional" json:"options" yaml:"options"`
	// The S3 prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#prefix ScnDataIntegrationFlow#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

