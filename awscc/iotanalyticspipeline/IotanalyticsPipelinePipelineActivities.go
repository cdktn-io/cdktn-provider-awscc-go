// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticspipeline


type IotanalyticsPipelinePipelineActivities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_pipeline#add_attributes IotanalyticsPipeline#add_attributes}.
	AddAttributes *IotanalyticsPipelinePipelineActivitiesAddAttributes `field:"optional" json:"addAttributes" yaml:"addAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_pipeline#channel IotanalyticsPipeline#channel}.
	Channel *IotanalyticsPipelinePipelineActivitiesChannel `field:"optional" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_pipeline#datastore IotanalyticsPipeline#datastore}.
	Datastore *IotanalyticsPipelinePipelineActivitiesDatastore `field:"optional" json:"datastore" yaml:"datastore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_pipeline#device_registry_enrich IotanalyticsPipeline#device_registry_enrich}.
	DeviceRegistryEnrich *IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrich `field:"optional" json:"deviceRegistryEnrich" yaml:"deviceRegistryEnrich"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_pipeline#device_shadow_enrich IotanalyticsPipeline#device_shadow_enrich}.
	DeviceShadowEnrich *IotanalyticsPipelinePipelineActivitiesDeviceShadowEnrich `field:"optional" json:"deviceShadowEnrich" yaml:"deviceShadowEnrich"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_pipeline#filter IotanalyticsPipeline#filter}.
	Filter *IotanalyticsPipelinePipelineActivitiesFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_pipeline#lambda IotanalyticsPipeline#lambda}.
	Lambda *IotanalyticsPipelinePipelineActivitiesLambda `field:"optional" json:"lambda" yaml:"lambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_pipeline#math IotanalyticsPipeline#math}.
	Math *IotanalyticsPipelinePipelineActivitiesMath `field:"optional" json:"math" yaml:"math"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_pipeline#remove_attributes IotanalyticsPipeline#remove_attributes}.
	RemoveAttributes *IotanalyticsPipelinePipelineActivitiesRemoveAttributes `field:"optional" json:"removeAttributes" yaml:"removeAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotanalytics_pipeline#select_attributes IotanalyticsPipeline#select_attributes}.
	SelectAttributes *IotanalyticsPipelinePipelineActivitiesSelectAttributes `field:"optional" json:"selectAttributes" yaml:"selectAttributes"`
}

