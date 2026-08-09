// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionmatchingworkflow


type EntityresolutionMatchingWorkflowResolutionTechniquesProviderPropertiesIntermediateSourceConfiguration struct {
	// The s3 path that would be used to stage the intermediate data being generated during workflow execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/entityresolution_matching_workflow#intermediate_s3_path EntityresolutionMatchingWorkflow#intermediate_s3_path}
	IntermediateS3Path *string `field:"optional" json:"intermediateS3Path" yaml:"intermediateS3Path"`
}

