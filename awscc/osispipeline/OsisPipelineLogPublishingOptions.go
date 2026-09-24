// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package osispipeline


type OsisPipelineLogPublishingOptions struct {
	// The destination for OpenSearch Ingestion Service logs sent to Amazon CloudWatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/osis_pipeline#cloudwatch_log_destination OsisPipeline#cloudwatch_log_destination}
	CloudwatchLogDestination *OsisPipelineLogPublishingOptionsCloudwatchLogDestination `field:"optional" json:"cloudwatchLogDestination" yaml:"cloudwatchLogDestination"`
	// Whether logs should be published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/osis_pipeline#is_logging_enabled OsisPipeline#is_logging_enabled}
	IsLoggingEnabled interface{} `field:"optional" json:"isLoggingEnabled" yaml:"isLoggingEnabled"`
}

