// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticspipeline


type IotanalyticsPipelinePipelineActivitiesLambda struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_pipeline#batch_size IotanalyticsPipeline#batch_size}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_pipeline#lambda_name IotanalyticsPipeline#lambda_name}.
	LambdaName *string `field:"optional" json:"lambdaName" yaml:"lambdaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_pipeline#name IotanalyticsPipeline#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_pipeline#next IotanalyticsPipeline#next}.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

