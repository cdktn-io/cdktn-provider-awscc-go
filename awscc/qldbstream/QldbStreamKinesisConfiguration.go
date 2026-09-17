// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qldbstream


type QldbStreamKinesisConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/qldb_stream#aggregation_enabled QldbStream#aggregation_enabled}.
	AggregationEnabled interface{} `field:"optional" json:"aggregationEnabled" yaml:"aggregationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/qldb_stream#stream_arn QldbStream#stream_arn}.
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
}

