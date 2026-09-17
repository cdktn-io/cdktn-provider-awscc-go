// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchmetricstream


type CloudwatchMetricStreamTags struct {
	// A unique identifier for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_metric_stream#key CloudwatchMetricStream#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// String which you can use to describe or define the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_metric_stream#value CloudwatchMetricStream#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

