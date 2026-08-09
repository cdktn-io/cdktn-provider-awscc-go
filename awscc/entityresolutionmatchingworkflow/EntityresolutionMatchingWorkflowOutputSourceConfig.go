// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionmatchingworkflow


type EntityresolutionMatchingWorkflowOutputSourceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_matching_workflow#output EntityresolutionMatchingWorkflow#output}.
	Output interface{} `field:"required" json:"output" yaml:"output"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_matching_workflow#apply_normalization EntityresolutionMatchingWorkflow#apply_normalization}.
	ApplyNormalization interface{} `field:"optional" json:"applyNormalization" yaml:"applyNormalization"`
	// The Customer Profiles integration configuration for the output source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_matching_workflow#customer_profiles_integration_config EntityresolutionMatchingWorkflow#customer_profiles_integration_config}
	CustomerProfilesIntegrationConfig *EntityresolutionMatchingWorkflowOutputSourceConfigCustomerProfilesIntegrationConfig `field:"optional" json:"customerProfilesIntegrationConfig" yaml:"customerProfilesIntegrationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_matching_workflow#kms_arn EntityresolutionMatchingWorkflow#kms_arn}.
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
	// The S3 path to which Entity Resolution will write the output table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_matching_workflow#output_s3_path EntityresolutionMatchingWorkflow#output_s3_path}
	OutputS3Path *string `field:"optional" json:"outputS3Path" yaml:"outputS3Path"`
}

