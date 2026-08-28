// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionmatchingworkflow


type EntityresolutionMatchingWorkflowOutputSourceConfigCustomerProfilesIntegrationConfig struct {
	// The Amazon Resource Name (ARN) of the Customer Profiles domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/entityresolution_matching_workflow#domain_arn EntityresolutionMatchingWorkflow#domain_arn}
	DomainArn *string `field:"optional" json:"domainArn" yaml:"domainArn"`
	// The Amazon Resource Name (ARN) of the Customer Profiles object type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/entityresolution_matching_workflow#object_type_arn EntityresolutionMatchingWorkflow#object_type_arn}
	ObjectTypeArn *string `field:"optional" json:"objectTypeArn" yaml:"objectTypeArn"`
}

