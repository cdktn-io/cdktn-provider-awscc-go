// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisedataset


type IotsitewiseDatasetDatasetSourceSourceDetailKendra struct {
	// The knowledgeBaseArn details for the Kendra dataset source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_dataset#knowledge_base_arn IotsitewiseDataset#knowledge_base_arn}
	KnowledgeBaseArn *string `field:"optional" json:"knowledgeBaseArn" yaml:"knowledgeBaseArn"`
	// The roleARN details for the Kendra dataset source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_dataset#role_arn IotsitewiseDataset#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

