// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelTopicConfigurationListRecordSchema struct {
	// ARN of Glue Schema Registry resource used for table schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/msk_channel#gsr_arn MskChannel#gsr_arn}
	GsrArn *string `field:"optional" json:"gsrArn" yaml:"gsrArn"`
}

