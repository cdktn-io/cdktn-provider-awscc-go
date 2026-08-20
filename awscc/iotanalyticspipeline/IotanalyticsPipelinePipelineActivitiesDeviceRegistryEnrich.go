// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticspipeline


type IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrich struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_pipeline#attribute IotanalyticsPipeline#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_pipeline#name IotanalyticsPipeline#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_pipeline#next IotanalyticsPipeline#next}.
	Next *string `field:"optional" json:"next" yaml:"next"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_pipeline#role_arn IotanalyticsPipeline#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_pipeline#thing_name IotanalyticsPipeline#thing_name}.
	ThingName *string `field:"optional" json:"thingName" yaml:"thingName"`
}

