// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan


type SsmincidentsResponsePlanActionsSsmAutomation struct {
	// The document name to use when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ssmincidents_response_plan#document_name SsmincidentsResponsePlan#document_name}
	DocumentName *string `field:"optional" json:"documentName" yaml:"documentName"`
	// The version of the document to use when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ssmincidents_response_plan#document_version SsmincidentsResponsePlan#document_version}
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// The parameters with dynamic values to set when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ssmincidents_response_plan#dynamic_parameters SsmincidentsResponsePlan#dynamic_parameters}
	DynamicParameters interface{} `field:"optional" json:"dynamicParameters" yaml:"dynamicParameters"`
	// The parameters to set when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ssmincidents_response_plan#parameters SsmincidentsResponsePlan#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The role ARN to use when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ssmincidents_response_plan#role_arn SsmincidentsResponsePlan#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The account type to use when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ssmincidents_response_plan#target_account SsmincidentsResponsePlan#target_account}
	TargetAccount *string `field:"optional" json:"targetAccount" yaml:"targetAccount"`
}

