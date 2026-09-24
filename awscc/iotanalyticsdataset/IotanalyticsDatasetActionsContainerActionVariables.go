// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetActionsContainerActionVariables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotanalytics_dataset#dataset_content_version_value IotanalyticsDataset#dataset_content_version_value}.
	DatasetContentVersionValue *IotanalyticsDatasetActionsContainerActionVariablesDatasetContentVersionValue `field:"optional" json:"datasetContentVersionValue" yaml:"datasetContentVersionValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotanalytics_dataset#double_value IotanalyticsDataset#double_value}.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotanalytics_dataset#output_file_uri_value IotanalyticsDataset#output_file_uri_value}.
	OutputFileUriValue *IotanalyticsDatasetActionsContainerActionVariablesOutputFileUriValue `field:"optional" json:"outputFileUriValue" yaml:"outputFileUriValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotanalytics_dataset#string_value IotanalyticsDataset#string_value}.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotanalytics_dataset#variable_name IotanalyticsDataset#variable_name}.
	VariableName *string `field:"optional" json:"variableName" yaml:"variableName"`
}

