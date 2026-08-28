// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParametersKinesisStreamParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#partition_key PipesPipe#partition_key}.
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
}

